# vi: ft=make

GOPATH:=$(shell go env GOPATH)

.PHONY: build
build:
	go build -o build/user cmd/main.go
    
.PHONY: test
test:
	@go get github.com/rakyll/gotest
	gotest -p 1 -race -cover -v ./...

.PHONY: lint
lint:
	@go get github.com/golangci/golangci-lint/cmd/golangci-lint
	golangci-lint run ./...

.PHONY: generate_sql
generate_sql:
	@go get -u -t github.com/volatiletech/sqlboiler
	@go get github.com/volatiletech/sqlboiler/drivers/sqlboiler-psql
	sqlboiler --wipe --no-tests -o ./server/models psql

.PHONY: generate_mocks
generate_mocks:
	@go get github.com/golang/mock/gomock
	@go install github.com/golang/mock/mockgen
	mockgen -package crypt -destination ./pkg/crypt/mock_crypt.go github.com/taeho-io/user/pkg/crypt Crypt

.PHONY: clean_mocks
clean_mocks:
	find . -name "mock_*.go" -type f -delete
	rm -rf mocks
