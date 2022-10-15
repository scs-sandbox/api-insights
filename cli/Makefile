PACKAGES = $(shell go list ./... | grep -v /vendor/)
CLI_NAME = "api-insights-cli"

all: test

test:
	go test -cover -v $(PACKAGES)

# build cli binaries
.PHONY: bin
bin:
	mkdir -p bin
	GOOS=linux   GOARCH=amd64 CGO_ENABLED=0 go build -o bin/linux/amd64/$(CLI_NAME)   main.go
	GOOS=linux   GOARCH=arm64 CGO_ENABLED=0 go build -o bin/linux/arm64/$(CLI_NAME)   main.go
	GOOS=linux   GOARCH=arm   CGO_ENABLED=0 go build -o bin/linux/arm/$(CLI_NAME)     main.go
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o bin/windows/amd64/$(CLI_NAME) main.go
	GOOS=darwin  GOARCH=amd64 CGO_ENABLED=0 go build -o bin/darwin/amd64/$(CLI_NAME)  main.go
	GOOS=darwin  GOARCH=arm64 CGO_ENABLED=0 go build -o bin/darwin/arm64/$(CLI_NAME)  main.go

# generate docs for CLI
.PHONY: docgen
docgen:
	mkdir -p docs
	go run main.go docs