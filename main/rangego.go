package main

import "fmt"

func main()  {
	arr := []int{10, 20, 30}
	for _, num := range arr {
		num++
	}
	fmt.Println(arr)
}
