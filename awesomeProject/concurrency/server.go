package main

import (
	"net"
	"time"
)

// 获取地址
func handler(c net.Conn, ch chan string) {
	addr := c.RemoteAddr().String()
	// 地址传进通道ch
	ch <- addr
	time.Sleep(100 * time.Millisecond)
	c.Write([]byte("ok"))
	c.Close()
}

func logger(wch chan int, results chan int) {
	for {
		// 从通道取值
		data := <-wch
		data++
		// 给通道值
		results <- data
	}
}

func parse(results chan int) {
	for {
		// 持续丢弃
		<-results
	}
}

func pool2(ch chan string, n int) {
	wch := make(chan int)
	results := make(chan int)
	for i := 0; i < n; i++ {
		// 给n个协程去把通道值传递
		go logger(wch, results)
	}
	go parse(results)
	for {
		addr := <-ch
		l := len(addr)
		wch <- l
	}
}

func server(l net.Listener, ch chan string) {
	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}
		go handler(c, ch)
	}
}

func main() {
	l, err := net.Listen("tcp", ":5000")
	if err != nil {
		panic(err)
	}
	ch := make(chan string)
	go pool2(ch, 4)
	go server(l, ch)
	time.Sleep(10 * time.Second)
}