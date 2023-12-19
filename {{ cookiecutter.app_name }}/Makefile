VERSION ?= $(shell ./scripts/semver current)
BUILD_COMMIT ?= $(shell git rev-parse HEAD)
BUILD_BY ?= $(shell git config user.name)
BUILD_TIME ?= $(shell date --utc)
BUILD_SYMBOL_PREFIX ?= $(shell cat go.mod | grep -e "^module" | cut -d" " -f2)/internal/build
BUILD_LDFLAGS=-ldflags "-X '$(BUILD_SYMBOL_PREFIX).Version=$(VERSION)' -X '$(BUILD_SYMBOL_PREFIX).Commit=$(BUILD_COMMIT)' -X '$(BUILD_SYMBOL_PREFIX).User=$(BUILD_BY)' -X '$(BUILD_SYMBOL_PREFIX).Time=$(BUILD_TIME)'"

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo 'Are you sure? [y/N] \c' && read ans && [ $${ans:-N} = y ]

.PHONY: check_dirty_workspace
check_dirty_workspace:
	@git diff --quiet

.PHONY: check_all
check_all: check_dirty_workspace audit test confirm

## install: bootstrap all dependencies for the project
.PHONY: install
install:
	$(if $(shell command -v asdf 2> /dev/null),,$(error ERROR: `asdf` [https://asdf-vm.com] is required))
	asdf install
	go get -u && go mod tidy
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/incu6us/goimports-reviser/v3@latest
	go install gotest.tools/gotestsum@latest
	go install mvdan.cc/gofumpt@latest

## run: run the app
.PHONY: run
run:
	go run .

## build: build the app with passed in `VERSION`, defaulting to current version
.PHONY: build
build: check_all
	go generate ./...
	@echo "Building version ${VERSION}"
	git fetch --all --tags --prune
	git checkout tags/$(VERSION)
	go build $(BUILD_LDFLAGS) -o bin/ ./...
	git checkout -

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
	goimports-reviser -rm-unused -set-alias -format ./...
	golangci-lint run

## tag-patch: create and push a git tag indicating the next patch version
.PHONY: tag-patch
tag-patch: check_all
	scripts/semver -w next-patch

## tag-minor: create and push a git tag indicating the next minor version
.PHONY: tag-minor
tag-minor: check_all
	scripts/semver -w next-minor

## tag-major: create and push a git tag indicating the next major version
.PHONY: tag-major
tag-major: check_all
	scripts/semver -w next-major

## release: create a release, using VERSION, defaulting to the current version
.PHONY: release
release: check_all
	@echo "Releasing version ${VERSION}"
	git publish
	git push --tags
