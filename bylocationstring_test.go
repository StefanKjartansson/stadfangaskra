package stadfangaskra

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestByLocationString(t *testing.T) {

	control := Location{
		Street:   "Auðbrekka",
		Number:   18,
		Postcode: 200,
	}

	testCases := make(map[string]Location)
	testCases["Litla-Fjarðarhorn  510 Hólmavík"] = Location{
		Street:       "Litla-Fjarðarhorn",
		Postcode:     510,
		Municipality: "Hólmavík",
	}
	testCases["Sætúni 10, 105 Reykjavík"] = Location{
		Street:       "Sætúni",
		Number:       10,
		Postcode:     105,
		Municipality: "Reykjavík",
	}
	testCases["Fornustekkum II  781 Höfn í Hornafirði"] = Location{
		Street:       "Fornustekkum II",
		Postcode:     781,
		Municipality: "Höfn í Hornafirði",
	}
	testCases["Dunhaga 5 Tæknigarði  107 Reykjavík"] = Location{
		Street:       "Dunhaga",
		Number:       5,
		Postcode:     107,
		Municipality: "Reykjavík",
	}
	testCases["Skútuvogi 1 b  104 Reykjavík"] = Location{
		Street:       "Skútuvogi",
		Number:       1,
		NumberChars:  "b",
		Postcode:     104,
		Municipality: "Reykjavík",
	}
	testCases["Domus Medica  Egilsgötu 3  101 Reykjavík"] = Location{
		Street:       "Egilsgötu",
		Number:       3,
		Postcode:     101,
		Municipality: "Reykjavík",
	}
	testCases["Domus Medica, Egilsgötu 3  101 Reykjavík"] = Location{
		Street:       "Egilsgötu",
		Number:       3,
		Postcode:     101,
		Municipality: "Reykjavík",
	}
	testCases["Fluggörðum 30d  101 Reykjavík"] = Location{
		Street:       "Fluggörðum",
		Number:       30,
		NumberChars:  "d",
		Postcode:     101,
		Municipality: "Reykjavík",
	}
	testCases["Hafnarstræti 20 4.hæð, 101 Reykjavík"] = Location{
		Street:       "Hafnarstræti",
		Number:       20,
		Postcode:     101,
		Municipality: "Reykjavík",
	}
	testCases["Austurstræti 17 (6.h), 101 Reykjavík"] = Location{
		Street:       "Austurstræti",
		Number:       17,
		Postcode:     101,
		Municipality: "Reykjavík",
	}
	testCases["Hringbraut Landsp., 101 Reykjavík"] = Location{
		Street:       "Hringbraut Landsp.",
		Postcode:     101,
		Municipality: "Reykjavík",
	}
	testCases["Laufásvegi  12, 101 Reykjavík"] = Location{
		Street:       "Laufásvegi",
		Number:       12,
		Postcode:     101,
		Municipality: "Reykjavík",
	}
	testCases["Lindargötu Fjármálar., 150 Reykjavík"] = Location{
		Street:       "Lindargötu Fjármálar.",
		Postcode:     150,
		Municipality: "Reykjavík",
	}
	testCases["Kirkjustræti Austurv., 101 Reykjavík"] = Location{
		Street:       "Kirkjustræti Austurv.",
		Postcode:     101,
		Municipality: "Reykjavík",
	}

	for s, exp := range testCases {
		t.Logf("String: %q\n", s)
		f, err := ByLocationString(s)
		assert.Nil(t, err)
		FilterMatches(t, f, &exp, true)
		FilterMatches(t, f, &control, false)
	}

}
