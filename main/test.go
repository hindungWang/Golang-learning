package main

import (
	"fmt"
)

func main()  {
	s := make([]int, 0, 3)
	fmt.Printf("%p\n",&s)
	for i := 0; i < 5; i++ {
		s = append(s, i)
		fmt.Printf("cap %v, len %v, %p\n", cap(s), len(s), s)
	}
}
