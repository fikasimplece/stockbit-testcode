package movie

import (
	"context"
	"fmt"
	"log"
	"stockbit-testcode/config"
	"stockbit-testcode/models"
	"time"
)

const (
	table      = "log_movie"
	layoutDate = "2006-01-02 15:04:05"
)

var dateNow = time.Now().Format(layoutDate)

func GetAll(ctx context.Context) ([]models.Movie, error) {
	var movies []models.Movie

	db, err := config.MySQL()

	if err != nil {
		log.Fatal(err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By ID DESC", table)

	rowQuery, _ := db.QueryContext(ctx, queryText)

	for rowQuery.Next() {
		var movie models.Movie
		var createdAt, updatedAt string

		if err = rowQuery.Scan(&movie.ID,
			&movie.ImdbID,
			&createdAt,
			&updatedAt,
			&movie.Path); err != nil {
			return nil, err
		}

		if err != nil {
			log.Fatal(err)
		}

		movies = append(movies, movie)
	}

	return movies, nil
}

func Insert(title string, path string) error {

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}
	insForm, err := db.Prepare("INSERT INTO log_movie (ImdbID, Path, createdAt, updatedAt) values(?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	insForm.Exec(title, path, dateNow, dateNow)

	if err != nil {
		return err
	}
	return nil

}
