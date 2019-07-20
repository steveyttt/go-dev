package main

import (
	"log"
	"os"
	"text/template"
)

//this var is accessible in the whole package
//do this to makesure the func init can be shared in the whole package
var tpl *template.Template

func init() {
	//template.must takes a pointer to a template.
	tpl = template.Must(template.ParseGlob("templates/*"))

}

func main() {
	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "vespa.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}

}
