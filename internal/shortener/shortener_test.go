package shortener_test

import (
	"errors"
	"testing"

	"github.com/Mi3i4/tiny_url/internal/shortener"
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

func TestService(t *testing.T) {
	tests := []struct {
		name   string
		args   []shortener.Option
		result int
	}{
		{
			name:   "simple",
			args:   []shortener.Option{shortener.WithCodeLength(6)},
			result: 6,
		},
		{
			name:   "defaults",
			args:   nil,
			result: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := shortener.New(tt.args...).CodeLength()

			if got != tt.result {
				t.Errorf("got %v want %v", got, tt.result)
			}
		})
	}
}

func TestValidator(t *testing.T) {
	t.Run("validator calling", func(t *testing.T) {
		called := false
		custom := func(url string) error {
			called = true
			return nil
		}
		svc := shortener.New(shortener.WithValidator(custom))
		svc.Validate("https://x.com")

		if !called {
			t.Error("кастомный валидатор не был вызван")
		}
	})

	t.Run("errors validator", func(t *testing.T) {
		myErr := errors.New("bad url")
		svc := shortener.New(shortener.WithValidator(func(string) error { return myErr }))

		if err := svc.Validate("x"); !errors.Is(err, myErr) {
			t.Errorf("ожидали %v, получили %v", myErr, err)
		}
	})
}
