
output = cbox-calendar-puzzle

.PHONY: build clean test cover

build:
	go build -o $(output)

clean:
	rm -f $(output)

test:
	go test -coverprofile cp.out ./board ./dbsqlite3 .

cover:
	go tool cover -html=cp.out
