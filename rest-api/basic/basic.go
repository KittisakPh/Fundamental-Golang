package basic

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type movie struct {
	ImdbID      string  `json:"imdbID"`
	Title       string  `json:"title"`
	Year        int     `json:"year"`
	Rating      float64 `json:"rating"`
	IsSuperHero bool    `json:"isSuperHero"`
}

var movies []movie

func moviesHandler(w http.ResponseWriter, req *http.Request) {
	method := req.Method
	switch method {
	case "GET":
		movies, err := json.Marshal(movies)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "error : %v", err)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Write(movies)
		return
	case "POST":
		bodyBytes, err := io.ReadAll(req.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "error : %v", err)
			return
		}
		movie := movie{}
		err = json.Unmarshal(bodyBytes, &movie)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "error : %v", err)
			return
		}
		movies = append(movies, movie)
		fmt.Fprintf(w, "hello %s created movies\n", "POST")
		return
	case "DELETE":
		fmt.Fprintf(w, "hello %s movies", "DELETE")
		return
	}

}

func Basic() {
	http.HandleFunc("/movies", moviesHandler)
	// localhost = loop back
	err := http.ListenAndServe("localhost:2565", nil)
	if err != nil {
		log.Fatal(err)
	}
}