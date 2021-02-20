package main

import (
	"errors"
	"fmt"
)

func main() {
	// fmt.Println(sum(2, 3))
	// fmt.Println(sum2(-1, -2))
	// fmt.Println(sum3(3, 4))
	// fmt.Println(sum3(-3, -3))
	// fmt.Println(sum2(3, 4))
	// fmt.Println(sum4(1))
	// fmt.Println(sum4(1, 2))
	// fmt.Println(sum4(1, 2, 3, 4, 5))
	// sum5 := func(a, b int) int {
	// 	return a + b
	// }
	// fmt.Println(sum5(1, 2))
	// cl := colsure()
	// fmt.Println(cl())
	// fmt.Println(cl())
	// fmt.Println(cl())
	age := Age(28)
	age.getAge()
	(&age).modify()
	age.getName()

	sm := Age.getAge
	sm(age)
}

func sum(a int, b int) int {
	return a + b
}

func sum2(a int, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("参数 a 或者 b 不能是负数")
	}
	return a + b, nil
}

func sum3(a int, b int) (sum int, err error) {
	if a < 0 || b < 0 {
		sum = 0
		err = errors.New("参数 a 或者 b 不能是负数")
	} else {
		sum = a + b
		err = nil
	}

	return
}

func sum4(params ...int) int {
	sum := 0
	for _, v := range params {
		sum += v
	}
	return sum
}

// 闭包
func colsure() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// Age is uint
type Age uint

// 定义方法
func (age Age) getAge() {
	fmt.Println("the age is：", age)
}

func (age Age) getName() {
	fmt.Println("the name is：", age)
}

func (age *Age) modify() {
	*age = Age(30)
}
