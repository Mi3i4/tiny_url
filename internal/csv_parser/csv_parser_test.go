package csvparser

import (
	"slices"
	"testing"
)

func TestParseLine(t *testing.T) {
	tests := []struct {
		name string
		line string
		want []string
	}{
		{
			name: "simple line",
			line: "a,b,c",
			want: []string{"a", "b", "c"},
		},
		{
			name: "line with quotes",
			line: `"a","b, c","d"`,
			want: []string{"a", "b, c", "d"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseLine(tt.line); !slices.Equal(got, tt.want) {
				t.Errorf("ParseLine() = %q, want %q", got, tt.want)
			}
		})
	}
}
