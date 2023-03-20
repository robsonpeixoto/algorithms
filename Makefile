MAKEFLAGS += --always-make

test:
	go test -v ./...

fmt:
	go fmt ./...

lint:
	golangci-lint run
