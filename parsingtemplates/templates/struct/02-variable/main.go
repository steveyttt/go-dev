package main

import (
	"text/template"
	"log"
	"os"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	buddha := sage{
		Name:  "Buddha",
		Motto: "You are my man",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", buddha)
	if err != nil {
		log.Fatal(err)
	}
}
