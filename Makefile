.DEFAULT_GOAL := build

GO_TEST_FLAGS ?= -race -count=1 -v -timeout=5m -json

build:
	go build -o bin/cli ./cmd/cli

test:
	go test $(GO_TEST_FLAGS) -coverprofile=cp.out -coverpkg=./... $$(go list ./... | grep -v -e /bin -e /cmd -e /examples -e /mocks)|\
	 	tparse --follow -sort=elapsed -trimpath=auto -all

cover:
	go tool cover -html=cp.out -o cover.html