package main

import "encoding/json"

type tokenResponse struct {
	AccessToken             string `json:"access_token"`
	TokenType               string `json:"token_type"`
	Scope                   string `json:"scope"`
	ExpirationLengthSeconds int    `json:"expires_in"`
	RefreshToken            string `json:"refresh_token"`
}

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    playlist, err := UnmarshalPlaylist(bytes)
//    bytes, err = playlist.Marshal()

func UnmarshalPlaylist(data []byte) (Playlist, error) {
	var r Playlist
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Playlist) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Playlist struct {
	Collaborative bool         `json:"collaborative"`
	Description   string       `json:"description"`
	ExternalUrls  ExternalUrls `json:"external_urls"`
	Followers     Followers    `json:"followers"`
	Href          string       `json:"href"`
	ID            string       `json:"id"`
	Images        []Image      `json:"images"`
	Name          string       `json:"name"`
	Owner         Owner        `json:"owner"`
	PrimaryColor  interface{}  `json:"primary_color"`
	Public        bool         `json:"public"`
	SnapshotID    string       `json:"snapshot_id"`
	Tracks        Tracks       `json:"tracks"`
	Type          string       `json:"type"`
	URI           string       `json:"uri"`
}

type ExternalUrls struct {
	Spotify string `json:"spotify"`
}

type Followers struct {
	Href  interface{} `json:"href"`
	Total int64       `json:"total"`
}

type Image struct {
	Height int64  `json:"height"`
	URL    string `json:"url"`
	Width  int64  `json:"width"`
}

type Owner struct {
	DisplayName  *string      `json:"display_name,omitempty"`
	ExternalUrls ExternalUrls `json:"external_urls"`
	Href         string       `json:"href"`
	ID           string       `json:"id"`
	Type         OwnerType    `json:"type"`
	URI          string       `json:"uri"`
	Name         *string      `json:"name,omitempty"`
}

type Tracks struct {
	Href     string      `json:"href"`
	Items    []TrackItem `json:"items"`
	Limit    int64       `json:"limit"`
	Next     string      `json:"next"`
	Offset   int64       `json:"offset"`
	Previous interface{} `json:"previous"`
	Total    int64       `json:"total"`
}

type TrackItem struct {
	AddedAt        string         `json:"added_at"`
	AddedBy        Owner          `json:"added_by"`
	IsLocal        bool           `json:"is_local"`
	PrimaryColor   interface{}    `json:"primary_color"`
	Track          TrackClass     `json:"track"`
	VideoThumbnail VideoThumbnail `json:"video_thumbnail"`
}

type TrackClass struct {
	Album            AlbumClass   `json:"album"`
	Artists          []Owner      `json:"artists"`
	AvailableMarkets []string     `json:"available_markets"`
	DiscNumber       int64        `json:"disc_number"`
	DurationMS       int64        `json:"duration_ms"`
	Episode          bool         `json:"episode"`
	Explicit         bool         `json:"explicit"`
	ExternalIDS      ExternalIDS  `json:"external_ids"`
	ExternalUrls     ExternalUrls `json:"external_urls"`
	Href             string       `json:"href"`
	ID               string       `json:"id"`
	IsLocal          bool         `json:"is_local"`
	Name             string       `json:"name"`
	Popularity       int64        `json:"popularity"`
	PreviewURL       string       `json:"preview_url"`
	Track            bool         `json:"track"`
	TrackNumber      int64        `json:"track_number"`
	Type             TrackType    `json:"type"`
	URI              string       `json:"uri"`
}

type AlbumClass struct {
	AlbumType            AlbumTypeEnum        `json:"album_type"`
	Artists              []Owner              `json:"artists"`
	AvailableMarkets     []string             `json:"available_markets"`
	ExternalUrls         ExternalUrls         `json:"external_urls"`
	Href                 string               `json:"href"`
	ID                   string               `json:"id"`
	Images               []Image              `json:"images"`
	Name                 string               `json:"name"`
	ReleaseDate          string               `json:"release_date"`
	ReleaseDatePrecision ReleaseDatePrecision `json:"release_date_precision"`
	TotalTracks          int64                `json:"total_tracks"`
	Type                 AlbumTypeEnum        `json:"type"`
	URI                  string               `json:"uri"`
}

type ExternalIDS struct {
	Isrc string `json:"isrc"`
}

type VideoThumbnail struct {
	URL interface{} `json:"url"`
}

type OwnerType string

const (
	Artist OwnerType = "artist"
	User   OwnerType = "user"
)

type AlbumTypeEnum string

const (
	Album  AlbumTypeEnum = "album"
	Single AlbumTypeEnum = "single"
)

type ReleaseDatePrecision string

const (
	Day ReleaseDatePrecision = "day"
)

type TrackType string

const (
	Track TrackType = "track"
)

type Show struct {
	AvailableMarkets   []string     `json:"available_markets"`
	Copyrights         []Copyright  `json:"copyrights"`
	Description        string       `json:"description"`
	HTMLDescription    string       `json:"html_description"`
	Explicit           bool         `json:"explicit"`
	ExternalUrls       ExternalUrls `json:"external_urls"`
	Href               string       `json:"href"`
	ID                 string       `json:"id"`
	Images             []Image      `json:"images"`
	IsExternallyHosted bool         `json:"is_externally_hosted"`
	Languages          []string     `json:"languages"`
	MediaType          string       `json:"media_type"`
	Name               string       `json:"name"`
	Publisher          string       `json:"publisher"`
	Type               string       `json:"type"`
	URI                string       `json:"uri"`
	TotalEpisodes      int64        `json:"total_episodes"`
	Episodes           Episodes     `json:"episodes"`
}

type Copyright struct {
	Text string `json:"text"`
	Type string `json:"type"`
}

type Episodes struct {
	Href     string `json:"href"`
	Limit    int64  `json:"limit"`
	Next     string `json:"next"`
	Offset   int64  `json:"offset"`
	Previous string `json:"previous"`
	Total    int64  `json:"total"`
	Items    []Item `json:"items"`
}

type Item struct {
	AudioPreviewURL      string       `json:"audio_preview_url"`
	Description          string       `json:"description"`
	HTMLDescription      string       `json:"html_description"`
	DurationMS           int64        `json:"duration_ms"`
	Explicit             bool         `json:"explicit"`
	ExternalUrls         ExternalUrls `json:"external_urls"`
	Href                 string       `json:"href"`
	ID                   string       `json:"id"`
	Images               []Image      `json:"images"`
	IsExternallyHosted   bool         `json:"is_externally_hosted"`
	IsPlayable           bool         `json:"is_playable"`
	Language             string       `json:"language"`
	Languages            []string     `json:"languages"`
	Name                 string       `json:"name"`
	ReleaseDate          string       `json:"release_date"`
	ReleaseDatePrecision string       `json:"release_date_precision"`
	ResumePoint          ResumePoint  `json:"resume_point"`
	Type                 string       `json:"type"`
	URI                  string       `json:"uri"`
	Restrictions         Restrictions `json:"restrictions"`
}
type Restrictions struct {
	Reason string `json:"reason"`
}

type ResumePoint struct {
	FullyPlayed      bool  `json:"fully_played"`
	ResumePositionMS int64 `json:"resume_position_ms"`
}

type EpisodeList struct {
	Href     string `json:"href"`
	Limit    int64  `json:"limit"`
	Next     string `json:"next"`
	Offset   int64  `json:"offset"`
	Previous string `json:"previous"`
	Total    int64  `json:"total"`
	Items    []Item `json:"items"`
}
