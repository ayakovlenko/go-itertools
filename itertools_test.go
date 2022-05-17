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
