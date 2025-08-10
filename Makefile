build:
	@go build -o bin/go-http-template cmd/main.go

run: build
	@./bin/go-http-template