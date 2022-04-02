test:
	go test ./...
record-test-coverage:
	go test ./... -v -covermode=count -coverprofile=coverage.out
build:
	go build -o ./pwd-gen .
