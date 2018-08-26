package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "resutl 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}

	//第一次sleep 2 秒，只等待 1 秒 所以超时
	//第二次sleep 3 秒，等待了3 秒，所以成功
	//基本发送和接收是阻塞，阻塞的意思是比如
	// write(n) 写n个字节，没写完之前是被阻塞的
}
