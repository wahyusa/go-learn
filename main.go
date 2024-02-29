package main

import (
	"fmt"
	"os"
)

type Classmates struct {
	ID   int
	Name string
}

func main() {
	find()
}

func find() {
	query := os.Args
	if len(os.Args) < 2 {
		fmt.Println(`
      Missing argument.
      Correct example :
      go run main.go 123`)
	} else {
		argument := query[1]
		fmt.Println(argument)
	}
}
