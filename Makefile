# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOINSTALL=$(GOCMD) install
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
ENV=GO111MODULE=on CGO_ENABLED=1

default:
	$(ENV) $(GOTEST) ./...

.PHONY: install
install:
	$(ENV) $(GOINSTALL)

.PHONY: test
test:
	$(ENV) $(GOTEST) ./... -v

.PHONY: travis-test
travis-test:
	$(ENV) $(GOTEST) ./... -v -coverprofile=coverage.txt -covermode=atomic