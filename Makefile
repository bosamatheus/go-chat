GO_BUILD_ENV := CGO_ENABLED=0 GOOS=linux GOARCH=amd64

help:			## Show the help
	@echo "Usage: make <target>"
	@echo ""
	@echo "Targets:"
	@fgrep "##" Makefile | fgrep -v fgrep

install:		## Install the dependencies
	go mod download

build:			## Build the binary
	$(GO_BUILD_ENV) go build -o bin/gochat cmd/gochat/main.go

docker-run:		## Run the application in a docker container
	docker-compose up -d

docker-build:		## Build and run the application in a docker container
	docker-compose up -d --build

lint:			## Lint the code with golangci-lint
	golangci-lint run

test:			## Run the tests
	go test ./...

test-cover:		## Test the code with coverage
	go test ./... -coverprofile=coverage.out

clean:			## Clean the build
	rm -rf bin/
	rm coverage.out
	docker-compose down
