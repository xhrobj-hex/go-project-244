.PHONY: build test run clean lint

build:
	go build -o bin/gendiff ./cmd/gendiff

test:
	go test -v ./...

run: build
	./bin/gendiff testdata/file1.json testdata/file2.json

clean:
	rm -rf bin/gendiff

lint:
	golangci-lint run
