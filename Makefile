# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GORUN=$(GOCMD) run
BINARY_NAME=libinput-gestures
BUILD_FOLDER=build

all: test build
clean: 
	$(GOCLEAN)
	rm -fr $(BUILD_FOLDER)
test: 
	$(GOTEST) -v ./...
build: clean
	$(GOBUILD) -o "$(BUILD_FOLDER)/$(BINARY_NAME)" -v cmd/*.go
setup:
	$(GOGET) github.com/onsi/ginkgo/ginkgo
	$(GOGET) github.com/onsi/gomega/...
