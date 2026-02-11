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

#-----------------------------------------------------------------------------------------------------------------------
# Dependencies
#-----------------------------------------------------------------------------------------------------------------------
.PHONY: deps

deps: ## Download dependencies
	@echo "==> Downloading dependencies to vendor folder..."
	@go mod vendor -v

$(GO_BIN)/golangci-lint:
	${call print, "Installing golangci-lint"}
	@go install -v github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest

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

#-----------------------------------------------------------------------------------------------------------------------
# Testing
#-----------------------------------------------------------------------------------------------------------------------
.PHONY: test test-record test-e2e

WIREMOCK_COMPOSE_FILE := wiremock/docker-compose.test.yml

test: ## Run tests. To run a specific test pass the FILTER var. Usage `make test FILTER="TestResourceServer_Read"`
	@WIREMOCK_STARTED_BY_MAKE=false; \
	if [ -n "$$WIREMOCK_PORT" ]; then \
		WIREMOCK_URL="http://localhost:$$WIREMOCK_PORT"; \
		if curl -s "$$WIREMOCK_URL/__admin/mappings" > /dev/null 2>&1; then \
			echo "==> WireMock is already running at $$WIREMOCK_URL"; \
		else \
			echo "==> WIREMOCK_PORT=$$WIREMOCK_PORT set but WireMock is not reachable"; \
			exit 1; \
		fi; \
	else \
		echo "==> Starting WireMock container..."; \
		docker compose -f $(WIREMOCK_COMPOSE_FILE) up -d; \
		WIREMOCK_STARTED_BY_MAKE=true; \
		WIREMOCK_PORT=$$(docker compose -f $(WIREMOCK_COMPOSE_FILE) port wiremock 8080 | cut -d: -f2); \
		WIREMOCK_URL="http://localhost:$$WIREMOCK_PORT"; \
		echo "==> WireMock assigned port: $$WIREMOCK_PORT"; \
		until curl -sf "$$WIREMOCK_URL/__admin/mappings" > /dev/null 2>&1; do \
			sleep 0.2; \
		done; \
		echo "==> WireMock is ready at $$WIREMOCK_URL"; \
	fi; \
	echo "==> Running tests with http recordings..."; \
	AUTH0_HTTP_RECORDINGS=on \
		AUTH0_DOMAIN=go-auth0-dev.eu.auth0.com \
		WIREMOCK_PORT=$$WIREMOCK_PORT \
		WIREMOCK_URL=$$WIREMOCK_URL \
		go test \
		-run "$(FILTER)" \
		-cover \
		-coverpkg=./... \
		-covermode=atomic \
		-coverprofile=coverage.out \
		./... ; \
	EXIT_CODE=$$?; \
	if [ "$$WIREMOCK_STARTED_BY_MAKE" = "true" ]; then \
		echo "==> Stopping WireMock container..."; \
		docker compose -f $(WIREMOCK_COMPOSE_FILE) down -v; \
	fi; \
	exit $$EXIT_CODE

test-record: ## Run authentication tests and record http interactions. To run a specific test pass the FILTER var. Usage `make test-record FILTER="TestResourceServer_Read"`
	@echo "==> Running authentication tests and recording http interactions..."
	@AUTH0_HTTP_RECORDINGS=on \
		go test \
		-run "$(FILTER)" \
		./authentication/...

test-e2e: ## Run authentication tests without http recordings. To run a specific test pass the FILTER var. Usage `make test-e2e FILTER="TestResourceServer_Read"`
	@echo "==> Running authentication tests against a real Auth0 tenant..."
	@go test \
		-run "$(FILTER)" \
		-cover \
		-coverpkg=./... \
		-covermode=atomic \
		-coverprofile=coverage.out \
		-timeout 20m \
		./authentication/...
