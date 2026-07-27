package basics

func Map(s []int, fn func(int) int) []int {
	result := make([]int, len(s))

	for i, v := range s {
		result[i] = fn(v)
	}
	return result
}

func Filter(s []int, fn func(int) bool) []int {
	result := make([]int, 0, len(s))

	for _, v := range s {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

func Reduce(s []int, init int, fn func(acc, x int) int) int {
	acc := init

	for _, v := range s {
		acc = fn(acc, v)
	}

	return acc
}
