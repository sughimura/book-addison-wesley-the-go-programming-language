package main

import (
	"fmt"
	"os"
)

func main() {
	for i, s := range os.Args[0:] {
		fmt.Println(i, s)
	}
}
