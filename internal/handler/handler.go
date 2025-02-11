package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sonochiwa/avito-shop/api/rest/models"
	"github.com/sonochiwa/avito-shop/internal/service"
)

type Handler struct {
	service service.Service
}

func New(service service.Service) Handler {
	return Handler{
		service: service,
	}
}

func (h Handler) NewRouter() http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	// without authenticated routes
	router.Post("/auth", h.authHandler)

	// with authentication routes
	router.Route("/api", func(r chi.Router) {
		r.Get("/info", h.infoHandler)
		r.Post("/sendCoin", h.sendCoinHandler)
		r.Post("/buy/{item}", h.buyItemHandler)
	})

	return router
}

func (h Handler) infoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userID := 1

	response, err := h.service.GetInfo(context.Background(), userID)
	if err != nil {
		responseError(w, err.Error())
		return
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		responseError(w, err.Error())
		return
	}
}

func (h Handler) sendCoinHandler(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (h Handler) buyItemHandler(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (h Handler) authHandler(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func responseError(w http.ResponseWriter, message string) {
	response := models.ErrorsResponse{
		Errors: message,
	}

	w.WriteHeader(http.StatusInternalServerError)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Print("error when json.NewEncoder(r).Encode(response): %w", err)
	}
}
