package main

import (
	"time"
)

// channelの方向(送信or受診）を指定し、誤った使い方をした場合にコンパイラで検出してもらうことができる

/*
* チャネルのおさらい：
* channel := make(chan int)
* channel <- 10 チャネルに送信
* message := <- channel チャネルから受診
 */
func main() {
	channel := make(chan string, 1)

	// ch chan<-stringの指定で、このチャネルは、入力のみ受け付ける型となる
	// chをレシーバーチャネルとして利用すると、コンパイルエラーとなる
	go func(ch chan<- string) { // これは送信専用のチャネルの型
		ch <- "Hello World"
		println("Finishing goroutine")
	}(channel)

	time.Sleep(time.Second)

	receivingCh(channel)

	// message := <-channel
	// fmt.Println(message)
}

// これは受診専用のチャネルの型
func receivingCh(ch <-chan string) {
	msg := <-ch
	println(msg)
}
