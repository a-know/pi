VERSION = $(shell gobump show -r)
CURRENT_REVISION = $(shell git rev-parse --short HEAD)
BUILD_LDFLAGS = "-X github.com/a-know/pi.revision=$(CURRENT_REVISION)"

ifdef update
  u=-u
endif

GO ?= GO111MODULE=on go

.PHONY: deps
deps:
	env GO111MODULE=on go mod download

.PHONY: devel-deps
devel-deps: deps
	$(GO) get ${u} \
	  golang.org/x/lint/golint             \
	  github.com/rakyll/gotest             \
	  github.com/x-motemen/gobump \
	  github.com/Songmu/goxz/cmd/goxz      \
	  github.com/Songmu/ghch/cmd/ghch      \
	  github.com/tcnksm/ghr

.PHONY: test
test: devel-deps
	gotest -v -cover .

.PHONY: citest
citest: deps
	$(GO) test

.PHONY: lint
lint: devel-deps
	$(GO) vet
	golint -set_exit_status

.PHONY: build
build: deps
	$(GO) build -ldflags=$(BUILD_LDFLAGS) ./cmd/pi

.PHONY: crossbuild
crossbuild: devel-deps
	env GO111MODULE=on goxz -pv=v$(shell gobump show -r) -build-ldflags=$(BUILD_LDFLAGS) \
	  -os=linux,darwin,freebsd -arch=amd64 -d=./dist/v$(shell gobump show -r) \
	  ./cmd/pi
	env GO111MODULE=on goxz -pv=v$(shell gobump show -r) -build-ldflags=$(BUILD_LDFLAGS) \
	  -os=windows -arch=amd64 -d=./dist/v$(shell gobump show -r) \
	  -o=pi.exe ./cmd/pi

.PHONY: bump
bump: devel-deps
	_tools/releng

.PHONY: upload
upload:
	ghr v$(VERSION) dist/v$(VERSION)

.PHONY: release
release: bump crossbuild upload
