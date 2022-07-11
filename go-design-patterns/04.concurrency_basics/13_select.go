package main

import "time"

/*
* Select文はチャネルの送受信操作を多重化できる
* select文は上から順に評価されない
* →　チャネルの送受信が実行可能かを判断し、可能な場合に実行される
* どのチャネルも実行可能で無い場合は、defaultが実行される
* 複数のチャネルが実行可能な場合は、どれか１つのcaseがラウンダむに実行される
* select文では1回しかcaseが実行されないので、複数回実行したい場合はforを使う
 */

// チャネル送信メソッド
func sendString(ch chan<- string, s string) {
	ch <- s
}

// チャネル受診メソッド
func receiver(helloCh, goodbyCh <-chan string, quitCh chan<- bool) {
	// forで無限ループ化し、helloChとgoodbyChをリッスンし続ける
	for {
		select {
		case msg := <-helloCh:
			println(msg)
		case msg := <-goodbyCh:
			println(msg)
			// 2秒間他のチャネルを受診しなかった場合にこのケースに入る
			// selectはハンドルごとに再開されるので、これはタイムアウトとして働く
		case <-time.After(time.Second * 2):
			println("Nothing received in 2 seconds. Exiting")
			quitCh <- true
			break
		}
	}
}

func main() {
	helloCh := make(chan string, 1)
	goodbyeCh := make(chan string, 1)
	quitCh := make(chan bool, 1)
	go receiver(helloCh, goodbyeCh, quitCh)

	go sendString(helloCh, "Hello!")

	time.Sleep(time.Second)

	go sendString(goodbyeCh, "goodbye!")
	<-quitCh

}
