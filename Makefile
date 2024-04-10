#-----------------------------------------------------------------------------------------------------------------------
# Variables (https://www.gnu.org/software/make/manual/html_node/Using-Variables.html#Using-Variables)
#-----------------------------------------------------------------------------------------------------------------------
.DEFAULT_GOAL := help
GO_BIN ?= $(shell go env GOPATH)/bin

#-----------------------------------------------------------------------------------------------------------------------
# Rules (https://www.gnu.org/software/make/manual/html_node/Rule-Introduction.html#Rule-Introduction)
#-----------------------------------------------------------------------------------------------------------------------
.PHONY: help generate

help: ## Show this help
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

generate: ## Generate management accessor methods
	@echo "==> Generating management accessor methods..."
	@go generate ./...

#-----------------------------------------------------------------------------------------------------------------------
# Dependencies
#-----------------------------------------------------------------------------------------------------------------------
.PHONY: deps

deps: ## Download dependencies
	@echo "==> Downloading dependencies to vendor folder..."
	@go mod vendor -v

$(GO_BIN)/golangci-lint:
	${call print, "Installing golangci-lint"}
	@go install -v github.com/golangci/golangci-lint/cmd/golangci-lint@latest

$(GO_BIN)/govulncheck:
	${call print, "Installing govulncheck"}
	@go install -v golang.org/x/vuln/cmd/govulncheck@latest

#-----------------------------------------------------------------------------------------------------------------------
# Checks
#-----------------------------------------------------------------------------------------------------------------------
.PHONY: lint check-vuln check-getters

lint: $(GO_BIN)/golangci-lint ## Run linting on the go files
	@echo "==> Running linting on the library with golangci-lint..."
	@golangci-lint run -v --fix -c .golangci.yml ./...

check-vuln: $(GO_BIN)/govulncheck ## Check for vulnerabilities
	@echo "==> Checking for vulnerabilities..."
	@govulncheck -show verbose ./...

check-getters: ## Check that struct field getters were generated
	@echo "==> Checking that struct field getters were generated..."
	@$(MAKE) generate
	@if [ -n "$$(git status --porcelain)" ]; \
	then \
		echo "Go generate resulted in changed files:"; \
		echo "$$(git diff)"; \
		echo "Please run \`make generate\` to regenerate struct field getters."; \
		exit 1; \
	fi
	@echo "Struct field getters are generated correctly."

#-----------------------------------------------------------------------------------------------------------------------
# Testing
#-----------------------------------------------------------------------------------------------------------------------
.PHONY: test test-record test-e2e

test: ## Run tests with http recordings. To run a specific test pass the FILTER var. Usage `make test FILTER="TestResourceServer_Read"`
	@echo "==> Running tests with http recordings..."
	@AUTH0_HTTP_RECORDINGS=on \
		AUTH0_DOMAIN=go-auth0-dev.eu.auth0.com \
		go test \
		-run "$(FILTER)" \
		-cover \
		-covermode=atomic \
		-coverprofile=coverage.out \
		./...

test-record: ## Run tests and record http interactions. To run a specific test pass the FILTER var. Usage `make test-record FILTER="TestResourceServer_Read"`
	@echo "==> Running tests and recording http interactions..."
	@AUTH0_HTTP_RECORDINGS=on \
		go test \
		-run "$(FILTER)" \
		./...

test-e2e: ## Run tests without http recordings. To run a specific test pass the FILTER var. Usage `make test-e2e FILTER="TestResourceServer_Read"`
	@echo "==> Running tests against a real Auth0 tenant..."
	@go test \
		-run "$(FILTER)" \
		-cover \
		-covermode=atomic \
		-coverprofile=coverage.out \
		-timeout 20m \
		./...
