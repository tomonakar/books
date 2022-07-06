package main

import (
	"sync"
	"time"
)

// Concurrent applications(平行アプリケーション) を扱う場合、
// 特定のメモリ位置に複数のリソースが潜在的にアクセスすることに
// 対処する必要がある
// これは、「レースコンディション」と呼ばれる

// 具体的には、２つのリソースが同じメモリ位置に同時にアクセスし衝突する問題を
// 回避する必要がある

// Mutexを利用して、複数のリソースが同じメモリに同時にアクセスしようとしても、
// 1つのリソースしか許容しないようロックを行う
// ロックが解除されたら、次のリソースがアクセス可能となる

type Counter struct {
	sync.Mutex
	value int
}

func main() {
	counter := Counter{}

	for i := 0; i < 10; i++ {
		go func(i int) {
			// race counter.Lock()
			counter.value++
			// counter.Unlock()
		}(i)
	}
	time.Sleep(time.Second)
	counter.Lock()
	println(counter.value)
	counter.Unlock()
}

// go run -race <ファイル名> で レースコンディションの検出ができる
// ただし、レース検出器は、静的解析ではなく、実行時に動作するため、
// レース検出機が検出できない潜在的なレースコンディションが設計に存在する可能性がある
