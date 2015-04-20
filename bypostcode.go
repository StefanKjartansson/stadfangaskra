package stadfangaskra

import "errors"

var (
	ErrInvalidPostCodeRange = errors.New("Invalid post code range")
)

func ByPostCode(code int) (FindFilter, error) {
	if code >= 101 && code <= 902 {
		return func(l *Location) bool {
			return l.Postcode == code
		}, nil
	}
	return nil, ErrInvalidPostCodeRange
}
