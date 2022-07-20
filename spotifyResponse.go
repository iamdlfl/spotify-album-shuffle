package main

type tokenResponse struct {
	AccessToken             string `json:"access_token"`
	TokenType               string `json:"token_type"`
	Scope                   string `json:"scope"`
	ExpirationLengthSeconds int    `json:"expires_in"`
	RefreshToken            string `json:"refresh_token"`
}
