.PHONY: build run test

build:
	go build -o bin/gendiff ./cmd/gendiff

test:
	go test -v ./...

run:
	./bin/gendiff
