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
	//pass in the value 42 into the template called tpl.gohtml
	//look inside the .gohtml file and you will see {{.}} - This is what we replace with 42
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 42)
	if err != nil {
		log.Fatal(err)
	}
}
