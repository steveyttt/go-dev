package main

import (
	"text/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", `Release self focus, embrace other-focus`)
	if err != nil {
		log.Fatal(err)
	}
}
