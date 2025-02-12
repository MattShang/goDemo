package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 创建一个通道来接收信号
	sigChan := make(chan os.Signal, 1)

	// 注册要捕获的信号
	// SIGINT - Ctrl + C
	// SIGTERM - 系统终止 [如 kill 命令是 terminated ]
	// syscall.SIGKILL 命令无法捕获，因为它们是强制终止
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM) // ✅
	// signal.Notify(sigChan, os.Interrupt)  // ❌

	// 新开一个 goroutine 处理信号
	go func() {
		// 等待信号
		sig := <-sigChan
		fmt.Printf("接收到信号: %+v\n", sig)

		// 执行清理操作
		fmt.Println("开始执行清理操作...")
		time.Sleep(2 * time.Second)
		fmt.Println("清理完成，程序退出")

		// 退出程序
		os.Exit(0)
	}()

	// 模拟主程序运行
	fmt.Printf("PID: %d 程序运行中...，按 Ctrl+C 退出\n", os.Getpid())

	// 使用 select 阻塞主程序
	select {}
}
