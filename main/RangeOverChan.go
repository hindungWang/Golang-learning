package main

import (
	"fmt"
)

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "tow"
	/*go func() {
		mes := <- queue
		fmt.Println(mes)
	}()*/
	close(queue)
	/*mes := <- queue
	fmt.Println(mes)*/
	//time.Sleep(time.Second)
	for elem := range queue {
		fmt.Println(elem)
	}
}
