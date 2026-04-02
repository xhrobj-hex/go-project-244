.PHONY: \
	build \
	test test-with-coverage \
	run run-json run-yaml \
	run-plain run-json-plain run-yaml-plain \
	run-json-json run-yaml-json \
	lint \
	clean 

build:
	go build -o bin/gendiff ./cmd/gendiff

test:
	go test -v ./...

test-with-coverage:
	go test -coverprofile=coverage.out ./...

run: build
	./bin/gendiff testdata/fixture/file5.json testdata/fixture/file6.yml

run-json: build
	./bin/gendiff testdata/fixture/file5.json testdata/fixture/file6.json

run-yaml: build
	./bin/gendiff testdata/fixture/file5.yml testdata/fixture/file6.yml
	
run-plain: build
	./bin/gendiff -f plain testdata/fixture/file5.json testdata/fixture/file6.yml

run-json-plain: build
	./bin/gendiff -f plain testdata/fixture/file5.json testdata/fixture/file6.json

run-yaml-plain: build
	./bin/gendiff -f plain testdata/fixture/file5.yml testdata/fixture/file6.yml

run-json-json: build
	./bin/gendiff -f json testdata/fixture/file5.json testdata/fixture/file6.json

run-yaml-json: build
	./bin/gendiff -f json testdata/fixture/file5.yml testdata/fixture/file6.yml

lint:
	golangci-lint run

clean:
	rm -rf bin/gendiff coverage.out
