package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/Odraxs/go-z-v-mail/server/app"
	"github.com/Odraxs/go-z-v-mail/server/config"
)

func main() {
	config.LoadZincsearchCredentials()
	app := app.New()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx)
	if err != nil {
		log.Println("failed to start app:", err)
	}
}
