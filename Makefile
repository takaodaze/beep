BINDIR:=bin

APP_NAME:=beep
ROOT_PACKAGE:=$(shell go list .)
ALL_PACKAGES:=$(shell go list ./...)

GO_FILES:=$(shell find . -type f -name '*.go' -print)

.PHONY: build
build:
	@go build -o $(BINDIR)/$(APP_NAME)  .

.PHONY: test
test:
	@go test ./...

.PHONY: run
run:
	@go run .