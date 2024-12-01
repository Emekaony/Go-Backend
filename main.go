package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/greet", greet)

	http.HandleFunc("/", sayHello)

	http.ListenAndServe("localhost:8000", nil)
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Greetings")
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}
