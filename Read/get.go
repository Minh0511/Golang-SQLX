package Read

import (
	"Go_project/Config"
	"fmt"
	"github.com/jmoiron/sqlx"
)

// GetAllMovies get all movies from database
func GetAllMovies(db *sqlx.DB) {
	var movies []Config.GoodMovies
	rows, err := db.Queryx("SELECT * FROM GoodMovies ORDER BY Rating DESC")
	if err != nil {
		panic(err.Error())
	}
	for rows.Next() {
		var movie Config.GoodMovies
		err = rows.StructScan(&movie)
		if err != nil {
			panic(err.Error())
		}
		movies = append(movies, movie)
	}
	for _, movie := range movies {
		fmt.Println("Movie name: ", movie.MovieName, "| Movie genre: ", movie.MovieGenre, "| Movie Director: ", movie.Director, "| Movie Rating: ", movie.Rating)
	}
}

// GetMovieByName select with condition using sqlx.Queryx
func GetMovieByName(db *sqlx.DB, movieName string) {
	var movie Config.GoodMovies
	err := db.QueryRow("SELECT * FROM GoodMovies WHERE MovieName = ?", movieName).Scan(&movie.MovieName, &movie.MovieGenre, &movie.Director, &movie.Rating)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Movie name: ", movie.MovieName, "| Movie genre: ", movie.MovieGenre, "| Movie Director: ", movie.Director, "| Movie Rating: ", movie.Rating)

}
