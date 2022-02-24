.ONESHELL:

.PHONY: help
help:		## Show the help.
	@echo "Usage: make <target>"
	@echo ""
	@echo "Targets:"
	@fgrep "##" Makefile | fgrep -v fgrep

.PHONY: install
install:	## Run the project
	go mod download

.PHONY: run
run:		## Run the application
	go run cmd/gochat/main.go
