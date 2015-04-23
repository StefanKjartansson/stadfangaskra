gen = ./build_json.go
data_path = ./data
package = stadfangaskra

all: test

test: .PHONY
	go test -v

fixture:
	go run ${gen} ${data_path}/stadfang.dsv ${data_path}/postnumer.txt > db.json
	go-bindata -o database.go -pkg="${package}" db.json
	rm db.json

.PHONY:
