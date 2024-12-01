package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/greet", greet)

	http.HandleFunc("/", sayHello)

	fmt.Println("Server is listening on port 8080")
	http.ListenAndServe("localhost:8000", nil)
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Greetings")
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}
