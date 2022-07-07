package main

import (
	"fmt"
	"time"
)

func main() {
	// channel作成時に、２番目の引数にバッファ長を指定することで
	// buffered channelを作成できる
	channel := make(chan string, 1)

	go func() {
		channel <- "Hello World 1"
		channel <- "Hello World 2"
		println("Finishing goroutine")
	}()

	time.Sleep(time.Second)

	// buffered channelの場合、チャネルがいっぱいなら送信操作は、別のGoroutineによる受診操作によって空きが出るまで待たされる（ブロックされる）
	// buffered chanelに空きがあれば、送信操作は待たされない。
	// buffered channelが「空」の場合、受診操作は別のGoroutineが値を送信するまで待たされる（これは普通のchannelの挙動と同じ）

	// よって、このコードの場合、バッファは１なので、2つ目のメッセージの送信はブロックされる
	manage := <-channel

	// そして、２つ目のメッセージが送信する前にmain関数が終了する
	fmt.Println(manage)
}
