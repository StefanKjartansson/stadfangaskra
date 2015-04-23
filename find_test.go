package stadfangaskra

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func FilterMatches(t *testing.T, f FindFilter, l *Location, expected bool) {
	assert.Equal(t, f(l), expected, fmt.Sprintf("%v", l))
}

func TestFind(t *testing.T) {

	locs := Locations{
		&Location{Postcode: 101, Street: "Laugavegur", Number: 3},
		&Location{Postcode: 101, Street: "Laugavegur", Number: 3, NumberChars: "a"},
		&Location{Postcode: 101, Street: "Laugavegur", Number: 1},
	}

	p, err := ByPostCode(101)
	assert.Nil(t, err)
	s, err := ByStreet("Laugavegur")
	assert.Nil(t, err)
	n, err := ByNumber("3")
	assert.Nil(t, err)

	filtered, err := locs.Find(p, s, n)

	assert.Nil(t, err)
	assert.Equal(t, len(filtered), 2)

}

func TestByMunicipality(t *testing.T) {
	/*
		f1, err := ByMunicipality("Selfoss")
		assert.Nil(t, err)
		FilterMatches(t, f1, &Location{Municipality: "Selfossi"}, true)

	*/
}
