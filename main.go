package main

import (
	"Go_project/Update"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

func main() {
	db, err := sqlx.Open("mysql", "root:1@tcp(127.0.0.1:3306)/Movies")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	//insert into database
	//Create.InsertMovies(db)

	//insert random data
	//Create.InsertBatch(db)

	//delete all data from DB
	//Delete.DeleteAll(db)

	//delete movie with the name "Doom"
	//Delete.DeleteMovieFromDB(db, "Doom")

	//update rating of movie with the name "Avengers"
	Update.NewMovieRating(db, "Avengers", 9.1)

	// Query the database
	//Read.GetAllMovies(db)

	//get movie by name "Avengers"
	//Read.GetMovieByName(db, "Avengers")
}
