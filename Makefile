GOFMT ?= gofmt "-s"
GOFILES := $(shell find . -name "*.go" -type f -not -path "./vendor/*")
GOPATH:=$(shell go env GOPATH)

.PHONY: fmt
fmt:
	$(GOFMT) -w $(GOFILES)

.PHONY: fmt-check
fmt-check:
	@diff=$$($(GOFMT) -d $(GOFILES)); \
	if [ -n "$$diff" ]; then \
		echo "Please run 'make fmt' and commit the result:"; \
		echo "$${diff}"; \
		exit 1; \
	fi; \
	echo "\033[34m[Code] format perfect!\033[0m";

.PHONY: proto
proto:
	protoc -I. -I${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate --go_out=plugins=grpc:. --micro_out=. --validate_out=lang=go:.  ./proto/**/*.proto