all: 2020

.PHONY: 2020
2020:
	@for file in $(shell ls 2020/*/main.go); do \
		echo $$file; \
		go run $$file; \
		echo; \
	done
