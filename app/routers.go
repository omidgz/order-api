package app

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/omidgz/order-api/handler"
)

func (this *App) loadRouter() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/orders", this.loadOrderRouter)

	this.router = router
}

func (this *App) loadOrderRouter(router chi.Router) {
	orderHandler := handler.NewOrderHandler(this.DB)

	router.Post("/", orderHandler.Create)
	router.Get("/", orderHandler.List)
	router.Get("/{id}", orderHandler.GetByID)
	router.Put("/", orderHandler.UpdateByID)
	router.Delete("/{id}", orderHandler.DeleteByID)
	log.Println("Order routes loaded")

}
