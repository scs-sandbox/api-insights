PACKAGES = $(shell go list ./... | grep -v /vendor/)

all: test build

.PHONY: test
test:
	go test -cover $(PACKAGES)

.PHONY: ut
ut:
	go test -cover $(PACKAGES) -coverprofile coverage.out
	go tool cover -html=coverage.out

# build the release files
.PHONY: build
build:
	go mod tidy
	mkdir -p release
	GOOS=linux   GOARCH=amd64 CGO_ENABLED=0 go build -o release/linux/amd64/api-insights   github.com/cisco-developer/api-insights/api/cmd/api-insights
	GOOS=linux   GOARCH=arm64 CGO_ENABLED=0 go build -o release/linux/arm64/api-insights   github.com/cisco-developer/api-insights/api/cmd/api-insights
	GOOS=linux   GOARCH=arm   CGO_ENABLED=0 go build -o release/linux/arm/api-insights     github.com/cisco-developer/api-insights/api/cmd/api-insights
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o release/windows/amd64/api-insights github.com/cisco-developer/api-insights/api/cmd/api-insights
	GOOS=darwin  GOARCH=amd64 CGO_ENABLED=0 go build -o release/darwin/amd64/api-insights  github.com/cisco-developer/api-insights/api/cmd/api-insights

# set git hooks for current project
set-git-hooks:
	git config core.hooksPath .githooks

# unset git hooks to use default git hooks for current project
unset-git-hooks:
	git config --unset core.hooksPath
