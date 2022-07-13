package Update

import (
	"github.com/jmoiron/sqlx"
	"log"
)

// UpdateRating update database
func NewMovieRating(db *sqlx.DB, movieName string, newRating float32) {
	_, err := db.Exec("UPDATE GoodMovies SET Rating = ? WHERE MovieName = ?", newRating, movieName)
	if err != nil {
		log.Fatalln(err)
	}
}
