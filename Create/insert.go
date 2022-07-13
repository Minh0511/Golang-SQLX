package Create

import (
	"Go_project/Config"
	"fmt"
	"github.com/jmoiron/sqlx"
	"math/rand"
	"time"
)

//insert many random data to database using sqlx
func InsertBatch(db *sqlx.DB) {
	start := time.Now()
	var goodMovies []Config.GoodMovies

	for i := 0; i < 10000; i++ {
		goodMovies = append(goodMovies, Config.GoodMovies{
			MovieName:  Config.MoviesName[rand.Intn(len(Config.MoviesName))],
			MovieGenre: Config.MoviesGenre[rand.Intn(len(Config.MoviesGenre))],
			Director:   Config.MoviesDirector[rand.Intn(len(Config.MoviesDirector))],
			Rating:     float32(Config.MoviesRating[rand.Intn(len(Config.MoviesRating))]),
		})
	}

	_, err := db.NamedExec(`INSERT INTO GoodMovies(MovieName, MovieGenre, Director, Rating) VALUES (:MovieName, :MovieGenre, :Director, :Rating)`, goodMovies)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Time elapsed: ", time.Since(start))
}

//insert 1 data to database using sqlx
func InsertMovies(db *sqlx.DB) {
	tx := db.MustBegin()
	tx.MustExec(`INSERT INTO GoodMovies(MovieName, MovieGenre, Director, Rating) VALUES (?, ?, ?, ?)`,
		"Avengers", "Action", "John Wark", 9.1)
	tx.Commit()
}
