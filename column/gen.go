package main

import (
	"log"
	"os"
	"regexp"
	"strings"
	"text/template"
)

type Type struct {
	Type string
}

func (t Type) CapsName() string {
	rxp := regexp.MustCompile(`^.+\.(.+)$`)
	return strings.Title(rxp.ReplaceAllString(t.Type, "$1"))
}

func (t Type) Filename() string {
	rxp := regexp.MustCompile(`\W`)
	return "_" + rxp.ReplaceAllString(strings.ToLower(t.Type), "_") + ".go"
}

func main() {
	var types = []Type{
		{"bool"},
		{"byte"},
		{"rune"},
		{"string"},
		{"error"},
		{"complex128"},
		{"complex64"},
		{"float32"},
		{"float64"},
		{"int"},
		{"int8"},
		{"int16"},
		{"int32"},
		{"int64"},
		{"uint"},
		{"uint8"},
		{"uint16"},
		{"uint32"},
		{"uint64"},
		{"uintptr"},
		{"fmt.Stringer"},
	}

	tpl, err := template.ParseFiles("template.go.tpl")
	if err != nil {
		log.Fatal(err)
	}

	for _, t := range types {
		file, err := os.Create(t.Filename())
		if err != nil {
			log.Fatal(file, err)
		}
		defer file.Close()

		err = tpl.Execute(file, t)
		if err != nil {
			log.Fatal(err)
		}
	}
}