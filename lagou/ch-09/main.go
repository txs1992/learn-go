// 练习 sync 包，了解互斥锁
// tips：使用 go run -race 、 go bunild -race 等可以检查代码是否存在资源竞争
package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	sum int = 0
	// mutex   sync.Mutex
	rwMutex sync.RWMutex
)

func main() {
	// moreGoroutine()
	// useWaitGroup()
	// useOnce()
	race()
}

func race() {
	cond := sync.NewCond(&sync.Mutex{})
	var pendWg sync.WaitGroup
	var seloveWg sync.WaitGroup
	pendWg.Add(5)
	seloveWg.Add(6)

	for i := 0; i < 5; i++ {
		go func(num int) {
			defer seloveWg.Done()
			fmt.Println(num, "号已经就位")
			cond.L.Lock()
			pendWg.Done() // 清除等待 wg 计数器
			cond.Wait()   // 等待发令枪响
			fmt.Println(num, "号开始跑...")
			cond.L.Unlock()
		}(i)
	}

	pendWg.Wait()
	// time.Sleep(time.Second)

	go func() {
		defer seloveWg.Done()
		fmt.Println("裁判已经就位，准备发令枪")
		fmt.Println("比赛开始，大家准备跑")
		cond.Broadcast() // 发令枪响
	}()

	seloveWg.Wait() // 防止函数提前返回退出
}

func useOnce() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}

	// 用于等待协程执行完毕
	done := make(chan bool)

	for i := 0; i < 10; i++ {
		go func() {
			// 吧要执行的函数/方法作为参数传递给 once.Do 方法即可
			once.Do(onceBody)
			done <- true
		}()
	}

	for i := 0; i < 10; i++ {
		<-done
	}
}

func useWaitGroup() {
	var wg sync.WaitGroup
	// 添加计数器
	wg.Add(110)

	for i := 0; i < 100; i++ {
		go func() {
			// 计数器值减1
			defer wg.Done()
			mutexAdd(10)
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			fmt.Println("sum 和为：", readSum())
		}()
	}

	fmt.Println("等待中...")
	defer fmt.Println("退出程序...", sum)
	wg.Wait()
}

func moreGoroutine() {
	for i := 0; i < 1000; i++ {
		// go add(1)
		// 运行互斥锁添加操作
		go mutexAdd(1)
	}

	for i := 0; i < 10; i++ {
		go fmt.Println("sum 和为：", readSum())
	}

	// 防止提前退出
	time.Sleep(2 * time.Second)
	fmt.Println("和为：", sum)
}

func readSum() int {
	// 使用读锁
	rwMutex.RLock()
	defer rwMutex.RUnlock()
	b := sum
	return b
}

func mutexAdd(i int) {
	// 使用互斥锁
	rwMutex.Lock()
	// 通过 deferred 解锁
	defer rwMutex.Unlock()
	sum += i
}

func add(i int) {
	sum += i
}
