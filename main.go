package main

import (
	"fmt"
	"net/http"
	"postgres/api/configs"
	"postgres/api/handlers"

	"github.com/go-chi/chi/v5"
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
	r.Get("/", handlers.List)
	r.Get("/{id}", handlers.Get)
	fmt.Println("System UP on port:", configs.GetServerPort())
	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)

}
