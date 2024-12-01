package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customer_name := vars["customer_name"]
	fmt.Fprintf(w, "Post request received for new customer %s", customer_name)
	// vars := mux.Vars(r)
	// customer_name := vars["customer_name"]
}
