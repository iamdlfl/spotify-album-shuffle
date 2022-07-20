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

func (s server) handleHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, "This is the shuffle home!", http.StatusOK)
	}
}

func (s *server) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := make(map[string]interface{}, 0)
		data["message"] = "Hello, World!"
		data["token"] = s.token
		s.respond(w, r, data, http.StatusOK)
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

		apiUri := fmt.Sprintf("%sresponse_type=code&client_id=%s&scope=%s&redirect_uri=%s&state=%s", spotifyAuthorizeURI, clientID, scope, redirectUri, state)
		http.Redirect(w, r, apiUri, http.StatusOK)
	}
}

func (s *server) handleGetToken() http.HandlerFunc {
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

		requestBody := url.Values{}
		requestBody.Set("grant_type", "authorization_code")
		requestBody.Set("code", code[0])
		requestBody.Set("redirect_uri", redirectUri)

		encodedBody := requestBody.Encode()

		newReq, err := http.NewRequest("POST", spotifyTokenURI, strings.NewReader(encodedBody))
		if err != nil {
			log.Println(err)
			s.respond(w, r, err.Error(), http.StatusInternalServerError)
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

		if res.StatusCode >= http.StatusBadRequest {
			log.Println(string(buffer))
			info := fmt.Sprintf("Error communicating with spotify: %q", buffer)
			s.respond(w, r, info, http.StatusInternalServerError)
			return
		}

		var token tokenResponse
		json.Unmarshal(buffer, &token)

		s.token = token
		tokenLength := token.ExpirationLengthSeconds
		timeToRefresh := time.Now().Add(time.Second * time.Duration(tokenLength))
		s.timeToRefresh = timeToRefresh

		s.respond(w, r, token, http.StatusOK)
	}
}
