package handlers

import (
	"banking/models"
	"encoding/json"
	"encoding/xml"
	"log"
	"net/http"
)

func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	// make a list of customers to send to the client
	w.Header().Set("Content-Type", "application/json")
	customers := []models.Customer{
		{Name: "nnaemeka", City: "Sacramento", Zipcode: "223"},
		{Name: "kamsi", City: "Coulumbus", Zipcode: "8900"},
	}
	// check what type of content the client is requesting
	switch r.Header.Get("Content-Type") {
	case "application/json":
		json.NewEncoder(w).Encode(customers)
	case "application/xml":
		xml.NewEncoder(w).Encode(customers)
	default:
		log.Fatal("Unknown content type")
	}
}
