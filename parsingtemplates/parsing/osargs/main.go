package main

import (
	"fmt"
	"os"
)

func main() {

	firstName := os.Args[1]
	lastName := os.Args[2]
	fmt.Println("first name is", firstName)
	fmt.Println("last name is", lastName)
}
