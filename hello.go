package main

import (
	"fmt"
	"time"
)

func fy() int {

	return 4
}
func main() {
	fmt.Print("hello world!")

	ptr0 := new(int64)
	ptr1 := ptr0
	*ptr0 = 9
	fmt.Println(*ptr1)
	/*for i := 0 ; i < 5; i++ {
		fmt.Println("pre: " , i)
		defer fmt.Println(i)
	}*/
	a := [...]string{0: "no error", 1: "Eio",
		3: "invalid argument"}

	for i, v := range a {
		fmt.Println(i, " -> ", v)
	}
	fmt.Println(a)

	var p *[]int = new([]int)

	*p = make([]int, 5)
	(*p)[0] = 9
	fmt.Println(len(*p))

	var v []int = make([]int, 9)

	fmt.Println(v)
	YSize := 5
	XSize := 3
	// Allocate the top-level slice.
	picture := make([][]uint8, YSize) // One row per unit of y.
	// Loop over the rows, allocating the slice for each row.
	for i := range picture {
		picture[i] = make([]uint8, XSize)
	}
	fmt.Println(picture)

	c := make(chan int) // Allocate a channel.
	// Start the sort in a goroutine; when it completes, signal on the channel.
	go func() {
		time.Sleep(time.Second)
		c <- 1 // Send a signal; value does not matter.
		time.Sleep(time.Second)
		fmt.Println("sent after")
	}()

	if a := <-c; a == 1 { //收到立即
		fmt.Println("coming")
	}

	//panic(fmt.Sprintf("CubeRoot(%s) did not converge", a))

}

func init() {
	fmt.Println("hhh")
	fmt.Println(fmt.Sprint("d"))
}
