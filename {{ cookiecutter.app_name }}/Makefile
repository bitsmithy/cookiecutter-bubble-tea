## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

## install: bootstrap all dependencies for the project
.PHONY: install
install:
	asdf install
	go install honnef.co/go/tools/cmd/staticcheck@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install gotest.tools/gotestsum@latest

## run: run the app
.PHONY: run
run:
	go run .

## build: build the app
.PHONY: build
build:
	go build -o bin/{{ cookiecutter.app_name }} .

## clean: cleanup build files
.PHONY: clean
clean:
	rm -f bin/*

## dependencies: install all code dependencies
.PHONY: dependencies
dependencies:
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify

## test: run all tests
.PHONY: test
test: dependencies
	@echo 'Running tests...'
	go clean -testcache
	gotestsum --format dots-v2 -- -race -vet=off ./...

## audit: tidy dependencies and format, and vet all code
.PHONY: audit
audit: dependencies
	@echo 'Auditing code...'
	golangci-lint run
