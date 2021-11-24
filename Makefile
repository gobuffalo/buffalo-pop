test:
	go test -tags sqlite -failfast -short -cover ./...
	go mod tidy -v

cov:
	go test -short -coverprofile cover.out ./...
	go tool cover -html cover.out
	go mod tidy -v

install:
	go install -v .

