package main

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var port string = "3008"

var clientID string = "a9707d9bd77e483881a10560f5bdb42d"

//go:embed .secret
var clientSecret string

var spotifyAuthorizeURI = "https://accounts.spotify.com/authorize?"
var spotifyTokenURI = "https://accounts.spotify.com/api/token?"

var state = "ajd93hflahe93lhf"

type server struct {
	mux *mux.Router
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
	s.routes()

	log.Printf("server listening at http://localhost:%s\n", port)
	return http.ListenAndServe(":"+port, s)
}
