# If you update this file, please follow
# https://suva.sh/posts/well-documented-makefiles

# Ensure Make is run with bash shell as some syntax below is bash-specific
SHELL := /usr/bin/env bash

.DEFAULT_GOAL := help

# Active module mode, as we use go modules to manage dependencies
export GO111MODULE := on

# Directories
BIN_DIR       := bin
TOOLS_DIR     := hack/tools
TOOLS_BIN_DIR := $(TOOLS_DIR)/bin
export PATH := $(abspath $(BIN_DIR)):$(abspath $(TOOLS_BIN_DIR)):$(PATH)

# Tooling binaries
CONTROLLER_GEN     := $(TOOLS_BIN_DIR)/controller-gen
GOLANGCI_LINT      := $(TOOLS_BIN_DIR)/golangci-lint

# Allow overriding manifest generation destination directory
MANIFEST_ROOT ?= config
CRD_ROOT      ?= $(MANIFEST_ROOT)/crd/bases

.PHONY: all
all: lint tools generate ## Runs tests and generates all components

## --------------------------------------
##@ Help
## --------------------------------------

help: ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-18s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

## --------------------------------------
##@ Tooling
## --------------------------------------

TOOLING_BINARIES := $(CONTROLLER_GEN) $(GOLANGCI_LINT)
tools: $(TOOLING_BINARIES) ## Build tooling binaries
.PHONY: $(TOOLING_BINARIES)
$(TOOLING_BINARIES):
	make -C $(TOOLS_DIR) $(@F)

## --------------------------------------
##@ Generate
## --------------------------------------

.PHONY: modules
modules: ## Validates the modules
	go mod tidy

.PHONY: modules-download
modules-download: ## Downloads and caches the modules
	go mod download

.PHONY: generate
generate: ## Run all code generation targets
	$(MAKE) generate-go
	$(MAKE) generate-manifests

.PHONY: generate-go
generate-go: $(CONTROLLER_GEN) ## Runs Go related generate targets
	$(CONTROLLER_GEN) \
		paths=./api/... \
		object:headerFile="$(abspath hack/boilerplate/boilerplate.go.txt)"
ifneq (0,$(GENERATE_CODE))
	go generate ./...
endif

.PHONY: generate-manifests
generate-manifests: $(CONTROLLER_GEN) ## Generate manifests e.g. CRD, RBAC etc.
	$(CONTROLLER_GEN) \
		paths=./api/... \
		crd:crdVersions=v1 \
		output:crd:dir=$(CRD_ROOT) \
		output:none

## --------------------------------------
##@ Linting
## --------------------------------------

.PHONY: lint
lint: ## Run all the lint targets
	$(MAKE) lint-go-full
	$(MAKE) lint-markdown
#	$(MAKE) lint-shell

GOLANGCI_LINT_FLAGS ?= --fast=true
.PHONY: lint-go
lint-go: $(GOLANGCI_LINT) ## Lint codebase
	$(GOLANGCI_LINT) run -v $(GOLANGCI_LINT_FLAGS)

.PHONY: lint-go-full
lint-go-full: GOLANGCI_LINT_FLAGS = --fast=false
lint-go-full: lint-go ## Run slower linters to detect possible issues

.PHONY: lint-markdown
lint-markdown: ## Lint the project's markdown
	docker run --rm -v $$(pwd):/mnt mivok/markdownlint -c /mnt/.mdlrc /mnt

.PHONY: lint-shell
lint-shell: ## Lint the project's shell scripts
	docker run --rm -v "$$(pwd)":/mnt:ro koalaman/shellcheck:stable $$(ls hack/*.sh)

## --------------------------------------
##@ Cleanup
## --------------------------------------

.PHONY: clean
clean: # Clean all generated or compiled files
	$(MAKE) clean-bin 
	$(MAKE) clean-crd
	$(MAKE) modules

.PHONY: clean-bin
clean-bin: ## Remove all generated tooling binaries
	rm -rf hack/tools/bin
	rm -rf hack/samples/bin

.PHONY: clean-crd
clean-crd: ## Remove all generated crds
	rm -rf config/crd
