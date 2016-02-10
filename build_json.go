// +build ignore

package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	floatSize = 64
)

var postCodes = make(map[int]string)

type location struct {
	Hnitnum        int    `json:"id"`
	Fasteignaheiti string `json:"name"`
	Postnr         int    `json:"postcode"`
	Husnr          int    `json:"house_number,omitempty"`
	Bokst          string `json:"house_characters,omitempty"`
	Serheiti       string `json:"specific_name,omitempty"`
	HeitiNf        string `json:"street,omitempty"`
	HeitiÞf        string `json:"street_dative,omitempty"`
	Sveitarfelag   string `json:"municipality,omitempty"`
	Coordinates    struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
	} `json:"coordinates"`
}

func importFromRecord(record []string) (loc location, err error) {

	for idx, i := range record {

		switch idx {

		case 0, 3, 4, 7, 10:
			val, err := strconv.Atoi(i)
			if err != nil {
				val = 0
			}
			switch idx {
			case 0:
				loc.Hnitnum = val
			case 7:
				loc.Postnr = val
			case 10:
				loc.Husnr = val
			}
		case 5:
			loc.Fasteignaheiti = i
		case 8:
			loc.HeitiNf = i
		case 9:
			loc.HeitiÞf = i
		case 11:
			loc.Bokst = i
		case 13:
			loc.Serheiti = i

		}
	}

	loc.Coordinates.X, _ = strconv.ParseFloat(strings.Replace(record[23], ",", ".", -1), floatSize)
	loc.Coordinates.Y, _ = strconv.ParseFloat(strings.Replace(record[24], ",", ".", -1), floatSize)

	loc.Sveitarfelag = postCodes[loc.Postnr]

	return
}

func importStream(source io.Reader) (loc []location, err error) {

	reader := csv.NewReader(source)
	reader.Comma = '|'
	_, _ = reader.Read()

	for {
		r, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return loc, err
		}
		t, err := importFromRecord(r)

		if err != nil {
			return loc, err
		}

		loc = append(loc, t)
	}

	return loc, err
}

func importDatabase(pfile string) (loc []location, err error) {
	file, err := os.Open(pfile)
	if err != nil {
		return loc, err
	}
	return importStream(file)
}

func importPostcodes(postcode_file string) error {

	file, err := os.Open(postcode_file)
	if err != nil {
		return err
	}

	reader := csv.NewReader(file)
	reader.Comma = ';'
	_, _ = reader.Read()

	for {
		r, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		code, err := strconv.Atoi(r[0])
		if err != nil {
			return err
		}
		postCodes[code] = r[1]
	}

	return nil
}

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) != 2 {
		log.Fatalf("No filename supplied")
	}

	log.Println("Starting import of postcodes")
	err := importPostcodes(args[1])

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting import")
	locations, err := importDatabase(args[0])

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Import finished")

	for idx, l := range locations {
		locations[idx].Fasteignaheiti = fmt.Sprintf("%s, %d %s", l.Fasteignaheiti, l.Postnr, l.Sveitarfelag)
	}

	b, err := json.Marshal(locations)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("JSON serialized")

	os.Stdout.Write(b)
}
