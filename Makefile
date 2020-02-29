.DEFAULT_GOAL := all

download:
	@echo Download go.mod dependencies
	@go mod download

test:
	go test -v -cover  ./...
build:
	go build -o ./bin ./...
all:
	make download
	make build
	make test