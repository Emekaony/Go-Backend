package handlers

import (
	"fmt"
	"net/http"
)

func SayHello(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Home Page")
}
