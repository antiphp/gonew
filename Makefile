test:
	@go test -race -cover ./...
.PHONY: test

lint:
	@golangci-lint run --config=.golangci.yml ./...
.PHONY: lint

build:
	@goreleaser release --rm-dist --snapshot
.PHONY: build
