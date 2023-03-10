package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/prajapatiomkar/todo-app-in-golang/router"
	"github.com/rs/cors"
)

func main() {
	port := os.Getenv(("PORT"))
	fmt.Println("Hello")
	r := router.Router()
	fmt.Println("Starting Server On The Port: 8080")
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "DELETE", "POST", "PUT"},
	})

	handlers := c.Handler(r)
	log.Fatal(http.ListenAndServe(":"+port, handlers))
	// log.Fatal(http.ListenAndServe(":8080", handlers))

}
