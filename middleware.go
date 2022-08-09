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

func (s *server) CheckToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if s.timeToRefresh.IsZero() {
			log.Println("there is no token yet")
			next.ServeHTTP(w, r)
			return
		}
		if time.Now().After(s.timeToRefresh) {
			log.Println("token should have expired, attempting to refresh now")
			refreshToken := s.token.RefreshToken

			requestBody := url.Values{}
			requestBody.Set("grant_type", "refresh_token")
			requestBody.Set("refresh_token", refreshToken)

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
				info := fmt.Sprintf("Error communicating with spotify for refresh: %q", buffer)
				s.respond(w, r, info, http.StatusInternalServerError)
				return
			}

			var accessToken map[string]interface{}
			json.Unmarshal(buffer, &accessToken)

			tokenString, ok := accessToken["access_token"].(string)
			if !ok {
				log.Println("there was an issue getting the new token, please try to just login again")
				next.ServeHTTP(w, r)
				return
			}
			log.Println("new token: " + tokenString)

			s.token.AccessToken = tokenString
		}
		log.Println("token exists but has not expired")
		next.ServeHTTP(w, r)
	})
}

// Middleware that allows any origin and handles OPTIONS request
func CorsOriginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if r.Method == http.MethodOptions {
			return
		}
		next.ServeHTTP(w, r)
	})
}
