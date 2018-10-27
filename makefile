NAME := cmandroid-server
PKG := github.com/JoelPagliuca/$(NAME)

VERSION := $(shell cat VERSION.txt)
GITCOMMIT := $(shell git rev-parse --short HEAD)
GITUNTRACKEDCHANGES := $(shell git status --porcelain --untracked-files=no)
ifneq ($(GITUNTRACKEDCHANGES),)
	GITCOMMIT := $(GITCOMMIT)-dirty
endif

GO_LDFLAGS=-ldflags "-w -X main.GITCOMMIT=$(GITCOMMIT) -X main.VERSION=$(VERSION)"
PLATFORMS=$(shell cat platforms.txt)

BUILDDIR=build

.PHONY: help
help: ## Print this message and exit
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build: $(NAME) ## Builds an executable for host platform

$(NAME): static $(wildcard *.go) VERSION.txt
	@echo "+ $@"
	go build $(GO_LDFLAGS) -o $(NAME) .

# Cross compilation function
define buildcross
mkdir -p $(BUILDDIR)/$(1)/$(2);
GOOS=$(1) GOARCH=$(2) CGO_ENABLED=$(CGO_ENABLED) go build $(GO_LDFLAGS)\
	 -o $(BUILDDIR)/$(1)/$(2)/$(NAME)-$(1)-$(2);
md5sum $(BUILDDIR)/$(1)/$(2)/$(NAME)-$(1)-$(2) > $(BUILDDIR)/$(1)/$(2)/$(NAME)-$(1)-$(2).md5;
sha256sum $(BUILDDIR)/$(1)/$(2)/$(NAME)-$(1)-$(2) > $(BUILDDIR)/$(1)/$(2)/$(NAME)-$(1)-$(2).sha256;
endef

.PHONY: cross
cross: static *.go VERSION.txt ## Builds the cross-compiled binaries
	@echo "+ $@"
	$(foreach GOOSARCH,$(PLATFORMS), $(call buildcross,$(subst /,,$(dir $(GOOSARCH))),$(notdir $(GOOSARCH))))

.PHONY: docs
docs: ## Builds documentation
	@echo "+ $@"
	./scripts/build-docs.sh

.PHONY: static
static: ## Compile the static content for the server
	@echo "+ $@"
	go get -u github.com/rakyll/statik
	statik -src=./go/swaggerui/
	mv -f statik go/statik

.PHONY: test
test: ## Runs the go tests
	@echo "+ $@"
	@go test -v $(shell go list ./... | grep -v vendor) | grep -v "no test files"

.PHONY: vet
vet: ## Verifies `go vet` passes
	@echo "+ $@"
	@go vet $(shell go list ./... | grep -v vendor)

.PHONY: clean
clean: ## Cleanup any build binaries or packages
	@echo "+ $@"
	rm -f $(NAME)
	rm -rf $(BUILDDIR)
	rm -rf ./go/statik/