# VARIABLES
PACKAGE=""
BINARY_NAME="server"

default: run

clean: ## Trash binary files
	@echo "--> cleaning..."
	@go clean || (echo "Unable to clean project" && exit 1)
	@rm -rf $(GOPATH)/bin/$(BINARY_NAME) 2> /dev/null
	@echo "Clean OK"

test: ## Run all tests
	@echo "--> testing..."
	@go test -v ./tests

install: clean ## Compile sources and build binary
	@echo "--> installing..."
	@go install $(PACKAGE) || (echo "Compilation error" && exit 1)
	@echo "Install OK"

run: install ## Run your application
	@echo "--> running application..."
	@/$(GOBIN)/$(BINARY_NAME)

deploy: clean
	@echo "--> preparing deploy..."
	@GOARCH=amd64 GOOS=linux go build $(PACKAGE) -o $(GOBIN)/$(BINARY_NAME) || (echo "Compilation error" && exit 1)
	@echo "deploy OK, now push the /bin/server and use eb deploy"
	##falta hacer el webpack?

