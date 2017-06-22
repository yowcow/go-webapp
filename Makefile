.PHONY: all Godeps build test

BINARY = go-webapp

all: Godeps build

Godeps:
	godep save

build: $(BINARY)

$(BINARY):
	go build -o $@

test:
	GIN_MODE=release go test ./action ./ -v

run:
	go run app.go

clean:
	rm -rf $(BINARY)
