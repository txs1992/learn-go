package main

import (
	"os"
	"fmt"
)

func main() {
	fmt.Println(os.Args[0])

	for i := 0; i < len(os.Args); i++ {
		fmt.Println("line ", i, os.Args[i])
	}
}