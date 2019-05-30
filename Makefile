# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOINSTALL=$(GOCMD) install
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test

default:
	GO111MODULE=on $(GOTEST) ./...
install:
	GO111MODULE=on $(GOINSTALL)
test:
	GO111MODULE=on $(GOTEST) ./... -v