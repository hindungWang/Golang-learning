package main

import (
	"fmt"
)

func main()  {
	c := make(chan int32)
	d := make(chan int32)
	f := make(chan int32)

	j := 1
	go func() {
		for {
			_ = <-d
			fmt.Println(j)
			j += 2
			if j > 99 {
				c <- 1
				return
			}
			c <- 1
		}
	}()
	go func() {
		i := 2
		for {
			_ = <-c
			fmt.Println(i)
			i += 2
			if i > 100 {
				f <- 1
				return
			}
			d <- 1
		}
	}()
	d <- 1
	<- f
}
