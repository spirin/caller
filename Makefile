export PROJECT_ROOT=$(shell pwd)

deps:
	@go mod tidy

test:
	@go test -race -tags=integration  ./...  -coverpkg=./... -coverprofile coverage.out

coverage:
	@go tool cover -html=coverage.out

lint:
	@echo 'linter run'
	docker run --rm -it \
		-v $(GOPATH)/pkg/mod:/go/pkg/mod \
		-v $(PROJECT_ROOT):/app \
		-w /app \
		golangci/golangci-lint:latest golangci-lint run

format:
	@go install github.com/daixiang0/gci@latest
	gci write -s standard -s default -s localmodule --skip-generated --skip-vendor .

ready: deps format test lint