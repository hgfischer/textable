package main

import (
	"log"
	"os"
	"regexp"
	"strings"
	"text/template"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("go generator/main.go [type name]")
	}

	t := Type{os.Args[1]}

	tpl, err := template.ParseFiles("generator/template.go.tpl")
	if err != nil {
		log.Fatal(err)
	}

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

type Type struct {
	Type string
}

func (t Type) CapsName() string {
	rxp := regexp.MustCompile(`^.+\.(.+)$`)
	return strings.Title(rxp.ReplaceAllString(t.Type, "$1"))
}

func (t Type) Filename() string {
	rxp := regexp.MustCompile(`\W`)
	return "c_" + rxp.ReplaceAllString(strings.ToLower(t.Type), "_") + ".go"
}