.ONESHELL:

.PHONY: help
help:		## Show the help
	@echo "Usage: make <target>"
	@echo ""
	@echo "Targets:"
	@fgrep "##" Makefile | fgrep -v fgrep

.PHONY: install
install:	## Install the dependencies
	go mod download

.PHONY: build
build:		## Build and run the application
	docker-compose up -d --build

.PHONY: run
run:		## Run the application
	docker-compose up -d

.PHONY: lint
lint: 		## Lint the code with golangci-lint
	golangci-lint run

.PHONY: clean
clean:		## Clean the build
	docker-compose down
