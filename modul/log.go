package modul

import (
	"context"
	"fmt"
	"net/http"

	Mo "stockbit-testcode/movie"
	U "stockbit-testcode/utils"
)

func GetLogMovie(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ctx, cancel := context.WithCancel(context.Background())

		defer cancel()

		movie, err := Mo.GetAll(ctx)

		if err != nil {
			fmt.Println(err)
		}

		U.ResponseJSON(w, movie, http.StatusOK)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusNotFound)
	return
}
