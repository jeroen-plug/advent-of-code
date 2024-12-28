DAYS := $(shell seq 1 25)

.PHONY: all $(DAYS) test lint input

all: $(DAYS)

$(DAYS):
	@./main.py $@

test:
	@pytest

lint:
	@flake8 . --count --select=E9,F63,F7,F82 --show-source --statistics
	@flake8 . --count --max-complexity=10 --max-line-length=127 --statistics

input:
	@../tools/downloadInput.sh
