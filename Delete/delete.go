package Delete

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

// DeleteAll delete all data from database
func DeleteAll(db *sqlx.DB) {
	start := time.Now()
	_, err := db.Exec("DELETE FROM GoodMovies")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Time elapsed: ", time.Since(start))
}

// DeleteMovieFromDB delete movie with condition
func DeleteMovieFromDB(db *sqlx.DB, movieName string) {
	start := time.Now()
	_, err := db.Exec("DELETE FROM GoodMovies WHERE MovieName = ?", movieName)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Time elapsed: ", time.Since(start))

}
