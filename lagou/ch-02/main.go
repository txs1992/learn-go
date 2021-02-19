package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	const (
		num   = 10
		pi    = "3.14"
		isOk  = false
		notOk = "false"
	)

	fmt.Println(strconv.Itoa(num))
	fmt.Println(strconv.Atoi(pi))
	fmt.Println(strconv.FormatBool(isOk))
	fmt.Println(strconv.ParseBool(notOk))
	fmt.Println(strings.HasPrefix(pi, "3"))
}
