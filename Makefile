clean:
	rm -f ./bin/*
	go clean
.PHONY: clean

build:
	go -C . build -o ./bin
.PHONY: build

run:
	go run ./main.go
.PHONY: run

generate:
	go generate ./...
.PHONY: generate

update:
	go get -u ./...
.PHONY: update