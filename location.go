package stadfangaskra

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Location struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Number       int    `json:"house_number,omitempty"`
	NumberChars  string `json:"house_characters,omitempty"`
	SpecificName string `json:"specific_name,omitempty"`
	Street       string `json:"street,omitempty"`
	StreetDative string `json:"street_dative,omitempty"`
	Postcode     int    `json:"postcode"`
	Municipality string `json:"municipality,omitempty"`
	Coordinates  Point  `json:"coordinates"`
	JSONCache    []byte `json:"-"`
}
