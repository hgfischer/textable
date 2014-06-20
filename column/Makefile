# http://wiki.osdev.org/Makefile
.PHONY: all clean default

TEMPLATE := generator/template.go.tpl

TYPES := \
	bool byte rune string error uintptr \
	complex128 complex64 float32 float64 \
	int int8 int16 int32 int64 \
	uint uint8 uint16 uint32 uint64 \
	fmt.Stringer

default: all

all: $(TEMPLATE)
	for type in $(TYPES); do go run generator/main.go $$type; done;
	goimports -w c_*.go
	go install

clean: 
	rm -f c_*.go
