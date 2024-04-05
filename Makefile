BINDIR:=bin

ROOT_PACKAGE:=$(shell go list .)
ALL_PACKAGES:=$(shell go list ./...)

GO_FILES:=$(shell find . -type f -name '*.go' -print)

.PHONY: build
build:
	@go build -o $(BINDIR)/beep  .

.PHONY: test
test:
	@go test ./...