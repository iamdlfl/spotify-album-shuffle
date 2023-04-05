package main

import (
	crand "crypto/rand"
	"database/sql"
	_ "embed"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

var port string = "3008"

var clientID string = "a9707d9bd77e483881a10560f5bdb42d"

const lettersAndNumbers = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890"

//go:embed .secret_client
var clientSecret string

//go:embed .secret_state
var state string

var spotifyAuthorizeURI = "https://accounts.spotify.com/authorize?"
var spotifyTokenURI = "https://accounts.spotify.com/api/token"
var spotifyApiURI = "https://api.spotify.com/v1"

var redirectUri = "http://localhost:3008/token"

var myId = "onthe_dl"
var cookieStorageString = "spotify-shuffle-cookie"
var secretKey []byte = []byte{105, 109, 98, 98, 71, 54, 122, 86, 70, 56, 100, 65, 70, 78, 87, 103, 74, 50, 83, 55, 50, 69, 113, 106, 116, 114, 78, 89, 113, 111, 117, 53}

var localMode = false

var httpRoot http.FileSystem

type server struct {
	mux           *mux.Router
	token         tokenResponse
	timeToRefresh time.Time
	db            *sql.DB
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
	var err error
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	var s server
	httpRoot = http.Dir("./static/")
	s.routes()
	s.db, err = sql.Open("sqlite3", "./session.db")
	if err != nil {
		return err
	}
	defer s.db.Close()

	log.Printf("server listening at http://localhost:%s\n", port)
	return http.ListenAndServe(":"+port, s)
}

func RandByteArray(n int) []byte {
	b := make([]byte, n)
	for i := range b {
		index, _ := crand.Int(crand.Reader, big.NewInt(int64(len(lettersAndNumbers))))
		b[i] = lettersAndNumbers[index.Int64()]
	}
	return b
}
