.PHONY : get clean plugin

# LATEST_TAG := $(shell git describe $(shell git rev-list --tags --max-count=1))

default: get clean plugin 

get:
	go get ./...

test:
	go test -v ./...

plugin:
	go build -buildmode=plugin .

clean:
	rm -f cloudless-init.so
