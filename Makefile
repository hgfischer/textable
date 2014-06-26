# http://wiki.osdev.org/Makefile
.PHONY: all clean default

default: all

clean:
	rm coverage.out

deps:
	go get github.com/stretchr/testify

cover: deps
	go test -v -coverprofile=coverage.out
	go tool cover -html=coverage.out

test: deps
	go test -v

all: test
	go install
