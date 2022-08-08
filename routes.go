package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *server) routes() {
	s.mux = mux.NewRouter()
	token := s.mux.PathPrefix("/api").Subrouter()
	s.mux.Use(s.CheckToken)
	s.mux.Path("/").Methods(http.MethodGet).Handler(s.handleHome())
	s.mux.Path("/shuffle").Methods(http.MethodGet).Handler(s.handleShuffle())
	s.mux.Path("/hello").Methods(http.MethodGet).Handler(s.handleHello())
	s.mux.Path("/login").Methods(http.MethodGet).Handler(s.handleLogin())
	s.mux.Path("/token").Methods(http.MethodGet).Handler(s.handleGetToken())
	token.Path("/me").Methods(http.MethodGet).Handler(s.handleGetUser())
	token.Path("/playlists").Methods(http.MethodGet).Handler(s.handleGetPlaylists())
	token.Path("/shuffle/{playlist_id}").Methods(http.MethodGet).Handler(s.handleShufflePlaylist())
	s.mux.NotFoundHandler = s.Handle404()
}
