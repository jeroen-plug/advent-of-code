DAYS := $(shell seq 1 25)

.PHONY: all $(DAYS) test input

all: $(DAYS)

$(DAYS):
	@./main.py $@

test:
	@pytest

input:
	@../tools/downloadInput.sh
