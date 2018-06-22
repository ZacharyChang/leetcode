# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOINSTALL=$(GOCMD) install
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test

default:
	$(GOTEST) ./...
install:
	$(GOINSTALL)
test:
	$(GOTEST) ./... -v