.PHONY: build test run clean

build:
	go build -o bin/gendiff ./cmd/gendiff

test:
	go test -v ./...

run: build
	./bin/gendiff testdata/file1.json

clean:
	rm -rf bin/gendiff
