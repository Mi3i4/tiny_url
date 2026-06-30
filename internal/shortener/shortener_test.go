package shortener_test

import (
	"errors"
	"github.com/Mi3i4/tiny_url/internal/shortener"
	"testing"
)

func TestShorten(t *testing.T) {
	tests := []struct {
		name    string
		url     string
		codeLen int
		wantErr error
	}{
		{
			name:    "valid url",
			url:     "https://yandex.ru",
			codeLen: 8,
			wantErr: nil,
		},
		{
			name:    "empty url",
			url:     "",
			codeLen: 8,
			wantErr: shortener.ErrEmptyURL,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code, err := shortener.Shorten(tt.url, tt.codeLen)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("got err %v, want %v", err, tt.wantErr)
			}

			if err == nil && len(code) != tt.codeLen {
				t.Errorf("got code len %d, want %d", len(code), tt.codeLen)
			}
		})
	}
}
