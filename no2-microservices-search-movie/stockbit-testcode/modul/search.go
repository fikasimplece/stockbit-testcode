package modul

import (
	"encoding/json"
	"log"
	"net/http"

	R "stockbit-testcode/movie"
	U "stockbit-testcode/utils"
)

var client = new(http.Client)

func GetListData(w http.ResponseWriter, r *http.Request) {

	title := r.URL.Query()["searchword"][0]
	page := r.URL.Query()["pagination"][0]

	go R.Insert(title, "search")

	req, _ := http.NewRequest("GET", "http://www.omdbapi.com", nil)

	q := req.URL.Query()
	q.Add("apikey", "faf7e5bb")
	q.Add("s", title)
	q.Add("page", page)
	req.URL.RawQuery = q.Encode()

	response, err := client.Do(req)
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		log.Fatal(err)
	}

	var generic map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&generic)
	if err != nil {
		log.Fatal(err)
	}

	U.ResponseJSON(w, generic, http.StatusOK)

}

func GetSingleDetailMovie(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query()["title"][0]

	go R.Insert(title, "detail")

	req, _ := http.NewRequest("GET", "http://www.omdbapi.com", nil)

	q := req.URL.Query()
	q.Add("apikey", "faf7e5bb")
	q.Add("t", title)
	req.URL.RawQuery = q.Encode()

	response, err := client.Do(req)
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		log.Fatal(err)
	}

	var generic map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&generic)
	if err != nil {
		log.Fatal(err)
	}

	U.ResponseJSON(w, generic, http.StatusOK)
}
