package stadfangaskra

import (
	"testing"
)

func TestParseLocation(t *testing.T) {

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
		NumberChars:  "n",
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
	/*
	   testCases["Hafnarstræti 91-95, 600 Akureyri"] = Location{
	       Street:      "Hafnarstræti",
	       Number: "91-95",
	       Postcode:    600,
	       Municipality:       "Akureyri",
	   }
	*/
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
	/*
	   testCases["Klapparstíg 25.27, 105 Reykjavík"] = Location{
	       Street:      "Klapparstíg",
	       Number: 25-27,
	       Postcode:    105,
	       Municipality:       "Reykjavík",
	   }
	*/
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
		a, err := ParseLocation(s)
		if a.Street != exp.Street {
			t.Errorf("ParseLocation: [Street] %v, expected: %v.", a, exp)
		}
		if a.Postcode != exp.Postcode {
			t.Errorf("ParseLocation: [Postcode] %v, expected: %v.", a, exp)
		}
		if a.Municipality != exp.Municipality {
			t.Errorf("ParseLocation: %v, expected: %v.", a, exp)
		}
		if a.Number != exp.Number {
			t.Errorf("ParseLocation: [Number] %v, expected: %v.", a, exp)
		}

		if err != nil {
			t.Errorf("Error: %v parsing %s (%v).\n", err, s, a)
			return
		}
	}

}

func TestQuery(t *testing.T) {

	q, err := ParseLocation("Vatnsstígur 3b, 101 Reykjavík")

	if err != nil {
		t.Errorf("Error: %+v.\n", err)
	}

	if q.Postcode != 101 {
		t.Errorf("Error: %+v.\n", q)
	}

	loc, err := DefaultStore.FindByQuery(q)

	if err != nil {
		t.Errorf("Error: %+v.\n", err)
	}

	if loc == nil {
		t.Error("Expected address, got nil")
	}

	if loc.Street != "Vatnsstígur" {
		t.Errorf("Expected Vatnsstígur, got %+v", loc)
		t.Errorf("%+v", q.GetSearchIndex())
		x, _ := DefaultStore.SearchIndex[q.GetSearchIndex()]
		for _, l := range x {
			t.Logf("%+v %d%s", l.Street, l.Number, l.NumberChars)
		}
	}

}
