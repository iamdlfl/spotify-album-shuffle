package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

func (s *server) handleShufflePlaylist() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//check mux vars
		vars := mux.Vars(r)
		playlistID, ok := vars["playlist_id"]
		if !ok {
			s.respond(w, r, "Error while finding necessary variables. Please contact the development team", http.StatusBadRequest)
			return
		}
		playlistUri := spotifyApiURI + "/playlists/" + playlistID
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

		jsonResponse, err := UnmarshalPlaylist(buffer)
		if err != nil {
			log.Println(err)
			s.respond(w, r, err.Error(), http.StatusInternalServerError)
			return
		}

		firstTrackAlbumId := jsonResponse.Tracks.Items[0].Track.Album.ID
		var albumNumberOfTracks int = 0
		for _, track := range jsonResponse.Tracks.Items {
			if track.Track.Album.ID == firstTrackAlbumId {
				albumNumberOfTracks++
			} else {
				break
			}
		}

		shuffleUri := spotifyApiURI + "/playlists/" + playlistID + "/tracks"
		createBody := map[string]interface{}{
			"range_start":   0,
			"insert_before": jsonResponse.Tracks.Total,
			"range_length":  albumNumberOfTracks,
		}
		jsonBody, _ := json.Marshal(createBody)
		body := bytes.NewReader(jsonBody)
		putReq, err := http.NewRequest("PUT", shuffleUri, body)
		if err != nil {
			log.Println(err)
			s.respond(w, r, err.Error(), http.StatusInternalServerError)
			return
		}
		putReq.Header.Set("Authorization", "Bearer "+s.token.AccessToken)
		putReq.Header.Set("Content-Type", "application/json")
		client2 := &http.Client{}
		res2, err := client2.Do(putReq)
		if err != nil {
			log.Println(err)
			s.respond(w, r, err.Error(), http.StatusInternalServerError)
			return
		}
		defer res2.Body.Close()
		buffer2, err := io.ReadAll(res2.Body)
		if err != nil {
			log.Println(err)
			s.respond(w, r, err.Error(), http.StatusInternalServerError)
			return
		}

		var jsonResponse2 map[string]interface{}
		err = json.Unmarshal(buffer2, &jsonResponse2)
		if err != nil {
			log.Println(err)
			s.respond(w, r, err.Error(), http.StatusInternalServerError)
			return
		}
		s.respond(w, r, jsonResponse2, http.StatusOK)
	}
}

func (s *server) handleGetShow() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		allEpisodes := make([]Item, 0)
		id := "04re9FvL1xviHGWvXKoAhZ"
		showUri := spotifyApiURI + "/shows/" + id + "/episodes"
		newReq, err := http.NewRequest("GET", showUri, nil)
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

		var jsonResponse EpisodeList
		err = json.Unmarshal(buffer, &jsonResponse)
		if err != nil {
			log.Println(err)
			s.respond(w, r, err.Error(), http.StatusInternalServerError)
			return
		}
		items := jsonResponse.Items
		allEpisodes = append(allEpisodes, items...)
		limit := 20
		lastLink := ""
		for jsonResponse.Next != "" && jsonResponse.Next != lastLink && limit > 0 {
			lastLink = jsonResponse.Next
			newReq, err := http.NewRequest("GET", jsonResponse.Next, nil)
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
			err = json.Unmarshal(buffer, &jsonResponse)
			if err != nil {
				log.Println(err)
				s.respond(w, r, err.Error(), http.StatusInternalServerError)
				return
			}
			limit--
			items := jsonResponse.Items
			allEpisodes = append(allEpisodes, items...)
		}

		s.respond(w, r, allEpisodes, http.StatusOK)
	}
}
