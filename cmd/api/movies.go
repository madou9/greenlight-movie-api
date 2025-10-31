package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/madou9/greenlight-movie-api.git/internal/data"
)

func(app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
}


func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	// Create a new instance of the Movie struct, containing the ID we extracted from
	// the URL and some dummy data. Also notice that we deliberately haven't set a 
	// value for the Year field

	movie := data.Movie{
		ID:			id,
		CreatedAt: 	time.Now(),
		Title: 		"Casablanca",
		Runtime: 	102,
		Genres:	 	[]string{"drama", "romance", "war"},
		Version: 	1,	
	}

	// Create an envelope{"movie": movie} instance and pass it to writeJSON(), instead
	// of passing the plan movie struct
	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
