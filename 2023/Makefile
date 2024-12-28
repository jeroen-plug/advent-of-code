DAYS := $(shell seq 1 25)

.PHONY: all $(DAYS) lint test bench fmt clean input

all:
	@go run .

$(DAYS):
	@go run . $@

lint:
	@go vet ./...

test:
	@go test ./day...

bench:
	@go test -benchmem -run=^$$ -bench '.*' ./day...

fmt:
	@go fmt ./...

clean:
	@go clean -r -testcache

input:
	@../tools/downloadInput.sh
