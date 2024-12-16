DAYS = $(wildcard day*)

FORMATTED_NUMBERS := $(sort $(foreach n, $(foreach d,$(DAYS),$(d:day%=%)), $(if $(filter 1 2 3 4 5 6 7 8 9,$n),0$n,$n)))
DAY_TARGETS := $(foreach n, $(FORMATTED_NUMBERS), $(if $(filter 0%, $n),$(subst 0,,$n),$n))
COMPILED_DAY_TARGETS := $(addprefix x,$(DAY_TARGETS))

EXE := advent-of-code-2024.exe
GOFILES := $(shell powershell Get-ChildItem . -Recurse -Name *.go)

.PHONY: all $(DAY_TARGETS) build test bench fmt clean

all: $(COMPILED_DAY_TARGETS)

$(DAY_TARGETS):
	@go run . $@

$(COMPILED_DAY_TARGETS): $(EXE)
	@./$(EXE) $(@:x%=%)

day%:
	@powershell ./create_day.ps1 $(@:day%=%)

build: $(EXE)

$(EXE): $(GOFILES)
	@go build .

lint:
	@go vet ./...

test:
	@go test ./day...

bench:
	@go test -benchmem -run=^$$ -bench .* ./day...

fmt:
	@go fmt ./...

clean:
	@go clean -r -testcache
