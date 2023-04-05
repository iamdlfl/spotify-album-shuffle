package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *server) routes() {
	s.mux = mux.NewRouter()
	token := s.mux.PathPrefix("/api").Subrouter()

	s.mux.Use(s.CheckToken)

	s.mux.Path("/login").Methods(http.MethodGet).Handler(s.handleLogin())
	s.mux.Path("/token").Methods(http.MethodGet).Handler(s.handleGetToken())
	s.mux.Path("/logged_in").Methods(http.MethodGet).Handler(s.handleIsLoggedIn())
	s.mux.Path("/logout").Methods(http.MethodGet).Handler(s.handleLogout())

	token.Path("/me").Methods(http.MethodGet).Handler(s.handleGetUser())
	token.Path("/playlists").Methods(http.MethodGet).Handler(s.handleGetPlaylists())
	token.Path("/shuffle/{playlist_id}").Methods(http.MethodGet).Handler(s.handleShufflePlaylist())
	token.Path("/show").Methods(http.MethodGet).Handler(s.handleGetShow())

	// Middleware
	if !localMode {
		s.mux.Use(mux.CORSMethodMiddleware(s.mux))
		s.mux.Use(CorsOriginMiddleware)
	}

	// static files
	s.mux.PathPrefix("/static/").Handler(http.FileServer(httpRoot))
	s.mux.PathPrefix("/").Handler(http.FileServer(httpRoot))

	s.mux.NotFoundHandler = s.Handle404()
}
