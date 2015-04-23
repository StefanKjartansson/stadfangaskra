package stadfangaskra

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestByStreet(t *testing.T) {

	f1, err := ByStreet("Laugavegur")
	assert.Nil(t, err)
	f2, err := ByStreet("Laugavegi")
	assert.Nil(t, err)
	f3, err := ByStreet("laugavegur")
	assert.Nil(t, err)

	FilterMatches(t, f1, &Location{Street: "Laugavegur"}, true)
	FilterMatches(t, f2, &Location{StreetDative: "Laugavegi"}, true)
	FilterMatches(t, f3, &Location{StreetDative: "laugavegur"}, true)
}
