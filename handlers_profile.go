package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func (s *server) handleGetUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		meUri := spotifyApiURI + "/me"
		newReq, err := http.NewRequest("GET", meUri, nil)
		if err != nil {
			log.Println(err)
			s.respond(w, r, err.Error(), http.StatusInternalServerError)
			return
		}
		newReq.Header.Set("Authorization", "Bearer "+s.token.AccessToken)
		client := &http.Client{}
		res, err := client.Do(newReq)
		if err != nil {
			log.Println(err)
			s.respond(w, r, err.Error(), http.StatusInternalServerError)
			return
		}
		defer res.Body.Close()

		buffer, err := io.ReadAll(res.Body)
		if err != nil {
			log.Println(err)
			s.respond(w, r, err.Error(), http.StatusInternalServerError)
			return
		}

		var jsonResponse map[string]interface{}
		err = json.Unmarshal(buffer, &jsonResponse)
		if err != nil {
			log.Println(err)
			s.respond(w, r, err.Error(), http.StatusInternalServerError)
			return
		}

		s.respond(w, r, jsonResponse, http.StatusOK)
	}
}

func (s *server) handleGetPlaylists() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		playlistUri := spotifyApiURI + "/users/" + myId + "/playlists"
		newReq, err := http.NewRequest("GET", playlistUri, nil)
		if err != nil {
			log.Println(err)
			s.respond(w, r, err.Error(), http.StatusInternalServerError)
			return
		}
		newReq.Header.Set("Authorization", "Bearer "+s.token.AccessToken)
		client := &http.Client{}
		res, err := client.Do(newReq)
		if err != nil {
			log.Println(err)
			s.respond(w, r, err.Error(), http.StatusInternalServerError)
			return
		}
		defer res.Body.Close()

		buffer, err := io.ReadAll(res.Body)
		if err != nil {
			log.Println(err)
			s.respond(w, r, err.Error(), http.StatusInternalServerError)
			return
		}

		var jsonResponse map[string]interface{}
		err = json.Unmarshal(buffer, &jsonResponse)
		if err != nil {
			log.Println(err)
			s.respond(w, r, err.Error(), http.StatusInternalServerError)
			return
		}

		s.respond(w, r, jsonResponse, http.StatusOK)
	}
}
