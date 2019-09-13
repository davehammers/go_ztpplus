SUBDIRS := $(shell ls */Makefile 2>/dev/null)
BUILD_DIRS := $(subst Makefile,build,$(SUBDIRS))
TEST_DIRS := $(subst Makefile,test,$(SUBDIRS))
LINT_DIRS := $(subst Makefile,lint,$(SUBDIRS))
STATICCHECK = $(HOME)/go/bin/staticcheck
GO_FILES = $(wildcard *.go)

.PHONY: all $(SUBDIRS) dev

all: $(SUBDIRS) dev

build: $(BUILD_DIRS)
ifneq ($(GO_FILES),)
	go fmt
	go build
endif
$(BUILD_DIRS):
	$(MAKE) -C $(@D) build

test: $(TEST_DIRS)
ifneq ($(GO_FILES),)
	go test -cover
endif
$(TEST_DIRS):
	$(MAKE) -C $(@D) test

lint: $(LINT_DIRS)
ifneq ($(GO_FILES),)
	$(HOME)/go/bin/staticcheck
endif
$(LINT_DIRS):
	$(MAKE) -C $(@D) lint

$(SUBDIRS): dev
	$(MAKE) -C $(@D)

$(STATICCHECK):
	go get honnef.co/go/tools/cmd/staticcheck

dev: $(STATICCHECK)
ifneq ($(GO_FILES),)
	go fmt
	go build
	go test -cover
	$(STATICCHECK)
endif
