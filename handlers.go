package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (s *server) respond(w http.ResponseWriter, r *http.Request, data interface{}, status int) {
	if data != nil {
		jData, err := json.Marshal(data)
		if err != nil {
			log.Println(err)
			w.WriteHeader(500)
			w.Write([]byte("Internal Server Error"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		w.Write(jData)
	}
}

func (s server) Handle404() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("hit 404 handler with %s\n", r.URL.String())
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Sorry, that page could not be found"))
	}
}

func (s server) handleHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, "This is the shuffle home!", http.StatusOK)
	}
}

func (s server) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, "Hello, world!", http.StatusOK)
	}
}

func (s server) handleShuffle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, "Handler not yet implemented", http.StatusOK)
	}
}

func (s server) handleLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		scope := "user-read-private user-read-email"

		apiUri := fmt.Sprintf("%sresponse_type=code&client_id=%s&scope=%s&redirect_uri=%s&state=%s", spotifyAuthorizeURI, clientID, scope, "http://localhost:3008/token", state)
		http.Redirect(w, r, apiUri, http.StatusOK)
	}
}

func (s server) handleGetToken() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		code, ok1 := r.Form["code"]
		response_state, ok2 := r.Form["state"]
		if !ok1 || !ok2 {
			s.respond(w, r, "could not get a proper response from the spotify API", http.StatusBadRequest)
			return
		}
		if state != response_state[0] {
			s.respond(w, r, "state was different", http.StatusBadRequest)
			return
		}

		tokenUri := fmt.Sprintf("%sgrant_type=%s&code=%s&redirect_uri=%s", spotifyTokenURI, "authorization_code", code[0], "http://localhost:3008/token")
		log.Println(tokenUri)

		s.respond(w, r, "not implemented", http.StatusOK)
	}
}
