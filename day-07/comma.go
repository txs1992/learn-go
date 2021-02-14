package main

import (
	"fmt"
	// "bytes"
	"strconv"
)

func main()  {
	fmt.Printf("%s\n", comma(123.4))
}

func comma(f float64) string {
	// var buf bytes.Buffer
	substr := strconv.FormatFloat(f,'G', -1, 64)
	return substr
	// for i := len(s) - 1; i >=0; i-- {

	// }
	
	// n := len(s)
	// if n <= 3 {
	// 	return s;
	// }
	// return comma(s[:n-3]) + "," + s[n-3:]
}
