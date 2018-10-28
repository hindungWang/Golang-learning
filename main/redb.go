package main

import (
	"net/http"
	"time"
	"fmt"
)



func A(a,b,c int) bool {
	if a*b*c == c*c {
		return true
	}
	return false
}
func B(a,b,c int) bool  {
	if C(a,b) && C(a,c) && C(c,b) {
		return true
	}
	return false
}
func C(a,b int) bool  {
	for b != 0 {
		t := b
		b = a%b
		a = t
	}
	if a == 1{
		return true
	}
	return false
}
func main()  {
	c := &http.Client{
		Timeout: 15 * time.Millisecond,
	}
	_, err := c.Get("http://192.168.99.102:30884")
	if err != nil {
		fmt.Println(err.Error())
	}

}


