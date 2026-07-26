package deferdemo

import (
	"testing"
)

func TestDefer(t *testing.T) {
	tests := []struct {
		name   string
		args   string
		result string
	}{
		{
			name:   "lifo",
			args:   "",
			result: "321",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DeferLifo()

			if got != tt.result {
				t.Errorf("got result %v, want %v", got, tt.result)
			}
		})
	}

	t.Run("double", func(t *testing.T) {
		got := DeferDouble(5)
		want := 10
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("snapshot", func(t *testing.T) {
		got := DeferSnapshot()
		want := 10

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
