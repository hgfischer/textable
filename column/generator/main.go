package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

const PREFIX = "g_"

var TYPES = []string{
	"bool", "uintptr", "complex128", "complex64", "float32", "float64",
	"int", "int8", "int16", "int32", "int64",
	"uint", "byte", "uint16", "uint32", "uint64",
	"string", // the last one is also the default type of the constructor
	// byte is an alias to uint8
	// rune is an alias to int32
}

func main() {
	allTypes := &AllTypes{}
	lastType := TYPES[len(TYPES)-1]
	for _, st := range TYPES {
		t := Type{st, lastType}
		allTypes.Types = append(allTypes.Types, &t)
		lastType = st
		ExecuteTemplate("generator/column_type.go.tpl", t, ".go")
		ExecuteTemplate("generator/test.go.tpl", t, "_test.go")
		allTypes.DefaultType = t
	}
	ExecuteTemplate("generator/constructor.go.tpl", allTypes, ".go")
}

func ExecuteTemplate(templateFile string, t HasFilename, suffix string) {
	tpl, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create(t.Filename(suffix))
	if err != nil {
		log.Fatal(file, err)
	}
	defer file.Close()

	err = tpl.Execute(file, t)
	if err != nil {
		log.Fatal(err)
	}
}

type HasFilename interface {
	Filename(suffix string) string
}

type Type struct {
	Type        string
	AnotherType string
}

func (t Type) CapsName() string {
	return strings.Title(t.Type) + "Col"
}

func (t Type) Filename(suffix string) string {
	return "gen_" + t.Type + "_col" + suffix
}

type AllTypes struct {
	Types       []*Type
	DefaultType Type
}

func (t AllTypes) Filename(suffix string) string {
	return "gen_constructor" + suffix
}