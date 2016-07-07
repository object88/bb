.DEFAULT_GOAL=build
UNAME=$(shell uname)
PREFIX=github.com/object88/bbgraph
GOVERSION=$(shell go version)
BUILD_SHA=$(shell git rev-parse HEAD)

ifeq "$(UNAME)" "Darwin"
    BUILD_FLAGS=-ldflags="-s -X main.Build=$(BUILD_SHA)"
else
    BUILD_FLAGS=-ldflags="-X main.Build=$(BUILD_SHA)"
endif

# Workaround for GO15VENDOREXPERIMENT bug (https://github.com/golang/go/issues/11659)
ALL_PACKAGES=$(shell go list ./... | grep -v /vendor/ | grep -v /scripts)

# We must compile with -ldflags="-s" to omit
# DWARF info on OSX when compiling with the
# 1.5 toolchain. Otherwise the resulting binary
# will be malformed once we codesign it and
# unable to execute.
# See https://github.com/golang/go/issues/11887#issuecomment-126117692.
ifeq "$(UNAME)" "Darwin"
	TEST_FLAGS=-exec=$(shell pwd)/scripts/testsign
	DARWIN="true"
endif

# If we're on OSX make sure the proper CERT env var is set.
check-cert:
ifdef DARWIN
ifeq "$(CERT)" ""
	$(error You must provide a CERT environment variable in order to codesign the binary.)
endif
endif

build: check-cert
	go build $(BUILD_FLAGS) github.com/object88/bbgraph/bbgraph
ifdef DARWIN
ifneq "$(GOBIN)" ""
	codesign -s "$(CERT)"  $(GOBIN)/bbgraph
else
	codesign -s "$(CERT)"  $(GOPATH)/bin/bbgraph
endif
endif

install: check-cert
	go install $(BUILD_FLAGS) github.com/object88/bbgraph/bbgraph
ifdef DARWIN
ifneq "$(GOBIN)" ""
	codesign -s "$(CERT)"  $(GOBIN)/bbgraph
else
	codesign -s "$(CERT)"  $(GOPATH)/bin/bbgraph
endif
endif
