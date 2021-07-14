VERSION := 0.2

build:
	go build -o json-schema-gen

test:
	go test ./... -coverprofile=coverage.out 

coverage:
	go tool cover -html=coverage.out

lint:
	golangci-lint run