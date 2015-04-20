package stadfangaskra

import "strings"

func ByStreet(name string) (FindFilter, error) {

	// TODO error if not in street/street-dative list

	return func(l *Location) bool {
		return l.Street == name ||
			l.StreetDative == name ||
			strings.ToLower(l.Street) == name ||
			strings.ToLower(l.StreetDative) == name
	}, nil
}
