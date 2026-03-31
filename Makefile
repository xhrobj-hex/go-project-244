.PHONY: \
	build \
	test test-with-coverage \
	run \
	lint \
	clean 

build:
	go build -o bin/gendiff ./cmd/gendiff

test:
	go test -v ./...

test-with-coverage:
	go test -coverprofile=coverage.out ./...

run: build
	./bin/gendiff testdata/file1.json testdata/file2.json

lint:
	golangci-lint run

clean:
	rm -rf bin/gendiff coverage.out
