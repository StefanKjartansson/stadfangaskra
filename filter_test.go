package stadfangaskra

import "testing"

func TestFilter(t *testing.T) {

	f := Filter{
		Street: []string{"*vegur", "*gata"},
	}

	l := Location{
		Street: "Laugavegur",
	}

	if !f.Match(&l) {
		t.Fatalf("Filter %v should match location %+v", f, l)
	}

	l.Postcode = 101
	f.Postcode = []int{200}

	if f.Match(&l) {
		t.Fatalf("Filter %v should not match location %+v", f, l)
	}

}

func TestFilterHash(t *testing.T) {

	f := Filter{
		Street:   []string{"*vegur"},
		Postcode: []int{101, 200},
	}
	f2 := Filter{
		Street:   []string{"*vegur"},
		Postcode: []int{101, 200},
	}

	if f.Hash() != f2.Hash() {
		t.Fatalf("f1: %+v should be equal to f2: %+v", f, f2)
	}

}
