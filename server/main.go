package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/Odraxs/go-z-v-mail/server/app"
)

func main() {
	config := app.LoadZincsearchCredentials()
	app := app.New(config)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx)
	if err != nil {
		log.Println("failed to start app:", err)
	}
}
