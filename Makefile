clean:
	rm -rf ./bin
	rm coverage.out

build:
	go build -o bin/jsonschema ./

clean.build: clean build

run:
	go run ./...

fmt:
	gofmt -w ./

test:
	go clean -testcache
	go test ./... -cover -coverprofile=coverage.out

test.v:
	go clean -testcache
	go test ./... -cover -coverprofile=coverage.out -v

test.cov:
	go tool cover -html=coverage.out

test.bench:
	go clean -testcache
	go test ./... -run=None -bench=. -benchmem

.PHONY: test
