gen = ./build_json.go
data_path = ./data
package = stadfangaskra
outfile = database.go

all: test

test: .PHONY
	go test -v

fixture:
	go run ${gen} ${data_path}/stadfang.dsv ${data_path}/postnumer.txt > db.json
	go-bindata -o ${outfile} -pkg="${package}" db.json
	sed -i -e 's/bindata_read/bindataRead/g' ${outfile}
	sed -i -e 's/bindata_file_info/bindataFileInfo /g' ${outfile}
	sed -i -e 's/_db_json/_dbJSON/g' ${outfile}
	sed -i -e 's/db_json_bytes/dbJSONBytes/g' ${outfile}
	sed -i -e 's/db_json/dbJSON/g' ${outfile}
	sed -i -e 's/_bintree_t/_bintreeT/g' ${outfile}
	gofmt -s -w ${outfile}
	rm ${outfile}-e
	rm db.json

.PHONY:
