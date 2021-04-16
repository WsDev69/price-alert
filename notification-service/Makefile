export GO111MODULE=on

BIN_NAME := $(or $(PROJECT_NAME),'api')
PKG_PATH := $(or $(PKG),'.')
PKG_LIST := $(shell go list ${PKG_PATH}/src/... | grep -v /vendor/)

VERSION='1.0.0'
COMMIT=`git rev-parse HEAD`
BUILD_TIME=`date +%FT%T%z`

LD_FLAGS="-X ${PKG}/src/services/stats.version=${VERSION} \
 -X ${PKG}/src/services/stats.buildTime=${BUILD_TIME} \
 -X ${PKG}/src/services/stats.commit=${COMMIT}"

.PHONY: all build clean citest test cover lint

all: build lint test

dep: # Download required dependencies
	go mod vendor

lint: dep ## Lint the files local env
	cd src/ && golangci-lint run -c ../.golangci.yml && cd ../

test: dep test-db-prepare ## Run unit tests
	cd src/ && go test -race -count=1 -short ./... && cd ../

msan: ## Run memory sanitizer
	go test -msan -short ${PKG_LIST}

cover:
	go test $(shell go list ./... | grep -v /vendor/;) -cover -v

build: dep ## Build the binary file
	@cd src; \
	go build -o ../bin/${BIN_NAME} -a -tags netgo -ldflags '-w -extldflags "-static"' -ldflags ${LD_FLAGS} .

clean: ## Remove previous build
	rm -f bin/$(BIN_NAME)

run: ## run application
	go run src/main.go

fmt: ## format source files
	go fmt github.com/wsdev69/price-alert/api-service/v0.0.1/src/...

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'