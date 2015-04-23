package stadfangaskra

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestByDistance(t *testing.T) {
	// less than 5 km away
	f1, err := ByDistance(Point{X: 64.11438660967, Y: -21.7702940057837}, 0.5)
	assert.Nil(t, err)
	FilterMatches(t, f1, Stadfangaskra[0], true)
	FilterMatches(t, f1, Stadfangaskra[1], true)
	FilterMatches(t, f1, Stadfangaskra[2], true)
	_, err = ByPostCode(1)
	assert.NotNil(t, err)
}
