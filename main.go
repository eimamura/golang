package main

import (
	"fmt"
	"time"
)

func main() {
	// 新しいGoroutineを起動
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("Goroutine:", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// 重い処理のシミュレーション（メインGoroutine）
	for i := 0; i < 5; i++ {
		fmt.Println("Main:     ", i)
		time.Sleep(100 * time.Millisecond)
	}
	
	fmt.Println("Done!")
}
