all: build

build:
	go mod tidy
	go vet ./...
	make test
	go build

test:
	go test -race -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
	rm -f coverage.out

clean:
	go clean