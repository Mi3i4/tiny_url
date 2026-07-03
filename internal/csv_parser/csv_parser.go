package csvparser

import (
	"strings"
)

func ParseLine(s string) []string {
	result := []string{}
	inQuotes := false
	sb := strings.Builder{}
	for _, r := range s {
		if r == '"' {
			inQuotes = !inQuotes
		} else if !inQuotes && r == ',' {
			result = append(result, sb.String())
			sb.Reset()
		} else {
			sb.WriteRune(r)
		}
	}
	result = append(result, sb.String())
	return result
}
