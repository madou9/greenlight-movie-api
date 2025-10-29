package main

import (
	"encoding/json"
	"net/http"
)



func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request){

	// create a map which holds the information that we want to send in the response
	data := map[string]string{
		"status" : "available",
		"environment" : app.config.env,
		"version" : version,
	}

	// pass the map to the json.Marshal() function. This returns a []byte slice 
	// containing thee encoded JSON. If there was an error, we log it and send the client
	// a generic error message.
	js, err := json.Marshal(data)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
		return
	}

	// Append a newline to the JSON. This is just a small nicety to make it easier to 
	// view in terminal applications.
	js = append(js, '\n')

	// At this point we know that encoding the data worked without any problems, so we 
	// can safely set any neecessary HTTP headers for a sucessful response.
	w.Header().Set("content-Type", "application/json")

	// Use w.write() to send the []byte slice containing the JSON as the response body.
	w.Write(js)

}