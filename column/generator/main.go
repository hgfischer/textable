package main

import (
	"log"
	"os"
	"regexp"
	"strings"
	"text/template"
)

const PREFIX = "c_"

func main() {
	if len(os.Args) != 3 {
		log.Fatal("go generator/main.go [type name] [another type]")
	}

	t := Type{os.Args[1], os.Args[2]}

	ExecuteTemplate("generator/column_type.go.tpl", t, ".go")
	ExecuteTemplate("generator/test.go.tpl", t, "_test.go")
}

func ExecuteTemplate(templateFile string, t Type, outFnameSuffix string) {
	tpl, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create(t.Filename(outFnameSuffix))
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
	Type        string
	AnotherType string
}

func (t Type) CapsName() string {
	rxp := regexp.MustCompile(`^.+\.(.+)$`)
	return strings.Title(rxp.ReplaceAllString(t.Type, "$1")) + "Col"
}

func (t Type) Basename() string {
	rxp := regexp.MustCompile(`\W`)
	return PREFIX + rxp.ReplaceAllString(strings.ToLower(t.Type), "_")
}

func (t Type) Filename(suffix string) string {
	return t.Basename() + suffix
}