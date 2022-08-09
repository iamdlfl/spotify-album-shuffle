package main

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

var port string = "3008"

var clientID string = "a9707d9bd77e483881a10560f5bdb42d"

//go:embed .secret
var clientSecret string

var spotifyAuthorizeURI = "https://accounts.spotify.com/authorize?"
var spotifyTokenURI = "https://accounts.spotify.com/api/token"
var spotifyApiURI = "https://api.spotify.com/v1"

var state = "ajd93hflahe93lhf"
var redirectUri = "http://localhost:3008/token"

var myId = "onthe_dl"

var httpRoot http.FileSystem

type server struct {
	mux           *mux.Router
	token         tokenResponse
	timeToRefresh time.Time
}

func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	var s server
	httpRoot = http.Dir("./static/")
	s.routes()

	log.Printf("server listening at http://localhost:%s\n", port)
	return http.ListenAndServe(":"+port, s)
}
