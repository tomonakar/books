package main

import "fmt"

// テレビのチャンネルは、スタジオからの放送を何百万ものテレビ（受信機）に接続する
// Goのチャンネルも、テレビのチャンネルと同じように動く
// 複数のgoroutineがエミッターとして働き、またレシーバとして働く
// また、goのチャネルは、デフォルトで何かが受信されまで、goroutineの実行をブロックする

func main() {
	// チャネル作成
	channel := make(chan string)
	go func() {

		// <- でチャネルにメッセージを渡している
		channel <- "Hello World!"
	}()

	// ここでmessage変数にchannelが代入されるまで、
	// goroutineの実行がブロックされる
	// そのため、main関数が先に終了することがない
	// なので、main関数とGoroutineを同期するためにWait.groupを使用する必要がない
	// では、その逆はどうだろうか？　Goroutineがメッセージを受信したときに、受信者がいない場合はどうなるか？次のファイルで検証する
	message := <-channel
	fmt.Println(message)
}
