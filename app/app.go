package app

import (
	"context"
	"fmt"
	"net/http"
)

type App struct {
	router http.Handler
}

func New() *App {
	return &App{
		router: loadRouter(),
	}
}

func (this *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":8080",
		Handler: this.router,
	}

	if err := server.ListenAndServe(); err != nil {
		return fmt.Errorf("starting the server: %w", err)
	}

	return nil
}
