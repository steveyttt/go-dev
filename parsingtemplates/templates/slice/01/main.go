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
	sages := []string{"George", "Muzz", "Grace", "Cid"}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", sages)
	if err != nil {
		log.Fatal(err)
	}
}
