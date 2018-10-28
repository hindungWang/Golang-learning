package main

import (

	"fmt"
)

func t()  {
	s := "sd"
	defer fmt.Println(s)
}
func main(){

	t()
}
