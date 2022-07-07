package main

import "time"

func main() {
	ch := make(chan int)
	go func() {
		ch <- 1
		time.Sleep(time.Second)

		ch <- 2
		// closeでチャネルを閉じることができる
		// closeは受信者ではなく、送信者が行う必要がある
		// 受信者がcloseするとパニックする
		close(ch)
	}()

	// rangeでチャネルがcloseされるまで、チャネルから値を繰り返し受け取る
	for v := range ch {
		println(v)
	}

}
