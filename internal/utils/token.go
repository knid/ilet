package utils

import (
	"errors"
	"net/http"
	"strings"
)

func ExtractTokenFromHeader(r *http.Request) (string, error) {
	token := r.Header.Get("Authorization")
	if token == "" {
		return "", errors.New("token not found")
	}

	if !strings.HasPrefix(token, "Token: ") {
		return "", errors.New("invalid token format")
	}

	splittedTokenHeader := strings.Split(token, "Token: ")
	if len(splittedTokenHeader) < 2 {
		return "", errors.New("invalid token format")
	}

	return splittedTokenHeader[1], nil
}
