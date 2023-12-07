package main

import (
	"fmt"
	"os"
)

func main() {
	for index, sep := range os.Args[1:] {
		fmt.Printf("%s, %d\n", sep, index)
	}
}
