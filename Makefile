DAYS = $(wildcard day*)

FORMATTED_NUMBERS := $(sort $(foreach n, $(foreach d,$(DAYS),$(d:day%=%)), $(if $(filter 1 2 3 4 5 6 7 8 9,$n),0$n,$n)))
DAY_TARGETS := $(foreach n, $(FORMATTED_NUMBERS), $(if $(filter 0%, $n),$(subst 0,,$n),$n))

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
