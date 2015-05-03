package main

import (
	"fmt"
	"log"

	"github.com/StefanKjartansson/stadfangaskra"
)

func main() {
	// Lists even numbered addresses in Kópavogur within 1km of
	// the Landspítalinn hospital
	distance, err := stadfangaskra.ByDistance(stadfangaskra.Point{X: 64.1195478, Y: -21.8807021}, 1.0)
	if err != nil {
		log.Fatal(err)
	}
	locations, err := stadfangaskra.Stadfangaskra.Find(
		// Filter by the postcodes of Kópavogur
		func(l *stadfangaskra.Location) bool {
			return (200 <= l.Postcode && l.Postcode <= 203)
		},
		// and even numbered house numbers
		func(l *stadfangaskra.Location) bool {
			return l.Number%2 == 0
		},
		// within 1km of the hospital
		distance,
	)
	if err != nil {
		log.Fatal(err)
	}
	for _, loc := range locations {
		fmt.Printf("%s is within 1km of the hospital\n", loc.Name)
	}
}
