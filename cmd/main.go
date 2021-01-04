package main

import (
	"github.com/StanislavDimitrenco/todoapp"
	"github.com/StanislavDimitrenco/todoapp/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)

	server := new(todoapp.Server)

	if err := server.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server %s", err.Error())
	}
}
