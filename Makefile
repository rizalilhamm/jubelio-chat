.PHONY : format install build

run:
	go run main.go

format:
	gofmt -s -w .

run-this:
	echo "hello"

everything-oke:
	go run main.go

install:
	go mod download

build:
	go build -tags musl -o main .

start:
	./main

migrate:
	go run migrations/migrate.go
