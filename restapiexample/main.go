package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/jackc/pgx/v4/pgxpool"
)

const insert = "insert into video (name, duration) values ('test', 20)"

type Video struct {
	Name     string `json:"name" db:"name"`
	Duration int    `json:"duration" db:"duration"`
	Actor    string `json:"actor,omitempty"`
	Director string `json:"director"`
}

func main() {
	fmt.Println("server started")
	var dbPool *pgxpool.Pool
	databaseUrl := "postgres://postgres:admin@localhost:5432/postgres"

	// this returns connection pool
	dbPool, err := pgxpool.Connect(context.Background(), databaseUrl)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	// to close DB pool
	defer dbPool.Close()
	r := mux.NewRouter()
	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the book: %s and page %s\n", title, page)
	})

	r.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		_, err := dbPool.Exec(request.Context(), insert)
		if err != nil {
			fmt.Printf("error : %v", err)
			_, err := writer.Write([]byte("error occurred"))
			if err != nil {
				return
			}
			return
		} else {
			writer.Write([]byte("inserted successfully"))
		}
	})

	http.ListenAndServe(":8080", r)
}

//type Server struct {
//	*http.Server
//	Router         *mux.Router
//}
//
//func NewServer() Server {
//	s := Server{}
//	s.Router = mux.NewRouter()
//	s.Router.PathPrefix("/").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
//		s.Router.ServeHTTP(writer, request)
//	})
//	s.Server = &http.Server{
//		Addr:           "8080",
//		Handler:        s.Router,}
//	return s
//}

//func (s Server) saveMovie(writer http.ResponseWriter, request *http.Request) {
//
//}
