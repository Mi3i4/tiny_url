package fizzbuzz

import (
	"maps"
	"slices"
	"strconv"
)

func CountFizzBuzz(nums []int, rules map[int]string) []string {
	result := make([]string, 0, len(nums))
	sortedKeys := slices.Sorted(maps.Keys(rules))

	for _, n := range nums {
		word := ""
		for _, key := range sortedKeys {
			if n%key == 0 {
				word += rules[key]
			}
		}
		if word == "" {
			word = strconv.Itoa(n)
		}
		result = append(result, word)
	}
	return result
}
