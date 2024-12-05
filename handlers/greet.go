package handlers

import (
	"fmt"
	"net/http"
)

func Greet(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Greetings, this is Emeka.")
}
