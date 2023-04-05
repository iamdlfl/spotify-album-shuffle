package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"net/http"
)

func writeSignedCookie(w http.ResponseWriter, cookie http.Cookie, secretKey []byte) error {
	mac := hmac.New(sha256.New, secretKey)
	mac.Write([]byte(cookie.Name))
	mac.Write([]byte(cookie.Value))
	signature := mac.Sum(nil)

	cookie.Value = string(signature) + cookie.Value

	cookie.Value = base64.URLEncoding.EncodeToString([]byte(cookie.Value))

	http.SetCookie(w, &cookie)

	return nil
}

func readSignedCookie(r *http.Request, name string, secretKey []byte) (string, error) {
	cookie, err := r.Cookie(name)
	if err != nil {
		return "", err
	}
	temp, err := base64.URLEncoding.DecodeString(cookie.Value)
	signedValue := string(temp)
	if err != nil {
		return "", errors.New("invalid value")
	}

	if len(signedValue) < sha256.Size {
		return "", errors.New("invalid value")
	}

	signature := signedValue[:sha256.Size]
	value := signedValue[sha256.Size:]

	mac := hmac.New(sha256.New, secretKey)
	mac.Write([]byte(name))
	mac.Write([]byte(value))
	expectedSignature := mac.Sum(nil)

	if !hmac.Equal([]byte(signature), expectedSignature) {
		return "", errors.New("invalid value")
	}

	return value, nil
}
