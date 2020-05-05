 # Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
ZIPCMD=zip
BINARY_NAME=localbitcoin_rate
BINARY_VERSION=0.2.0
BUILD_FOLDER=./build
BUILD_WINDOWS_FOLDER=$(BUILD_FOLDER)/windows
BUILD_LINUX_FOLDER=$(BUILD_FOLDER)/linux
    
all: test build
build: 
	[ -d $(BUILD_WIN_FOLDER) ] || mkdir -p $(BUILD_WIN_FOLDER)
	[ -d $(BUILD_LINUX_FOLDER) ] || mkdir -p $(BUILD_LINUX_FOLDER)
	GOOS=linux GOARCH=amd64 $(GOBUILD) -ldflags="-s -w -X localbitcoin_rate/cli.version=$(BINARY_VERSION)" -o $(BUILD_LINUX_FOLDER)/$(BINARY_NAME) -v
	GOOS=windows GOARCH=386 $(GOBUILD) -ldflags="-s -w -X localbitcoin_rate/cli.version=$(BINARY_VERSION)" -o $(BUILD_WINDOWS_FOLDER)/$(BINARY_NAME).exe -v
	zip -r $(BUILD_LINUX_FOLDER)/$(BINARY_NAME)-linux-amd64-$(BINARY_VERSION).zip $(BUILD_LINUX_FOLDER)/$(BINARY_NAME)
	zip -r $(BUILD_WINDOWS_FOLDER)/$(BINARY_NAME)-windows-i386-$(BINARY_VERSION).zip $(BUILD_WINDOWS_FOLDER)/$(BINARY_NAME).exe  
test: 
	$(GOTEST) -v ./...
clean: 
	$(GOCLEAN) 
	if [ -d $(BUILD_FOLDER) ]; then rm -Rf $(BUILD_FOLDER); fi
	rm -f $(BINARY_NAME)