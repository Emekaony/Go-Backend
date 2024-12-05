package main

import (
	"banking/handlers"
	"fmt"
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
	// get requests
	router.HandleFunc("/", handlers.SayHello).Methods(http.MethodGet)
	router.HandleFunc("/greet", handlers.Greet).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", handlers.GetCustomer).Methods(http.MethodGet) // that regex automatically handles cases where customer_id is not a number
	router.HandleFunc("/customers", handlers.GetAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/api/time", handlers.GetTime).Methods(http.MethodGet)
	router.HandleFunc("/file", handlers.GetFile).Methods(http.MethodGet)

	// create requests
	router.HandleFunc("/customers/{customer_name}", handlers.CreateCustomer).Methods(http.MethodPost)

	fmt.Println("Server is listening on port 8000")
	http.ListenAndServe("localhost:8000", router)
}
