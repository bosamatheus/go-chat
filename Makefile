.ONESHELL:

.PHONY: help
help:			## Show the help
	@echo "Usage: make <target>"
	@echo ""
	@echo "Targets:"
	@fgrep "##" Makefile | fgrep -v fgrep

.PHONY: install
install:		## Install the dependencies
	go mod download

.PHONY: build
build:			## Build and run the application
	docker-compose up -d --build

.PHONY: run
run:			## Run the application
	docker-compose up -d

.PHONY: lint
lint: 			## Lint the code with golangci-lint
	golangci-lint run

.PHONY: test
test:			## Run the tests
	go test ./...

.PHONY: test-cover
test-cover:		## Test the code with coverage
	go test ./... -coverprofile=coverage.out

.PHONY: clean
clean:			## Clean the build
	rm coverage.out
	docker-compose down
