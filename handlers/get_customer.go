package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["customer_id"])
}
