package main

import (
	"fmt"
	"strings"
	"sync"
)

var wait sync.WaitGroup

func main() {
	wait.Add(1)

	// This is Callback Hell!!!
	toUpperAsync("Hello Callbacks!", func(v string) {
		toUpperAsync(fmt.Sprintf("Callback: %s\n", v), func(v string) {
			fmt.Printf("Callback within %s", v)
			// 外側のコールバックは内側のコールバックの前に実行されるので、
			// waitGroupは１つで良い
			wait.Done()
		})
	})

	println("Waiting async response...")
	wait.Wait()

}

func toUpperAsync(word string, f func(string)) {
	go func() {
		f(strings.ToUpper(word))
	}()
}
