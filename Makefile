tidy:
	go mod tidy

build: tidy
	go build -o bin/api

run: build
	./bin/api
