package stadfangaskra

import (
	"encoding/base64"
	"fmt"
	"strings"
)

type Filter struct {
	Street   []string `schema:"street"`
	Number   []int    `schema:"number"`
	Postcode []int    `schema:"postcode"`
}

func StringMatchesWildcard(value, wildcard string) bool {
	if strings.HasSuffix(wildcard, "*") {
		v := wildcard[0:strings.Index(wildcard, "*")]
		if strings.HasPrefix(value, v) {
			return true
		}
	} else if strings.HasPrefix(wildcard, "*") {
		v := wildcard[strings.Index(wildcard, "*")+1:]
		if strings.HasSuffix(value, v) {
			return true
		}
	} else if wildcard == value {
		return true
	}
	return false
}

func StringMatchesAnyWildcard(value string, wildcardList []string) bool {
	for _, s := range wildcardList {
		if s == "*" {
			continue
		}
		if StringMatchesWildcard(value, s) {
			return true
		}
	}
	return false
}

// Returns true if a number is in the slice
func intInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func (f *Filter) Hash() string {
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%v%v%v", f.Street, f.Number, f.Postcode)))
}

func (f *Filter) Match(l *Location) bool {

	if len(f.Street) > 0 && !StringMatchesAnyWildcard(l.Street, f.Street) {
		return false
	}

	if len(f.Number) > 0 && !intInSlice(l.Number, f.Number) {
		return false
	}

	if len(f.Postcode) > 0 && !intInSlice(l.Postcode, f.Postcode) {
		return false
	}

	return true
}
