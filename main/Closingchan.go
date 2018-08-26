package main

import (
	"fmt"
)

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("recevied job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// mian
	for i := 1; i <= 3; i++ {
		jobs <- i
		fmt.Println("sent job", i)
		//time.Sleep(time.Second)
	}
	close(jobs)

	fmt.Println("sent all jobs")
	<-done

	// the more value will be
	// false if jobs has been
	// closed and all values
	// in the channel have
	// already been received.
	// We use this to notify
	// on done when we’ve worked
	// all our jobs.
}
