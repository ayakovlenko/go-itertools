package maps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaps(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		menu := map[string]int{
			"Salad": 11,
			"Soup":  8,
			"Pasta": 13,
		}

		have := FilterEntries(menu, func(item string, price int) bool {
			return item != "Pasta" && price < 10
		})
		want := map[string]int{
			"Soup": 8,
		}

		assert.Equal(t, want, have)
	})
}
