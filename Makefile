.PHONY: all test

all: Godeps

Godeps:
	godep save

test:
	GIN_MODE=release go test -v
