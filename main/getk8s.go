package main

import "fmt"

func main()  {
	ipmap := make(map[string]string)
	ipmap["s"] = "ds"
	if ipmap["ds"] == "" {
		fmt.Println("d")
	}
	b := true
	if b {
		fmt.Println(b)
	}
}