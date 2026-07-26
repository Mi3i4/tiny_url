package shortener

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
)

var ErrEmptyURL = errors.New("url must not be empty")

func Shorten(rawURL string, codeLen int) (string, error) {
	if err := ValidateURL(rawURL); err != nil {
		return "", err
	}

	b := make([]byte, codeLen)
	if _, err := rand.Read(b); err != nil {
		return "", fmt.Errorf("generate code: %w", err)
	}

	return base64.RawURLEncoding.EncodeToString(b)[:codeLen], nil
}
