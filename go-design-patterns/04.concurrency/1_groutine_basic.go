package main

import "time"

func main() {

	// go routineを実行する
	go helloWorld()
	// helloWorld関数を実行する前にmain関数が終了しないようwaitを掛けている
	time.Sleep(time.Second)

	// 無名関数
	go func() {
		println("Hello World2")
	}()
	time.Sleep(time.Second)

	// 無名関数の引数指定
	go func(msg string) {
		println(msg)
	}("Hello World3")
	time.Sleep(time.Second)

	// 関数定義
	messagePrinter := func(msg string) {
		println(msg)
	}

	go messagePrinter("HelloWorld4")
	go messagePrinter("Hello goroutine")
	time.Sleep(time.Second)

}

func helloWorld() {
	println("Hello World!")
}
