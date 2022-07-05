package main

import (
	"fmt"
	"sync"
)

func main() {
	// 最初に WaitGroupを定義する
	var wait sync.WaitGroup
	// 次に、Goroutineを実行する前に、wait.Add(1)メソッドで、
	// goroutineを１つ実行するまで待つよう依頼
	wait.Add(1)

	go func() {
		fmt.Println("Hello World!")
		// Goroutineの終了を宣言
		wait.Done()
	}()
	// WaitGroupの実行を宣言
	wait.Wait()
}
