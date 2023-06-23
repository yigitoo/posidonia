all: run
all-wsetup: run-wsetup
run:
	python3 scripts/make.py
run-wsetup:
	python3 scripts/make.py --wsetup
.PHONY: all all-wsetup
