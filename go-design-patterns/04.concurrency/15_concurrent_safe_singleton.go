package main

var addCh chan bool = make(chan bool)

// 「整数を受け取るチャネル」を受け取る
var getCountCh chan chan int = make(chan chan int)
var quitCh chan bool = make(chan bool)

// どのパッケージのinit()関数もプログラム実行時に実行されるので、
// このコードを個別に実行する必要っはない
func init() {
	var count int

	go func(addCh <-chan bool, getCountCh <-chan chan int, quitCh <-chan bool) {
		for {
			select {
			case <-addCh:
				count++
			case ch := <-getCountCh:
				ch <- count
			case <-quitCh:
				return
			}
		}
	}(addCh, getCountCh, quitCh)
}

type singleton struct{}

var instance singleton

func GetInstance() *singleton {
	return &instance
}

func (s *singleton) AddOne() {
	addCh <- true
}

func (s *singleton) GetCount() int {
	resCh := make(chan int)
	defer close(resCh)
	getCountCh <- resCh
	return <-resCh
}

func (s *singleton) Stop() {
	quitCh <- true
	close(addCh)
	close(getCountCh)
	close(quitCh)
}
