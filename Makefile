clean:
	rm -f ./bin/*
	go clean
.PHONY: clean

build:
	go -C . build -o ./bin
.PHONY: build

run:
	go run ./main.go -configFile ./config/local.yaml
.PHONY: run

generate:
	go generate ./...
.PHONY: generate

update:
	go get -u ./...
.PHONY: update

test:
	go test `go list ./... | grep -vE "./internal/api" | grep -vE "./internal/model" | grep -vE "./internal/mocks" | grep -vE "./internal/persistence/inmem"`
.PHONY: test

test-coverage:
	go test -short `go list ./... | grep -vE "./internal/api" | grep -vE "./internal/model" | grep -vE "./internal/mocks" | grep -vE "./internal/persistence/inmem"` \
 	-race -covermode=atomic -coverprofile=./bin/coverage.out
	go tool cover -func=./bin/coverage.out
	go tool cover -html=./bin/coverage.out -o ./bin/coverage.html
.PHONY: test-coverage