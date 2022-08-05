BINDIR      := $(CURDIR)/bin
BINNAME     ?= xapi-go

GOBIN         = $(shell go env GOBIN)
ifeq ($(GOBIN),)
GOBIN         = $(shell go env GOPATH)/bin
endif
GOX           = $(GOBIN)/gox
GOIMPORTS     = $(GOBIN)/goimports
AIR     	  = $(GOBIN)/air
ARCH          = $(shell uname -m)

# List all our actual files, excluding vendor
GOPKGS ?= $(shell go list $(FILES) | grep -v /vendor/)
GOFILES ?= $(shell find . -name '*.go' | grep -v /vendor/)

check-all: check-imports check-fmt check-mod ## Check all
.PHONY: check-all

fix-all: vendor tidy fiximports fmt vet ## Fix all
.PHONY: fix-all

build-linux: ## Build for linux
	@echo "==> Building"
	@go build -o $(BINDIR)/$(BINNAME) ./cmd/client
.PHONY: build

dev-dependencies: ## Downloads the necessesary dev dependencies.
	@echo "==> Downloading development dependencies"
	@go install golang.org/x/tools/cmd/goimports@latest
.PHONY: dev-dependencies

dependencies: ## Downloads modules
	@echo "==> Downloading dependencies"
	@go mod download
.PHONY: dependencies

tidy: ## Cleans the Go module.
	@echo "==> Tidying module"
	@go mod tidy
.PHONY: tidy

vendor: ## Copy of all packages needed to support builds and tests in the vendor directory
	@go mod vendor
.PHONY: vendor

check-imports: ## A check which lists improperly-formatted imports, if they exist.
	@$(shell pwd)/scripts/check-imports.sh
.PHONY: check-imports

check-fmt: ## A check which lists improperly-formatted files, if they exist.
	@$(shell pwd)/scripts/check-gofmt.sh
.PHONY: check-fmt

check-mod: ## A check which lists extraneous dependencies, if they exist.
	@$(shell pwd)/scripts/check-mod.sh
.PHONY: check-mod

fiximports: ## Properly formats and orders imports.
	@echo "==> Fixing imports"
	@$(GOIMPORTS) -d ${GOFILES}
.PHONY: fiximports

fmt: ## Properly formats Go files and orders dependencies.
	@echo "==> Running gofmt"
	@gofmt -s -w ${GOFILES}
.PHONY: fmt

vet: ## Identifies common errors.
	@echo "==> Running go vet"
	@go vet ./...
.PHONY: vet

run:
	@go run ./cmd/client
.PHONY: run

help: ## Prints this help menu.
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
.PHONY: help
