// 练习并发基础 Goroutines 和 Channels
package main

import (
	"fmt"
	"time"
)

func main() {
	// go learnGoroutine()
	// learnChannel()
	// oneWayChannel()
	selectChannel()
}

func selectChannel() {
	downloadFile := func(chanName string) string {
		time.Sleep(time.Second)
		return chanName + ": filePath"
	}

	sendCh := func(ch chan string, chanName string) {
		ch <- downloadFile(chanName)
	}

	// 声明三个 channel
	firstCh := make(chan string)
	secondCh := make(chan string)
	lastCh := make(chan string)

	go sendCh(firstCh, "firstCh")
	go sendCh(secondCh, "secondCh")
	go sendCh(lastCh, "lastCh")

	select {
	case filePath := <-firstCh:
		fmt.Println(filePath)
	case filePath := <-secondCh:
		fmt.Println(filePath)
	case filePath := <-lastCh:
		fmt.Println(filePath)
	}
}

func oneWayChannel() {
	onlySendCh := make(chan<- int)
	onleyReceiveCh := make(<-chan int)
	go sendDataCh(onlySendCh)
	fmt.Println(onlySendCh, onleyReceiveCh)
}

func sendDataCh(send chan<- int) {
	send <- 1
	send <- 2
	send <- 3
}

func learnChannel() {
	// 创建一个有缓冲的通道
	ch := make(chan string, 5)

	// 开启协程并向缓冲通道写入数据
	go chanGoroutine(ch)

	fmt.Println("我是 ch 主协程")

	// 获取有缓冲通道的数据并打印
	v := <-ch
	v2 := <-ch
	fmt.Println("接收到的 channel 中的值为：", v, v2)
	fmt.Println("ch 容量为：", cap(ch), "，元素个数为：", len(ch))

	// 关闭通道
	close(ch)
}

func chanGoroutine(ch chan string) {
	fmt.Println("我是通道写入协程")
	ch <- "chanGoroutine 写入1"
	ch <- "chanGoroutine 写入2"
	ch <- "chanGoroutine 写入3"
	ch <- "chanGoroutine 写入4"
	ch <- "chanGoroutine 写入5"
	ch <- "chanGoroutine 写入6"
}

func learnGoroutine() {
	// 通过 go 关键字加函数/方法来开启协程
	go fmt.Println("MT")
	fmt.Println("我是主协程")
	time.Sleep(time.Second)
}
