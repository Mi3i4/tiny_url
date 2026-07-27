package basics

import (
	"slices"
	"testing"
)

func TestBasics(t *testing.T) {
	t.Run("map", func(t *testing.T) {
		slice := []int{1, 2, 3}
		fn := func(n int) int { return n * 2 }

		got := Map(slice, fn)
		want := []int{2, 4, 6}

		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("filter", func(t *testing.T) {
		slice := []int{1, 2, 3}
		fn := func(n int) bool { return n%2 == 0 }

		got := Filter(slice, fn)
		want := []int{2}

		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("reduce", func(t *testing.T) {
		slice := []int{1, 2, 3, 4}
		fn := func(acc, x int) int { return acc + x }

		got := Reduce(slice, 0, fn)
		want := 10

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
