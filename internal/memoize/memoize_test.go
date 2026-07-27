package memoize

import (
	"testing"
	"time"
)

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func TestMemoize(t *testing.T) {
	tests := []struct {
		name   string
		fn     func(int) int
		input  int
		result int
	}{
		{
			name:   "double",
			fn:     func(n int) int { return n * 2 },
			input:  5,
			result: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			memo := Memoize(tt.fn)
			got := memo(tt.input)

			if got != tt.result {
				t.Errorf("got %v want %v", got, tt.result)
			}
		})
	}
}

func TestMemoizeCache(t *testing.T) {
	t.Run("cache check 1", func(t *testing.T) {
		calls := 0

		slow := func(n int) int {
			calls++
			return n * 2
		}

		memo := Memoize(slow)
		memo(5)
		memo(5)

		if calls != 1 {
			t.Errorf("no cache!")
		}
	})

	t.Run("cache check 2", func(t *testing.T) {
		calls := 0

		slow := func(n int) int {
			calls++
			return n * 2
		}

		memo := Memoize(slow)
		memo(5)
		memo(5)
		got := memo(6)
		want := 12

		if calls != 2 {
			t.Errorf("Counter broken")
		}
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestFibTiming(t *testing.T) {
	t.Run("fib test", func(t *testing.T) {
		if got := fib(10); got != 55 {
			t.Errorf("fib(10) = %d, want 55", got)
		}
	})

	t.Run("fib time", func(t *testing.T) {
		start := time.Now()
		result := fib(35)
		elapsed := time.Since(start)
		t.Logf("наивный fib(35) = %d, заняло %v", result, elapsed)
	})

	t.Run("fib memoized", func(t *testing.T) {
		memoFib := Memoize(fib)

		start := time.Now()
		memoFib(35)
		t.Logf("memoFib(35) 1-й раз: %v", time.Since(start))

		start = time.Now()
		memoFib(35)
		t.Logf("memoFib(35) 2-й раз: %v", time.Since(start))
	})

	t.Run("fib memoized recursion", func(t *testing.T) {
		var fibMemo func(int) int

		fibMemo = func(n int) int {
			if n < 2 {
				return n
			}
			return fibMemo(n-1) + fibMemo(n-2)
		}

		fibMemo = Memoize(fibMemo)

		start := time.Now()
		result := fibMemo(35)
		t.Logf("Способ 2: fib(35) = %d, 1-й раз: %v", result, time.Since(start))
	})
}
