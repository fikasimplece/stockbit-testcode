package modul

import (
	"context"
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
			message := err.Error()
			U.ResponseJSON(w, "error-message : "+message, http.StatusNotFound)
			return

		}

		U.ResponseJSON(w, movie, http.StatusOK)
		return
	}

	http.Error(w, "error-message : Tidak di ijinkan", http.StatusNotFound)
	return
}
