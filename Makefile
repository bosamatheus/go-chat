.ONESHELL:

.PHONY: help
help:		## Show the help.
	@echo "Usage: make <target>"
	@echo ""
	@echo "Targets:"
	@fgrep "##" Makefile | fgrep -v fgrep

.PHONY: install
install:	## Install the dependencies
	go mod download

.PHONY: run
run:		## Run the application
	go run cmd/gochat/main.go

.PHONY: lint
lint: 		## Lint the code with golangci-lint
	golangci-lint run
