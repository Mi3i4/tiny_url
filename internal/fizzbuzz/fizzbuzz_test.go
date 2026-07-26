package fizzbuzz

import (
	"slices"
	"testing"
)

func TestCountFizzBuzz(t *testing.T) {
	tests := []struct {
		name          string
		range_numbers []int
		rules         map[int]string
		result        []string
	}{
		{
			name:          "base one",
			range_numbers: []int{1, 2, 3, 4, 5},
			rules:         map[int]string{3: "Fizz", 5: "Buzz"},
			result:        []string{"1", "2", "Fizz", "4", "Buzz"},
		},
		{
			name:          "1..15 with 3=Fizz 5=Buzz",
			range_numbers: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			rules:         map[int]string{3: "Fizz", 5: "Buzz"},
			result: []string{
				"1", "2", "Fizz", "4", "Buzz",
				"Fizz", "7", "8", "Fizz", "Buzz",
				"11", "Fizz", "13", "14", "FizzBuzz",
			},
		},
		{
			name:          "empty rules",
			range_numbers: []int{1, 2, 3},
			rules:         map[int]string{},
			result:        []string{"1", "2", "3"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountFizzBuzz(tt.range_numbers, tt.rules)

			if !slices.Equal(got, tt.result) {
				t.Errorf("got err %v, want %v", got, tt.result)
			}
		})
	}
}
