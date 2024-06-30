package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Video struct {
	Name     string `json:"name" db:"name"`
	Duration int    `json:"duration" db:"duration"`
	Actor    string `json:"actor,omitempty"`
	Director string `json:"director"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/helloworld", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello world"))
	}).Methods("POST", "GET")

	router.HandleFunc("/book/{name}/{pageNumber}", func(writer http.ResponseWriter, request *http.Request) {
		values := mux.Vars(request)
		name := values["name"]
		pageNumber := values["pageNumber"]
		// search :=
		fmt.Fprintf(writer, "you request the page number :%v of the book: %v", pageNumber, name)
		// fmt.Fprintf(writer, "you searched for word : %v", search)
	}).Methods("POST", "GET")

	router.HandleFunc("/saveMovie", func(writer http.ResponseWriter, request *http.Request) {
		var movie Video
		err := json.NewDecoder(request.Body).Decode(&movie)
		if err != nil {
			fmt.Fprintf(writer, "error occurred")
		}
		fmt.Fprintf(writer, "movie created:%v, %v", movie.Name, movie.Duration)
	}).Methods("POST")

	router.HandleFunc("/getMovie", func(writer http.ResponseWriter, request *http.Request) {
		var movie = Video{
			Name:     "hello movie",
			Duration: 200,
			Actor:    "test",
			Director: "director",
		}

		b, err := json.Marshal(movie)
		if err != nil {

		} else {
			writer.Write(b)
		}
	}).Methods("GET")

	http.ListenAndServe(":8080", router)
}
