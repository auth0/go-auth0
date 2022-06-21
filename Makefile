.DEFAULT_GOAL := help

.PHONY: help
help: ## Shows this help
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: deps
deps: ## Downloads dependencies
	@echo "Downloading dependencies to vendor folder..."
	@go mod vendor -v

.PHONY: lint
lint: ## Runs linting on the go files. Requires docker to be installed
	@echo "Running linting on the library..."
	@docker run --rm -v $(CURDIR):/go-auth0 -w /go-auth0 golangci/golangci-lint:v1.44.0 golangci-lint -c .golangci.yaml run ./...

.PHONY: test-e2e
test-e2e: ## Runs tests without http recordings
	@echo "Running tests..."
	@go test -cover -covermode=atomic -coverprofile=coverage.out ./...

.PHONY: test
test: ## Runs tests with http recordings
	@echo "Running tests with vcr recordings..."
	@AUTH0_HTTP_RECORDINGS=on AUTH0_DOMAIN=go-auth0-dev.eu.auth0.com go test -cover -covermode=atomic -coverprofile=coverage.out ./...

.PHONY: generate
generate: ## Generate management accessor methods
	@echo "Generating management accessor methods..."
	@go generate ./...
