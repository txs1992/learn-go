package main

import "fmt"

func practiceSwitch() {
	switch i := 2; i {
	case 1:
		fmt.Println("1")
		fallthrough
	case 2:
		fmt.Println("2")
		fallthrough
	case 10:
		fmt.Println("10")
	default:
		fmt.Println("default")
	}
}

func practiceEach() {
	for i := 1; i < 3; i++ {
		fmt.Println(i)
	}

	sum := 0
	j := 1

	// 模拟 while
	for j <= 3 {
		sum += j
		j++
	}

	fmt.Println("the sum is:", sum)

	for i := 1; i <= 10; i++ {
		if i <= 5 {
			continue
		}
		fmt.Println("the i is：", i)
	}
}

func main() {
	practiceSwitch()
	practiceEach()
}
