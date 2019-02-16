GO ?= GO111MODULE=on go

deps:
	env GO111MODULE=on go mod download

build: deps
	$(GO) build ./cmd/pi

.PHONY: deps build
