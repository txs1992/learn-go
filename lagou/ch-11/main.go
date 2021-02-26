package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// selectTimeout()
	pipeline()
}

func pipeline() {
	coms := buy(10)       // 采购 10 套配件
	phones := build(coms) // 组装 10 部手机
	packs := pack(phones) // 打包它们以便售卖

	// 查看输出结果
	for p := range packs {
		fmt.Println(p)
	}
}

// 扇入函数，吧多个 channel 中的数据发送到一个 channel 中
func merge(ins ...<-chan string) <-chan string {
	var wg sync.WaitGroup
	out := make(chan string)

	// 把一个 channel 中的数据发送到 out 中
	emit := func(in <-chan string) {
		defer wg.Done()
		for c := range in {
			out <- c
		}
	}

	wg.Add(len(ins))

	// 扇入，需要启动多个协程用于处于多个 channel 中的数据
	for _, cs := range ins {
		go emit(cs)
	}

	// 等待所有输入的数据 ins 处理完，再关闭输出 out

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func pack(in <-chan string) <-chan string {
	out := make(chan string)

	go func() {
		defer fmt.Println("打包结束...")
		defer close(out)

		for c := range in {
			out <- fmt.Sprintf("打包(%s)", c)
		}
	}()

	return out
}

func build(in <-chan string) <-chan string {
	out := make(chan string)

	go func() {
		defer fmt.Println("组装结束...")
		defer close(out)

		for c := range in {
			out <- fmt.Sprintf("组装(%s)", c)
		}
	}()

	return out
}

func buy(n int) <-chan string {
	out := make(chan string)

	go func() {
		defer fmt.Println("采购结束...")
		defer close(out)

		for i := 1; i <= n; i++ {
			out <- fmt.Sprintf("配件%d", i)
		}
	}()

	return out
}

// select timeout 模式
func selectTimeout() {
	result := make(chan string)

	go func() {
		// 模拟网络访问
		time.Sleep(8 * time.Second)
		result <- "服务端结果"
	}()

	select {
	case v := <-result:
		fmt.Println(v)
	case <-time.After(5 * time.Second):
		fmt.Println("网络访问超时了")
	}
}
