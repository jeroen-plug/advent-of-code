DAYS := $(shell seq 1 12)

.PHONY: all $(DAYS) lint test input

all: $(DAYS)

$(DAYS):
	@cargo solve $@

lint:
	@cargo clippy

test:
	@cargo test

input:
	@../tools/downloadInput.sh
