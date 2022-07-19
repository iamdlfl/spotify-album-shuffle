package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *server) routes() {
	s.mux = mux.NewRouter()
	s.mux.Path("/").Methods(http.MethodGet).Handler(s.handleHome())
	s.mux.Path("/shuffle").Methods(http.MethodGet).Handler(s.handleShuffle())
	s.mux.Path("/hello").Methods(http.MethodGet).Handler(s.handleHello())
	s.mux.Path("/login").Methods(http.MethodGet).Handler(s.handleLogin())
	s.mux.Path("/token").Methods(http.MethodGet).Handler(s.handleGetToken())
	s.mux.NotFoundHandler = s.Handle404()
}