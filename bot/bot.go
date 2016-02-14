package bot

import (
	"github.com/cpalone/maimai.v3/proto"

	"euphoria.io/scope"
	"github.com/apex/log"
)

// Bot represents a connection to a euphoria room. It may be a mock room.
type Bot struct {
	dialer   proto.Dialer
	conn     proto.Connection
	handlers []proto.Handler
	ctx      scope.Context
}

// NewBot creates a new Bot with the given context, handlers, and dialer.
func NewBot(ctx scope.Context, handlers []proto.Handler, dialer proto.Dialer) *Bot {
	return &Bot{
		dialer:   dialer,
		conn:     nil,
		handlers: handlers,
		ctx:      ctx,
	}
}

// Run dials the room, starts the handlers, and starts listening for incoming
// packets.
func (b *Bot) Run() error {
	logger := b.ctx.Get("logger").(*log.Logger)
	conn, err := b.dialer.Dial()
	if err != nil {
		logger.Fatalf("Bot.Run: error connecting (%s)", err)
		return nil
	}
	b.conn = conn
	defer b.conn.Close()

	errChan := make(chan error)
	for _, handler := range b.handlers {
		go func() {
			if err := handler.Run(b.ctx.Fork()); err != nil {
				logger.Fatalf("Bot.Run: error running handler (%s)", err)
				errChan <- err
			}
		}()
	}

	go func() {
		for {
			p, err := b.conn.Receive()
			if err != nil {
				logger.Fatalf("Bot.Run: error receiving packet (%s)", err)
				errChan <- err
				return
			}
			for _, handler := range b.handlers {
				if err := handler.HandleIncoming(b.conn, p); err != nil {
					logger.Fatalf("Bot.Run: error in Handler.HandleIncoming (%s)", err)
					errChan <- err
					return
				}
			}
		}
	}()

	select {
	case err := <-errChan:
		logger.Fatalf("Fatal error in Bot.Run (%s)", err)
		return err
	case <-b.ctx.Done():
		logger.Infof("Bot.Run: context is cancelled (%s)", b.ctx.Err())
		return nil
	}
}
