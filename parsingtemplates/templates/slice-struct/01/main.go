//range over tpl.gohtml template with a list of sages
package main

import (
	"log"
	"os"
	"text/template"
)

type sage struct {
	Name  string
	Motto string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	buddha := sage{
		Name:  "Buddha",
		Motto: "You are the man",
	}

	lincoln := sage{
		Name:  "lincon",
		Motto: "burning bush",
	}

	captain := sage{
		Name:  "captai",
		Motto: "down periscope",
	}

	sages := []sage{buddha, lincoln, captain}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", sages)
	if err != nil {
		log.Fatal(err)
	}
}
