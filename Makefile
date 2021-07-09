MODULE = $(shell go list -m)
#VERSION ?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || echo "1.0.0")
VERSION = "1.0.0"
PACKAGES := $(shell go list ./... | grep -v /vendor/)
LDFLAGS := -ldflags "-X main.Version=${VERSION}"
CONFIG_FILE ?= ./config/local.yml
# generate help info from comments: thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help
help: ## help information about make commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: version
version: ## display the version of the API server
	@echo $(VERSION)

.PHONY: run
run: ## run the API server
	go run ${LDFLAGS} cmd/server/main.go

.PHONY: build
build:  ## build the API server binary
	CGO_ENABLED=0 go build ${LDFLAGS} -a -o server $(MODULE)/cmd/server

.PHONY: build-docker-local
build-docker-local: ## build the API server as a docker image
	docker build --build-arg APP_ENV=local -f cmd/server/Dockerfile -t go-rest-dev .

.PHONY: build-docker-staging
build-docker-staging: ## build the API server as a docker image
	docker build --build-arg APP_ENV=staging -f cmd/server/Dockerfile -t go-rest-dev .

.PHONY: run-docker
run-docker: ## run the API server as a docker image
	docker run --name microservice -d -p 8080:8080 go-rest-dev 

.PHONY: clean
clean: ## remove temporary files
	rm -rf server