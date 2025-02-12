package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// CancelDemo()
	// TimeoutDemo()
	ValueDemo()
}

func CancelDemo() {
	// 创建一个可取消的上下文
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		// 监听上下文信号
		select {
		case <-ctx.Done():
			fmt.Println("收到 cancel 信号: ", ctx.Err())
		}
	}()

	// 达成条件，发送取消信号
	cancel()

	// 测试代码可使用忙循环阻塞主线程
	for {
		time.Sleep(1 * time.Second)
	}
}

func TimeoutDemo() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go func() {
		select {
		case <-time.After(3 * time.Second):
			fmt.Println("3 second timeout")
		case <-ctx.Done():
			fmt.Println("接收到 cancel 信号：:", ctx.Err())
		}
	}()

	for {
		time.Sleep(1 * time.Second)
	}
}

func ValueDemo() {
	ctx := context.WithValue(context.Background(), "age", 15)
	if value := ctx.Value("age"); value != nil {
		fmt.Println(value)
	}
}
