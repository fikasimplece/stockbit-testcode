//Nama : Rafika Sihombing
//28-04-2021
package main

import (
	"log"
	"net/http"
	"stockbit-testcode/config"

	M "stockbit-testcode/modul"
)

func main() {

	db, e := config.MySQL()

	if e != nil {
		log.Fatal(e)
	}

	eb := db.Ping()
	if eb != nil {
		panic(eb.Error())
	}

	http.HandleFunc("/movie-log", M.GetLogMovie)
	http.HandleFunc("/get/search", M.GetListData)
	http.HandleFunc("/get/detail", M.GetSingleDetailMovie)

	errs := http.ListenAndServe(":7000", nil)

	if errs != nil {
		log.Fatal(errs)
	}

}
