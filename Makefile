ifdef update
  u=-u
endif

GO ?= GO111MODULE=on go

deps:
	env GO111MODULE=on go mod download

devel-deps: deps
	$(GO) get ${u} \
	  golang.org/x/lint/golint             \
	  github.com/rakyll/gotest

test: devel-deps
	gotest -v -cover .

citest: deps
	$(GO) test

lint: devel-deps
	$(GO) vet
	golint -set_exit_status

build: deps
	$(GO) build ./cmd/pi

.PHONY: deps test build
