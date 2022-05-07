package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] { // NOTE: _ -> ブランク識別子(blank identifier)
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
