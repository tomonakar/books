package main

import "sync"

type Counter struct {
	sync.Mutex
	value int
}

func main() {
	counter := Counter{}

	// レース検出器では検出絵できない
	// for i := 0; i < 1; i++ {

	// レース検出器で検出可能。ただし、GOMAXPROCS=1とすると検出できなくなる
	for i := 0; i < 2; i++ {
		go func(i int) {
			counter.value++
		}(i)
	}
}

// -raceは、レースコンディションを実行時に引き起こす場合には検出できるが、
// 設計上、レースコンディションを引き起こす可能性があるコードの検出はできない
