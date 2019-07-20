package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	str := fmt.Sprint(`
	I
	am
	not
	valide
	html
	`)

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error creating file")
	}
	defer nf.Close()

	io.Copy(nf, strings.NewReader(str))

}
