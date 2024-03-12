package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

type App struct {
	router      http.Handler
}

func New() *App {
	app := &App{}

	app.loadRoutes()
	return app
}

func (app *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":3001",
		Handler: app.router,
	}

	log.Println("Starting server")

	ch := make(chan error, 1)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			ch <- fmt.Errorf("failed to start the server: %w", err)
		}
		close(ch)
	}()

	select {
	case err := <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		return server.Shutdown(timeout)
	}
}
