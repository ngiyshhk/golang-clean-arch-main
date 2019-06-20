NAME     := golang-clean-arch-main
VERSION  := v0.0.1
REVISION := $(shell git rev-parse --short HEAD)

SRCS    := $(shell find . -type f -name '*.go')
LDFLAGS := -ldflags="-s -w -X \"main.Version=$(VERSION)\" -X \"main.Revision=$(REVISION)\" -extldflags \"-static\""

.PHONY: dep
dep:
	go get -u github.com/golang/dep/cmd/dep
	go get -u golang.org/x/lint/golint

.PHONY: deps
deps: dep
	dep ensure

.PHONY: lint
lint:
	golint -set_exit_status $$(go list ./...)
	go vet ./...

.PHONY: all
all: deps lint $(SRCS)
	GOOS=linux GOARCH=amd64 go build -a -tags netgo -installsuffix netgo $(LDFLAGS) -o bin/clean_arch_api main.go

.PHONY: clean
clean:
	rm -rf bin/*
	rm -rf vendor/*