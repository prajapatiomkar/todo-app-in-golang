package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prajapatiomkar/todo-app-in-golang/router"
)

func main() {
	fmt.Println("Hello")
	r := router.Router()
	fmt.Println("Starting Server On The Port: 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
