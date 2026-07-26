package shortener

import (
	"errors"
	"net/url"
	"strings"
)

var (
	ErrInvalidURL        = errors.New("invalid url format")
	ErrUnsupportedScheme = errors.New("only http and https are supported")
	ErrURLTooLong        = errors.New("url exceeds maximum length")
)

const maxURLLength = 2048

// ValidateURL проверяет, что URL корректен и поддерживается.
func ValidateURL(raw string) error {
	if raw == "" {
		return ErrEmptyURL
	}

	if len(raw) > maxURLLength {
		return ErrURLTooLong
	}

	u, err := url.Parse(raw)
	if err != nil {
		return ErrInvalidURL
	}

	switch {
	case u.Scheme != "http" && u.Scheme != "https":
		return ErrUnsupportedScheme
	case u.Host == "":
		return ErrInvalidURL
	case strings.Contains(u.Host, " "):
		return ErrInvalidURL
	}

	return nil
}
