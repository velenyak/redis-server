BINARY_NAME=redis-server

build:
	go build -o bin/$(BINARY_NAME) -v ./cmd/$(BINARY_NAME)
tidy:
	go mod tidy
