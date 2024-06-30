package api

import (
	"context"
	"github.com/gorilla/mux"
	"net"
	"net/http"
)

type Server struct {
	*http.Server

	Router   *mux.Router // add your API routes here
	Listener net.Listener
}

func New() {
	s := Server{}
	s.Router = mux.NewRouter()

	s.route()

	s.Server = &http.Server{
		Addr:    ":8080",
		Handler: &s,
	}
	err := s.ListenAndServe()
	if err != nil {
		return
	}
}

// ServeHTTP implements http.Handler
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	s.Router.ServeHTTP(w, r.WithContext(ctx))
}

func (s *Server) route() {
	s.Router.Handle("/test", http.HandlerFunc(s.testHandler))
}

func (s *Server) testHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test"))
}
