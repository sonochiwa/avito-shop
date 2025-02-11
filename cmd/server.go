package cmd

import (
	"log"
	"net/http"

	"github.com/sonochiwa/avito-shop/internal/database"
	"github.com/sonochiwa/avito-shop/internal/handler"
	"github.com/sonochiwa/avito-shop/internal/repository"
	"github.com/sonochiwa/avito-shop/internal/service"
)

func Run() {
	pool := database.NewPool()
	defer pool.Close()

	repositoryInstance := repository.New(pool)
	serviceInstance := service.New(repositoryInstance)
	handlerInstance := handler.New(serviceInstance)

	srv := http.Server{
		Addr:    ":8080",
		Handler: handlerInstance.NewRouter(),
	}

	log.Print("starting server...")
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("error when starting server: %v", err)
	}
}
