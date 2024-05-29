
output = cbox-calendar-puzzle

build:
	go build -o $(output)

clean:
	rm -f $(output)

test:
	go test -coverprofile cp.out ./board ./dbsqlite3 .

cover:
	go tool cover -html=cp.out
