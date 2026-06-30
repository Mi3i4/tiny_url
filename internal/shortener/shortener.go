package shortener

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
)

var ErrEmptyURL = errors.New("url must not be empty")

func Shorten(url string, codeLen int) (string, error) {
	if url == "" {
		return "", ErrEmptyURL
	}

	b := make([]byte, codeLen)
	if _, err := rand.Read(b); err != nil {
		return "", fmt.Errorf("generate code: %w", err)
	}

	code := base64.RawURLEncoding.EncodeToString(b)
	return code[:codeLen], nil
}
