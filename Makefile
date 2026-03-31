.PHONY: \
	build \
	test test-with-coverage \
	run run-json run-yaml \
	lint \
	clean 

build:
	go build -o bin/gendiff ./cmd/gendiff

test:
	go test -v ./...

test-with-coverage:
	go test -coverprofile=coverage.out ./...

run: build
	./bin/gendiff testdata/file1.json testdata/file2.yml

run-json: build
	./bin/gendiff testdata/file1.json testdata/file2.json

run-yaml: build
	./bin/gendiff testdata/file1.yml testdata/file2.yml

lint:
	golangci-lint run

clean:
	rm -rf bin/gendiff coverage.out
