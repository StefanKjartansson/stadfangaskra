package stadfangaskra

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	reNumber       = regexp.MustCompile(`\d+-?\.?`)
	rePostcode     = regexp.MustCompile(`\d{3}\s+`)
	reRemainder    = regexp.MustCompile(`^[a-zA-Z]{1}$`)
	reStrictNumber = regexp.MustCompile(`^\d+$`)
	excemptionList = []string{
		"Domus",
		"Medica",
	}
)

func ByLocationString(s string) (FindFilter, error) {

	pcl := rePostcode.FindStringSubmatchIndex(s)
	if pcl == nil {
		return nil, fmt.Errorf("No postcode found for: '%s'\n", s)
	}

	if len(pcl) < 2 {
		return nil, fmt.Errorf("Postcode location error: '%s', %+v", s, pcl)
	}

	pstart := pcl[0]
	pend := pcl[1]

	postcode, err := strconv.Atoi(strings.TrimSpace(s[pstart:pend]))

	if err != nil {
		return nil, err
	}

	postcodeFilter, err := ByPostCode(postcode)
	if err != nil {
		return nil, err
	}

	// Municipality follows the postcode
	municipality := strings.TrimSpace(s[pend:])

	// Isolate the address part
	addressPart := strings.Trim(s[:pstart], ", ")

	// Find the house number
	anl := reNumber.FindStringSubmatchIndex(addressPart)

	// No house number, set the street and return
	if len(anl) == 0 {
		streetFilter, err := ByStreet(addressPart)
		if err != nil {
			return nil, err
		}
		return func(l *Location) bool {
			// TODO, municipality filter
			if municipality != "" {
				if l.Municipality != municipality {
					return false
				}
			}
			return postcodeFilter(l) && streetFilter(l)
		}, nil
	}

	houseNumber := addressPart[anl[0]:anl[1]]

	// The address part trailing the number is larger than the capturing regex,
	// this indicates that there's either a house character in the housenumber
	// or a range of building numbers
	if len(addressPart) > anl[1] {

		remainder := strings.TrimSpace(addressPart[anl[1]:])

		// We only care about trailing house characters and building ranges
		if reRemainder.MatchString(remainder) || reStrictNumber.MatchString(remainder) {
			houseNumber += remainder
		}

		// Some building ranges are delimited by a dot, replace with a dash
		houseNumber = strings.Replace(houseNumber, ".", "-", -1)
	}

	var numberFilter FindFilter
	if houseNumber != "" {
		numberFilter, err = ByNumber(houseNumber)
		if err != nil {
			return nil, err
		}
	}

	street := ""

	//Find first number from the house number
	//Find the first character from the house number

	// Street name part, usually there is just a single name but in some cases
	// this part is a place name (not unusual to encounter farm names here).
	for _, s := range strings.Split(strings.TrimSpace(addressPart[:anl[0]]), " ") {

		s = strings.Trim(s, ", ")

		// Ignore empty strings and excempt strings
		// TODO: Expand excemption list to return f.i. the address of mall instead of
		// it's print name.
		if s == "" || stringInSlice(s, excemptionList) {
			continue
		}

		// Add space if there are more than one parts
		if street != "" {
			street += " "
		}
		street += s
	}

	streetFilter, err := ByStreet(strings.TrimSpace(street))
	if err != nil {
		return nil, err
	}

	return func(l *Location) bool {
		// TODO, municipality filter
		if municipality != "" {
			if l.Municipality != municipality {
				return false
			}
		}
		m := postcodeFilter(l) && streetFilter(l)
		if !m {
			return m
		}
		if numberFilter != nil {
			return numberFilter(l)
		}
		return true
	}, nil

}
