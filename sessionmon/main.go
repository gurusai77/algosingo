package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

	req.Context()
	fmt.Println("server: handler started")
	defer fmt.Println("server: handler ended")

	fmt.Fprintf(w, "hello\n")

}

func main() {

	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}
