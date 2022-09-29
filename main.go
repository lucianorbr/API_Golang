package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/lucianorbr/API_Golang/configs"
	"github.com/lucianorbr/API_Golang/handlers"
	"net/http"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Post("/", handlers.Create)
	r.Put("/{id}", handlers.Update)
	r.Delete("/{id}", handlers.Delete)
	r.Get("/", handlers.Get)
	r.Get("/{id}", handlers.Get)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}
