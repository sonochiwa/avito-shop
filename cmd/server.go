package cmd

import (
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

	handlerInstance.NewRouter()

	srv := http.Server{
		Addr:    ":8080",
		Handler: handlerInstance.NewRouter(),
	}

	srv.ListenAndServe()
}
