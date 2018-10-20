NAME := cmandroid-server
PKG := github.com/JoelPagliuca/$(NAME)

VERSION := $(shell cat VERSION.txt)
GITCOMMIT := $(shell git rev-parse --short HEAD)
GITUNTRACKEDCHANGES := $(shell git status --porcelain --untracked-files=no)
ifneq ($(GITUNTRACKEDCHANGES),)
	GITCOMMIT := $(GITCOMMIT)-dirty
endif

GO_LDFLAGS=-ldflags "-w -X main.GITCOMMIT=$(GITCOMMIT) -X main.VERSION=$(VERSION)"

.PHONY: help
help: ## Print this message and exit
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build: $(NAME) ## Builds a dynamic executable or package

$(NAME): $(wildcard *.go) VERSION.txt
	@echo "+ $@"
	go build $(GO_LDFLAGS) -o $(NAME) .

.PHONY: docs
docs: ## Builds documentation
	@echo "+ $@"
	./scripts/build-docs.sh

.PHONY: test
test: ## Runs the go tests
	@echo "+ $@"
	@go test -v $(shell go list ./... | grep -v vendor)

.PHONY: vet
vet: ## Verifies `go vet` passes
	@echo "+ $@"
	@go vet $(shell go list ./... | grep -v vendor)

.PHONY: clean
clean: ## Cleanup any build binaries or packages
	@echo "+ $@"
	$(RM) $(NAME)