//range over tpl.gohtml template with a list of sages
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

	sages := map[string]string{
		"Father":   "Steve",
		"Mother":   "Marianne",
		"Daughter": "Grace",
		"Son":      "George",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", sages)
	if err != nil {
		log.Fatal(err)
	}
}
