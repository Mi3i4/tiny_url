package textstat

import (
	"testing"
)

func TestCount(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		wantBytes     int
		wantRunes     int
		wantGraphemes int
	}{
		{
			name:          "Simple ASCII",
			input:         "hello",
			wantBytes:     5,
			wantRunes:     5,
			wantGraphemes: 5,
		},
		{
			name:          "Japanese hiragana",
			input:         "こんにちは",
			wantBytes:     15,
			wantRunes:     5,
			wantGraphemes: 5,
		},
		{
			name:          "Family emoji",
			input:         "👨‍👩‍👧‍👦",
			wantBytes:     25,
			wantRunes:     7,
			wantGraphemes: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBytes, gotRunes, gotGraphemes := Count(tt.input)
			if gotBytes != tt.wantBytes || gotRunes != tt.wantRunes || gotGraphemes != tt.wantGraphemes {
				t.Errorf("Count() = (%d, %d, %d), want (%d, %d, %d)", gotBytes, gotRunes, gotGraphemes, tt.wantBytes, tt.wantRunes, tt.wantGraphemes)
			}
		})
	}
}
