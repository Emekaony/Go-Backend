package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ApiResponse struct {
	Message string `json:"message"`
}

// main entry point of every go code
func main() {
	// create the handler that will handle routes for you
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// make sure to write this to the responseWriter
		fmt.Fprintln(w, "Hello world")
	})

	http.HandleFunc("/api", apiHandler)

	// now how do we make sure this server is available somewhere?
	fmt.Println("server is listening on port 3000")
	http.ListenAndServe("localhost:3000", nil)
}

func apiHandler(w http.ResponseWriter, _ *http.Request) {
	apiResponse := ApiResponse{
		Message: "Authorized for this API endpoint",
	}

	// now make sure the response header is in JSON format
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(apiResponse)
}
