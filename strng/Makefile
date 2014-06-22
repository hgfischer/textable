# http://wiki.osdev.org/Makefile
.PHONY: all clean default

default: all

clean:
	rm coverage.out

cover:
	go test -v -coverprofile=coverage.out
	go tool cover -html=coverage.out

test:
	go test -v

all: test
	go install
