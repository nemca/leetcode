#!/usr/bin/env make -f

.PHONY: all test

all: test

test:
	@pkg=$(filter-out test,$(MAKECMDGOALS)); \
	if [ -n "$$pkg" ]; then \
		go test -race -v ./$$pkg; \
	else \
		go test -race ./...; \
	fi
