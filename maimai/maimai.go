package main

import (
	"fmt"
	"os"

	"github.com/cpalone/maimai.v3/bot"
	"github.com/cpalone/maimai.v3/connection"
	"github.com/cpalone/maimai.v3/handlers"
	"github.com/cpalone/maimai.v3/proto"

	"euphoria.io/scope"
	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
)

const url = "wss://euphoria.io/room/%s/ws"

func main() {
	logger := &log.Logger{
		Handler: cli.New(os.Stdout),
		Level:   log.DebugLevel,
	}
	ctx := scope.New()
	ctx.Set("logger", logger)
	b := bot.NewBot(ctx, []proto.Handler{&handlers.PingHandler{}}, connection.NewWSDialer(ctx.Fork(), fmt.Sprintf(url, "test")))
	if err := b.Run(); err != nil {
		logger.Fatalf("Bot.Run: Fatal error (%s)", err)
	}
}
