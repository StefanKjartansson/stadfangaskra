package stadfangaskra

import (
	"errors"
	"strconv"
	"strings"
)

var (
	ErrInvalidHouseNumber = errors.New("Invalid house number")
)

func isDigit(r rune) bool {
	if n := r - '0'; n >= 0 && n <= 9 {
		return false
	}
	return true
}

// ByNumber returns a FindFilter matching house numbers.
func ByNumber(number string) (FindFilter, error) {

	numPart := strings.TrimFunc(number, isDigit)
	charPart := strings.TrimFunc(number, func(r rune) bool { return !isDigit(r) })

	n, err := strconv.Atoi(numPart)
	if err != nil {
		return nil, err
	}

	if n <= 0 && charPart == "" {
		return nil, ErrInvalidHouseNumber
	}

	return func(l *Location) bool {
		if charPart == "" {
			return l.Number == n
		}
		return l.Number == n && l.NumberChars == charPart
	}, nil
}
