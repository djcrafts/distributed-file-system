
build:
	@go build -o bin/file-system

run: build
	@./bin/file-system

test:
	go test ./... -v