package memoize

func Memoize(f func(int) int) func(int) int {
	cache := map[int]int{}

	return func(n int) int {
		if v, ok := cache[n]; ok {
			return v
		}

		v := f(n)
		cache[n] = v
		return v
	}
}
