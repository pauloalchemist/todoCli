include .env 
export

build:
	go build -o bin/${BINARY_NAME} ./cmd/cli/main.go

run-list:
	go run ./cmd/cli/main.go -list

run-add:
	go run ./cmd/cli/main.go -add

clean:
	go clean
	rm -rf ./bin/todo

test:
	go test -v ./...

test_coverage:
	go test -v ./... -coverprofile=coverage.out
