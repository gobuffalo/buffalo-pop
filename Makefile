TAGS ?= "sqlite"
GO_BIN ?= go

deps:
	$(GO_BIN) get -v github.com/gobuffalo/packr/packr
	$(GO_BIN) get -tags ${TAGS} -v -t ./...

build: deps
	packr
	$(GO_BIN) build -v .

install: deps
	packr
	$(GO_BIN) install -v .

test:
	$(GO_BIN) test -tags ${TAGS} ./...

ci-test: deps
	$(GO_BIN) test -tags ${TAGS} -race ./...

update:
	$(GO_BIN) get -u
	$(GO_BIN) mod tidy
	make test
