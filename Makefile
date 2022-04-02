test:
  go mod download
  go mod tidy
	go test ./...
record-test-coverage:
  go mod download
  go mod tidy
	go test ./... -v -covermode=count -coverprofile=coverage.out
build:
	go build -o ./pwd-gen .
version: build
	./pwd-gen version 
