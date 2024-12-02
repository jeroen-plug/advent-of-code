DAYS = $(wildcard day*)
DAY_TARGETS = $(foreach d,$(DAYS),$(d:day%=%))

.PHONY: all $(DAY_TARGETS) build test clean

all: $(DAY_TARGETS)

$(DAY_TARGETS):
	@go run . $@

day%:
	@powershell ./create_day.ps1 $(@:day%=%)

build:
	@go build .

test:
	@go test ./day...

clean:
	@go clean -r -testcache
