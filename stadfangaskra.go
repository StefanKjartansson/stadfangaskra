package stadfangaskra

import (
	"bytes"
	"encoding/json"
	"log"
)

var (
	Stadfangaskra = Locations{}
)

func init() {
	db, err := dbJSONBytes()
	if err != nil {
		log.Fatal(err)
	}
	decoder := json.NewDecoder(bytes.NewBuffer(db))
	err = decoder.Decode(&Stadfangaskra)
	if err != nil {
		log.Fatal(err)
	}
}
