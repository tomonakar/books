package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	channel := make(chan string)

	var waitGroup sync.WaitGroup

	waitGroup.Add(1)
	go func() {
		channel <- "Hello World!"
		println("Finishing goroutine")
		waitGroup.Done()
	}()

	// 1秒waitしてから、チャネルをリッスンしてデータを取得する
	// チャネルもエミッタの向こう側（リスナー側）からの合図を待って、実行を継続する
	time.Sleep(time.Second)
	message := <-channel
	fmt.Println(message)
	waitGroup.Wait()
}

// チャネルとは、Goroutine同士で通信するためのもので、パイプのように
// 一方からデータを送り、もう一方からデータを受け取る
// デフォルト状態では、エミッタはGoroutineはレシーバのGoroutineがデータを受け取るまで、その実行をブロックする
// 受信側のGoroutineも同様に、エミッタがチャネルを通じてデータを送信するまで、ブロックする
// つまり、受動的なリスナー（データ待ち）と受動的なエミッター（リスナー待ち）を持つことができる
