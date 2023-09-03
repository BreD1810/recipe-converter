.PHONY: build-cli install-cli test lint

build-cli:
	go build -o convert-recipe ./cmd/convert-recipe

install-cli:
	go install ./cmd/convert-recipe

test:
	go test -race -v ./...

lint:
	golangci-lint run
