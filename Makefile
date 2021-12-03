TAGS ?= "sqlite"

test:
	go test -tags ${TAGS} -failfast -short -cover ./...
	go mod tidy -v

cov:
	go test -tags ${TAGS} -short -coverprofile cover.out ./...
	go tool cover -html cover.out
	go mod tidy -v

install:
	go install -v -tags ${TAGS} .

