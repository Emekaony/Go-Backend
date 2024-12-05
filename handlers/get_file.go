package handlers

import "net/http"

func GetFile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "models/index.html")
}
