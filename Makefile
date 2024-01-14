MAKEFLAGS += --no-print-directory

GOBIN ?= $(shell go env GOPATH)/bin

.DEFAULT_GOAL := check

.PHONE: deps
deps:
	go mod download -x

.PHONE: testdeps
testdeps: deps
	go install honnef.co/go/tools/cmd/staticcheck@2023.1.6

.PHONE: tidy
tidy:
	go mod verify
	go mod tidy

.PHONE: vet
vet: testdeps
	go vet ./...

.PHONE: staticcheck
staticcheck: testdeps
	$(GOBIN)/staticcheck ./...

.PHONE: lint
lint: vet staticcheck

.PHONE: test
test:
	go test -v -covermode=atomic -coverprofile=coverage.out ./...

.PHONE: check
check: test lint

.PHONE: clean
clean:
	go clean ./...
