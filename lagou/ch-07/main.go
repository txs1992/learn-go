package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// 自定义错误结构体
type customError struct {
	errorCode    int
	errorMessage string
}

type customErrorWrapper struct {
	error
	cerr customError
	msg  string
}

func main() {
	// testError()
	// testAdd()
	// errorAssert()
	// testErrorWrapper()
	// testErrorf()
	// testAs()
	// e := errors.New("元素错误 e")
	// w := fmt.Errorf("Wrap 了一个错误: %w", e)
	// fmt.Println(w)
	// fmt.Println(errors.Unwrap(w))
	// fmt.Println(errors.Is(w, e))
	// fmt.Println(readFile("../ch-06/walk-run.go"))
	// moreDefer()
}

func (cew *customErrorWrapper) Error() string {
	return cew.Error() + cew.msg
}

func (ce *customError) Error() string {
	return ce.errorMessage
}

func moreDefer() {
	defer fmt.Println("First defer")
	defer fmt.Println("Second defer")
	defer fmt.Println("Last defer")
	fmt.Println("函数自身代码")
}

func readFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)

	if err != nil {
		return nil, err
	}
	defer f.Close()

	return nil, err
}

func testAs() {
	sum, err := add(-1, 2)
	var ce *customError

	if errors.As(err, &ce) {
		fmt.Println("错误代码为:", ce.errorCode, "，错误信息为：", ce.errorMessage)
	} else {
		fmt.Println(sum)
	}
}

func testErrorWrapper() {
	_, err := strconv.Atoi("A")
	errWrap := customErrorWrapper{err, customError{1, "参数 a 或者 b 不能为负数"}, "外层包裹容器信息"}
	fmt.Println(errWrap)
}

func testAdd() {
	sum, err := add(-1, 2)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sum)
	}
}

// 通过断言获取详细错误信息
func errorAssert() {
	sum, err := add(-1, 2)

	if cm, ok := err.(*customError); ok {
		fmt.Println("错误代码为:", cm.errorCode, "，错误信息为：", cm.errorMessage)
	} else {
		fmt.Println(sum)
	}
}

func add(a, b int) (int, error) {
	if a < 0 || b < 0 {
		// return 0, errors.New("参数 a 或者 b 不能为负数")
		return 0, &customError{1, "参数 a 或者 b 不能为负数"}
	} else {
		return a + b, nil
	}
}

func testError() {
	i, err := strconv.Atoi("A")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}
}
