TAGS ?= "sqlite"
GO_BIN ?= go

install: generate
	$(GO_BIN) install -tags ${TAGS} -v

tidy:
ifeq ($(GO111MODULE),on)
	$(GO_BIN) mod tidy
else
	echo skipping go mod tidy
endif

deps:
	$(GO_BIN) get github.com/gobuffalo/release
	$(GO_BIN) get github.com/gobuffalo/packr/v2/packr2
	$(GO_BIN) get -tags ${TAGS} -t ./...
	make tidy

generate:
	$(GO_BIN) generate

build: generate tidy
	$(GO_BIN) build -v .

test: generate tidy
	$(GO_BIN) test -tags ${TAGS} ./...

ci-test:
	$(GO_BIN) test -tags ${TAGS} -race ./...

lint:
	gometalinter --vendor ./... --deadline=1m --skip=internal

update:
	$(GO_BIN) get -u -tags ${TAGS}
	make generate
	make tidy
	make test
	make install

release-test:
	make test

release:
	release -y -f ./pop/version.go
