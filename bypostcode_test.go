package stadfangaskra

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestByPostCode(t *testing.T) {
	f1, err := ByPostCode(101)
	assert.Nil(t, err)
	FilterMatches(t, f1, &Location{Postcode: 101}, true)
	_, err = ByPostCode(1)
	assert.NotNil(t, err)
}
