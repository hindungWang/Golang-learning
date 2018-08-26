package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var ops uint64
	var op uint64

	for i := 0; i < 50; i++ {
		go func() {
			for {
				atomic.AddUint64(&ops, 1)
				op++
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println(ops)
	fmt.Println(opsFinal)

	fmt.Println(op)
}
