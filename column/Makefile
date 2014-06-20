# http://wiki.osdev.org/Makefile
.PHONY: all clean default

TYPES := \
	bool byte rune string uintptr \
	complex128 complex64 float32 float64 \
	int int8 int16 int32 int64 \
	uint uint8 uint16 uint32 uint64

LAST_TYPE := uint64

default: all

all: clean
	last_type=$(LAST_TYPE); \
	for type in $(TYPES); do \
		echo go run generator/main.go $$type $$last_type; \
		go run generator/main.go $$type $$last_type; \
		last_type=$$type; \
	done;
	goimports -w c_*.go
	go test -v
	go install

clean: 
	rm -f c_*.go
