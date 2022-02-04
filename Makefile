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
	@docker run --rm -v $(CURDIR):/go-auth0 -w /go-auth0 golangci/golangci-lint:v1.41.1 golangci-lint run ./...

.PHONY: tests
tests: ## Runs tests
	@echo "Running tests..."
	@go test ./...
