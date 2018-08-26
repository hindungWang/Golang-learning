package main

import "fmt"

func main() {
	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages

	fmt.Println(msg)

	//buffer
	mes := make(chan string, 2)

	mes <- "m1"
	mes <- "m2"

	fmt.Println(<-mes)
	fmt.Println(<-mes)
	//fmt.Println(<-mes)

}
