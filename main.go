package main

import (
	"banking/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	start()
}

func start() {
	// define our own multiplexer
	// mux := http.NewServeMux()
	router := mux.NewRouter()
	// http routes
	router.HandleFunc("/greet", handlers.Greet)
	router.HandleFunc("/", handlers.SayHello)
	router.HandleFunc("/customers/{customer_id}", handlers.GetCustomer)
	router.HandleFunc("/customers", handlers.GetAllCustomers)

	fmt.Println("Server is listening on port 8000")
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
