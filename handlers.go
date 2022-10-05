package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
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

func (s *server) handleIsLoggedIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := make(map[string]bool)
		data["logged_in"] = false
		if time.Now().After(s.timeToRefresh) {
			s.respond(w, r, data, http.StatusOK)
			return
		}
		// time to refresh will default to 0 time, so if user has not been logged in yet it will always return true
		data["logged_in"] = true
		s.respond(w, r, data, http.StatusOK)
	}
}

func (s server) handleLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		scope := "user-read-private user-read-email playlist-modify-private playlist-modify-public"
		apiUri := fmt.Sprintf("%sresponse_type=code&client_id=%s&scope=%s&redirect_uri=%s&state=%s&show_dialog=true", spotifyAuthorizeURI, clientID, scope, redirectUri, state)

		http.Redirect(w, r, apiUri, http.StatusSeeOther)
	}
}

func (s *server) handleLogout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.token = tokenResponse{}
		s.timeToRefresh = time.Time{}
		http.Redirect(w, r, "http://localhost:3000/", http.StatusSeeOther)
	}
}

func (s *server) handleGetToken() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := make(map[string]interface{})
		data["success"] = false

		r.ParseForm()
		code, ok1 := r.Form["code"]
		response_state, ok2 := r.Form["state"]
		if !ok1 || !ok2 {
			data["message"] = "could not get a proper response from the spotify API"
			s.respond(w, r, data, http.StatusBadRequest)
			return
		}
		if state != response_state[0] {
			data["message"] = "state was different - there may be a security issue"
			s.respond(w, r, data, http.StatusBadRequest)
			return
		}

		requestBody := url.Values{}
		requestBody.Set("grant_type", "authorization_code")
		requestBody.Set("code", code[0])
		requestBody.Set("redirect_uri", redirectUri)

		encodedBody := requestBody.Encode()

		newReq, err := http.NewRequest("POST", spotifyTokenURI, strings.NewReader(encodedBody))
		if err != nil {
			log.Println(err)
			data["message"] = err.Error()
			s.respond(w, r, data, http.StatusInternalServerError)
			return
		}

		newReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		authString := fmt.Sprintf("%s:%s", clientID, clientSecret)
		encodedAuth := base64.RawStdEncoding.EncodeToString([]byte(authString))
		newReq.Header.Set("Authorization", "Basic "+encodedAuth)

		client := &http.Client{}
		res, err := client.Do(newReq)
		if err != nil {
			log.Println(err)
			data["message"] = err.Error()
			s.respond(w, r, data, http.StatusInternalServerError)
			return
		}
		defer res.Body.Close()
		buffer, err := io.ReadAll(res.Body)
		if err != nil {
			log.Println(err)
			data["message"] = err.Error()
			s.respond(w, r, data, http.StatusInternalServerError)
			return
		}

		if res.StatusCode >= http.StatusBadRequest {
			log.Println(string(buffer))
			info := fmt.Sprintf("Error communicating with spotify: %q", buffer)
			data["message"] = info
			s.respond(w, r, data, http.StatusInternalServerError)
			return
		}

		var token tokenResponse
		json.Unmarshal(buffer, &token)

		s.token = token
		tokenLength := token.ExpirationLengthSeconds
		timeToRefresh := time.Now().Add(time.Second * time.Duration(tokenLength))
		s.timeToRefresh = timeToRefresh

		http.Redirect(w, r, "http://localhost:3008/", http.StatusSeeOther)
	}
}
