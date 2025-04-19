package arraysandslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("Collection of 5 ints", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("Got %d, want %d, given %v", got, want, numbers)
		}
	})

	t.Run("Collection of varying number of slices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		checkSums(t, got, want)
	})

	t.Run("Correctly sum collection of tails of varying number of slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{0, 9, 10})
		want := []int{5, 19}

		checkSums(t, got, want)
	})

	t.Run("Safely sum tails of empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9, 10})
		want := []int{0, 19}

		checkSums(t, got, want)
	})

}
