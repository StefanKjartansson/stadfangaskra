package stadfangaskra

import (
	"encoding/json"
	"fmt"
	"io"
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

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

type AddressCompound struct {
	PostCode   int
	StreetName string
}

type Store struct {
	Locations   []Location
	IdIndex     map[int]*Location
	SearchIndex map[AddressCompound][]*Location
}

func NewStore(file io.ReadCloser) (*Store, error) {

	s := Store{
		IdIndex:     make(map[int]*Location),
		SearchIndex: make(map[AddressCompound][]*Location),
	}

	decoder := json.NewDecoder(file)

	err := decoder.Decode(&s.Locations)
	if err != nil {
		return nil, err
	}

	for idx, l := range s.Locations {
		s.IdIndex[l.ID] = &s.Locations[idx]
		b, err := json.Marshal(l)
		if err != nil {
			return nil, err
		}
		s.Locations[idx].JSONCache = b

		tmp := l.GetSearchIndices()

		for _, t := range tmp {
			val, ok := s.SearchIndex[t]

			if !ok {
				val = []*Location{}
			}

			val = append(val, &s.Locations[idx])
			s.SearchIndex[t] = val

		}

	}

	return &s, nil
}

func (s *Store) FindByString(query string) (*Location, error) {

	q, err := ParseLocation(query)
	if err != nil {
		return nil, err
	}

	return s.FindByQuery(q)
}

// FindByQuery
func (s *Store) FindByQuery(query Location) (*Location, error) {

	// Check
	si := query.GetSearchIndex()
	possibles, ok := s.SearchIndex[si]
	if !ok {
		return nil, fmt.Errorf("Missing Search Index for: %+v", si)
	}

	for _, l := range possibles {
		if l.Number != query.Number {
			continue
		}
		if l.NumberChars != query.NumberChars {
			continue
		}
		return l, nil
	}

	return nil, nil

}

func (s *Store) GetById(ID int) *Location {
	val, ok := s.IdIndex[ID]
	if !ok {
		return nil
	}
	return val
}

func ParseLocation(s string) (query Location, err error) {

	pcl := rePostcode.FindStringSubmatchIndex(s)
	if pcl == nil {
		err = fmt.Errorf("No postcode found for: '%s'\n", s)
		return
	}

	if len(pcl) < 2 {
		err = fmt.Errorf("Postcode location error: '%s', %+v", s, pcl)
		return
	}

	pstart := pcl[0]
	pend := pcl[1]

	query.Postcode, err = strconv.Atoi(strings.TrimSpace(s[pstart:pend]))

	if err != nil {
		return
	}

	// Municipality follows the postcode
	query.Municipality = strings.TrimSpace(s[pend:])

	// Isolate the address part
	addressPart := strings.Trim(s[:pstart], ", ")

	// Find the house number
	anl := reNumber.FindStringSubmatchIndex(addressPart)

	// No house number, set the street and return
	if len(anl) == 0 {
		query.Street = addressPart
		return
	}

	houseNumber := addressPart[anl[0]:anl[1]]

	query.Number, err = strconv.Atoi(strings.Split(houseNumber, "-")[0])

	if err != nil {
		return
	}

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
		if query.Street != "" {
			query.Street += " "
		}
		query.Street += s
	}

	// Trim trailing spaces
	query.Street = strings.TrimSpace(query.Street)

	return
}
