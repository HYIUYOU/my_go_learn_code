package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	//s, sep := "", ""
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
		//fmt.Println("os.Args:" + os.Args[i])
	}
	fmt.Println(s)
}
