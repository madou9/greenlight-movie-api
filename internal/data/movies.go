package data

import "time"

// Annotate the Movie struct with struct tags to control how the keys appear in the
// JSON-encoded output.
type Movie struct {
	ID			int64 		`json:"id"`			// unique integer ID for movie
	CreatedAt	time.Time	`json:"created_At"` // Timestamp for when the movie is added to our database
	Title		string 		`json:"title"`  	// Movie title
	Year 		int32 		`json:"year"`   	// Movie release year
	Runtime		int32 		`json:"runtime"`	// Movie runtime (in minutes)
	Genres		[]string 	`json:"genres"` 	// Slice of genres for the movie (romance, comedy, etc.)
	Version		int32  		`json:"version"` 	// time the movie information is updated
}
