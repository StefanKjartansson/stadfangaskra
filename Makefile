fixture_gen = ./create_fixture/cf.go
data_path = ./create_fixture/data
package = stadfangaskra

all: fixture test

test: .PHONY
	go test -v

fixture:
	go run ${fixture_gen} ${data_path}/Stadfangaskra_20131028.dsv ${data_path}/postnumer.txt > db.json
	go-bindata -o database.go -pkg="${package}" db.json
	rm db.json

.PHONY:
