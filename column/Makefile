# http://wiki.osdev.org/Makefile
.PHONY: all clean default

default: all

all: clean
	go run generator/main.go
	goimports -w gen_*.go
	go test -v
	go install

clean: 
	rm -f gen_*.go
