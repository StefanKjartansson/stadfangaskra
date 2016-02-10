# stadfangaskra

stadfangaskra is a library for working with Icelandic addresses. The address data is sourced from the [national registry of Iceland](http://www.skra.is) and distributed as part of the library.

[![Build Status](https://travis-ci.org/StefanKjartansson/stadfangaskra.png?branch=develop)](https://travis-ci.org/StefanKjartansson/stadfangaskra)
[![Report Card](https://goreportcard.com/badge/github.com/StefanKjartansson/stadfangaskra)](https://goreportcard.com/badge/github.com/StefanKjartansson/stadfangaskra)

## Installation

	go get github.com/StefanKjartansson/stadfangaskra

## Usage:

Filtering the dataset is done by passing in `FindFilter` functions with the type signature `func(*stadfangaskra.Location) bool` to `Stadfangaskra.Find`.

### Example:

```go
package main

import (
	"fmt"
	"log"

	"github.com/StefanKjartansson/stadfangaskra"
)

func main() {
	// Lists even numbered addresses in Kópavogur within 1km of
	// the Landspítalinn hospital.

	// ByDistance returns a FindFilter function.
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
```

## Builtin match functions

### `ByDistance`

Returns a `FindFilter` matching addresses within the given radius in km.

```go
f, err := stadfangaskra.ByDistance(stadfangaskra.Point{X: 64.1195478, Y: -21.8807021}, 1.0)
```

### `ByLocationString`

Returns a `FindFilter` matching the criteria of the parsed input string.

```go
f, err := stadfangaskra.ByLocationString("Laufásvegur 12, 101 Reykjavík")
```

### `ByNumber`

Returns a `FindFilter` matching house numbers.

```go
f, err := stadfangaskra.ByNumber("2")
f, err := stadfangaskra.ByNumber("2a") // supports character part of house number.
```

### `ByPostCode`

Returns a `FindFilter` matching postcode.

```go
f, err := stadfangaskra.ByPostCode(200)
```

### `ByStreet`

Returns a `FindFilter` matching name of street.

```go
f, err := stadfangaskra.ByStreet("Laugavegur")
f, err := stadfangaskra.ByStreet("laugavegi") // supports dative & lowercase.
```



