package itertools

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {

	t.Run("filter 1", func(t *testing.T) {
		xs := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

		want := []int{0, 2, 4, 6, 8}
		have := Filter(xs, func(x int) bool {
			return x%2 == 0
		})

		assert.Equal(t, want, have)
	})
}

func TestMap(t *testing.T) {

	t.Run("map 1", func(t *testing.T) {
		xs := []int{1, 2, 3, 4}

		want := []string{"1", "2", "3", "4"}
		have := Map(xs, func(x int) string {
			return strconv.Itoa(x)
		})

		assert.Equal(t, want, have)
	})
}

func TestDropWhile(t *testing.T) {

	t.Run("DropWhile 1", func(t *testing.T) {
		xs := []int{1, 4, 6, 4, 1}

		want := []int{6, 4, 1}
		have := DropWhile(xs, func(x int) bool {
			return x < 5
		})

		assert.Equal(t, want, have)
	})

	t.Run("DropWhile 2", func(t *testing.T) {
		xs := []int{1, 4, 6, 4, 1}

		want := []int{}
		have := DropWhile(xs, func(x int) bool {
			return x < 10
		})

		assert.Equal(t, want, have)
	})

	t.Run("DropWhile 3", func(t *testing.T) {
		xs := []int{1, 4, 6, 4, 1}

		want := xs
		have := DropWhile(xs, func(x int) bool {
			return x < 1
		})

		assert.Equal(t, want, have)
	})
}

func TestTakeWhile(t *testing.T) {

	t.Run("TakeWhile 1", func(t *testing.T) {
		xs := []int{1, 4, 6, 4, 1}

		want := []int{1, 4}
		have := TakeWhile(xs, func(x int) bool {
			return x < 5
		})

		assert.Equal(t, want, have)
	})

	t.Run("TakeWhile 2", func(t *testing.T) {
		xs := []int{1, 4, 6, 4, 1}

		want := []int{}
		have := TakeWhile(xs, func(x int) bool {
			return x < 1
		})

		assert.Equal(t, want, have)
	})

	t.Run("TakeWhile 3", func(t *testing.T) {
		xs := []int{1, 4, 6, 4, 1}

		want := xs
		have := TakeWhile(xs, func(x int) bool {
			return x < 10
		})

		assert.Equal(t, want, have)
	})
}

func TestPairwise(t *testing.T) {

	t.Run("Pairwise 1", func(t *testing.T) {
		xs := []string{"A", "B", "C", "D", "E", "F", "G"}

		want := [][2]string{
			{"A", "B"},
			{"B", "C"},
			{"C", "D"},
			{"D", "E"},
			{"E", "F"},
			{"F", "G"},
		}
		have := Pairwise(xs)

		assert.Equal(t, want, have)
	})
}
