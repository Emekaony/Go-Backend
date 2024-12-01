package main

import (
	"banking/models"
	"banking/tinkering"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	tinkering.RandomStuff()
	// http routes
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/customers", getAllCustomers)

	fmt.Println("Server is listening on port 8080")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Greetings")
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	// make a list of customers to send to the client
	w.Header().Set("Content-Type", "application/json")
	customers := []models.Customer{
		{Name: "nnaemeka", City: "Sacramento", Zipcode: "223"},
		{Name: "kamsi", City: "Coulumbus", Zipcode: "8900"},
	}
	json.NewEncoder(w).Encode(customers)
}
