package stadfangaskra

import "math"

const R = 6371

func distanceInKm(p1, p2 Point) float64 {
	a := 0.5 -
		math.Cos((p2.X-p1.X)*
			math.Pi/180.0)/2 +
		math.Cos(p1.X*math.Pi/180)*
			math.Cos(p2.X*math.Pi/180)*
			(1-math.Cos((p2.Y-p1.Y)*math.Pi/180))/2
	return R * 2 * math.Asin(math.Sqrt(a))
}

func ByDistance(p Point, distance float64) (FindFilter, error) {
	return func(l *Location) bool {
		return distanceInKm(p, l.Coordinates) <= distance
	}, nil
}
