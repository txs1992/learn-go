package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	// fmt.Printf("循环开始")
	for input.Scan() {
		counts[input.Text()]++
		// 控制退出循环
		if input.Text() == "end" { break }
	}

	// 注意：忽略 input.Err() 中可能的错误
	for line, n := range counts {
		if n > 0 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}