package main

import (
	"banking/handlers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	start()
}

func start() {
	// http routes
	http.HandleFunc("/greet", handlers.Greet)
	http.HandleFunc("/", handlers.SayHello)
	http.HandleFunc("/customers", handlers.GetAllCustomers)

	fmt.Println("Server is listening on port 8080")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
