.PHONY: build run

build:
	go build -o bin/gendiff ./cmd/gendiff

run:
	./bin/gendiff
