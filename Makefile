fixture_gen = ./create_fixture/cf.go
data_path = ./create_fixture/data

all: fixture test

test: .PHONY
	go test -v

fixture:
	go run ${fixture_gen} ${data_path}/Stadfangaskra_20130326.dsv ${data_path}/postnumer.txt > db.json

.PHONY:
