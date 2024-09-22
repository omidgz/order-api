package app

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"gorm.io/gorm"
)

type App struct {
	router http.Handler
	DB     *gorm.DB
}

func (this *App) Start(ctx context.Context) error {
	this.loadDB()
	this.loadRouter()

	server := &http.Server{
		Addr:    ":8080",
		Handler: this.router,
	}

	log.Printf("Starting the server on %s...", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		return fmt.Errorf("starting the server: %w", err)
	}

	return nil
}
