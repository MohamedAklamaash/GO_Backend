run:
	@go run cmd/main.go

build:
	@go build -o bin/auth cmd/main.go

buildrun: build
	@./bin/auth

test:
	@go test -v ./...