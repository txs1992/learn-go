package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type key string

func main() {
	// runWatchDog()
	// useContextRunWatchDog()
	useContextValue()
}

func useContextValue() {
	var wg sync.WaitGroup
	ctx, stop := context.WithCancel(context.Background())

	const userIDKey key = "userId"

	valCtx := context.WithValue(ctx, userIDKey, "007")

	wg.Add(1)

	go func() {
		defer wg.Done()
		getUser(valCtx, userIDKey)
	}()

	time.Sleep(2 * time.Second)
	stop()
	wg.Wait()
}

func useContextRunWatchDog() {
	var wg sync.WaitGroup
	ctx, stop := context.WithCancel(context.Background())

	wg.Add(3)

	for i := 1; i < 4; i++ {
		go func(i int) {
			defer fmt.Println("协程运行结束")
			defer wg.Done()
			contextWatchDog(ctx, fmt.Sprintf("【监控狗 %d 号】", i))
		}(i)
	}

	time.Sleep(2 * time.Second)
	fmt.Println("停止 context...")
	stop()
	wg.Wait()
	fmt.Println("关闭 waitGroup...")
}

func runWatchDog() {
	stopCh := make(chan bool) //用来停止监控狗
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		watchDog(stopCh, "【监控狗 1 号】")
	}()

	time.Sleep(3 * time.Second) // 先让监控狗监控 3 秒
	stopCh <- true
	wg.Wait()
}

func watchDog(stopCh chan bool, name string) {
	// 开启 for select 循环，一直在后台监控
	for {
		select {
		case <-stopCh:
			fmt.Println(name, "停止指令已收到，立即停止")
			return
		default:
			fmt.Println(name, "正在监控...")
		}
		time.Sleep(time.Second / 2)
	}
}

func contextWatchDog(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "停止指令已收到，立即停止")
			return
		default:
			fmt.Println(name, "正在监控...")
		}
		time.Sleep(time.Second / 2)
	}
}

func getUser(ctx context.Context, userIDKey key) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("【获取用户】", "协程退出")
			return
		default:
			userID := ctx.Value(userIDKey)
			fmt.Println("【获取用户】", "用户 ID 为：", userID)
			time.Sleep(time.Second)
		}
	}
}
