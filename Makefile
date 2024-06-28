build:
	@go build -o bin/gosh cmd/gosh/main.go

run: build
	@./bin/gosh

test:
	@go test -v ./...
