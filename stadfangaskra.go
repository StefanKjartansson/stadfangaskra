package stadfangaskra

import (
	"log"
	"os"
)

var (
	DefaultStore *Store
)

// Exists reports whether the named file or directory exists.
func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func GetById(ID int) *Location {
	return DefaultStore.GetById(ID)
}

func Geocode(s string) (*Location, error) {
	return DefaultStore.FindByString(s)
}

func init() {

	path := os.Getenv("STADFANGASKRA_DB")

	if !Exists(path) {
		path = "/usr/share/stadfangaskra/db.json"
	}

	if !Exists(path) {
		path = "./db.json"
	}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	DefaultStore, err = NewStore(file)

	if err != nil {
		log.Fatal(err)
	}

}
