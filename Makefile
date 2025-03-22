build:
	go build -o ./bin/gochain

run: build
	./bin/gochain

test:
	go test -v ./...
