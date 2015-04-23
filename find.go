package stadfangaskra

// change to bool, error is never used
type FindFilter func(*Location) bool

type Locations []*Location

func (locs Locations) Find(filterOpts ...FindFilter) (Locations, error) {

	matched := []*Location{}
	numFilters := len(filterOpts)

LocationLoop:
	for lidx, l := range locs {
		for fidx, f := range filterOpts {
			if !f(l) {
				continue LocationLoop
			}
			if numFilters-1 == fidx {
				matched = append(matched, locs[lidx])
			}

		}
	}
	return matched, nil
}
