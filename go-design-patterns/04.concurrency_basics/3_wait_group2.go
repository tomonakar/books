package main

import (
	"fmt"
	"sync"
)

func main() {
	var wait sync.WaitGroup

	// 5つのGoroutineを実行する
	// Goroutineの実行順序は保証されていないことに注意
	goRoutines := 5
	wait.Add(goRoutines)

	for i := 0; i < goRoutines; i++ {
		go func(goRoutineID int) {
			fmt.Printf("ID:%d: Hello goroutines!\n", goRoutineID)
			wait.Done()
		}(i)
	}
	wait.Wait()
}
