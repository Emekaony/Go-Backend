package handlers

import (
	"banking/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func GetTime(w http.ResponseWriter, r *http.Request) {
	// make sure to set the content-type to json
	w.Header().Set("Content-Type", "application/json")
	responses := []models.Time{}
	// we access the query parameters from the response object
	queryParams := r.URL.Query()
	// this returns an array incase we passed multiple, pretty cool
	location_var := strings.Split(queryParams["tz"][0], ",")

	// handle multiple timezones
	for _, param := range location_var {
		location, err := time.LoadLocation(param)
		if err != nil {
			// this means that the location parameter was invalid
			fmt.Fprint(w, "invalid timezone")
			w.WriteHeader(http.StatusNotFound)
		} else {
			currentFormattedTime := time.Now().UTC().In(location).Format("2006-01-02 15:04:05 -0700 MST")
			currTime := models.Time{CurrentTime: currentFormattedTime}
			responses = append(responses, currTime)
		}
	}
	json.NewEncoder(w).Encode(responses)
}
