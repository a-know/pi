GO ?= GO111MODULE=on go

deps:
	env GO111MODULE=on go mod download

test: deps
	gotest -v -cover .

build: deps
	$(GO) build ./cmd/pi

.PHONY: deps test build
