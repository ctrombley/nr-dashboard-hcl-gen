GO_PKGS      ?= $(shell go list ./...)

all: clean snapshot

test:
	go test $(GO_PKGS)

clean:
	@rm -rf dist

snapshot:
	goreleaser --snapshot --skip-publish --rm-dist

publish:
	goreleaser

.PHONY: all build