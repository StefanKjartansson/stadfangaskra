package stadfangaskra

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestByNumber(t *testing.T) {
	f1, err := ByNumber("2")
	assert.Nil(t, err)
	f2, err := ByNumber("2a")
	assert.Nil(t, err)
	FilterMatches(t, f1, &Location{Number: 2}, true)
	FilterMatches(t, f2, &Location{Number: 2, NumberChars: "a"}, true)
}
