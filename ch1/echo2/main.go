package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:len(os.Args)] {
		s += sep + arg
	}
	fmt.Println(s)
}
