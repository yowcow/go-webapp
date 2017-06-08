.PHONY: all test

BINARY = go-webapp

all: Godeps $(BINARY)

Godeps:
	godep save

test:
	GIN_MODE=release go test -v

$(BINARY):
	go build -o $@

run:
	./$(BINARY)

clean:
	rm -rf $(BINARY)
