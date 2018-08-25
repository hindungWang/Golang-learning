package main

import "fmt"

func main()  {
	var a = "init"
	{gf := 9;fmt.Println(gf)}

	fmt.Println(a)


	var b, c int = 1, 2;
	fmt.Println(b, c)
	fmt.Println(&b)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "sort"
	fmt.Println(f)
	if (b + c == 3 || (fsd() == 3)) {
		fmt.Println("i")
	}
}
func fsd() int  {
	fmt.Println("dsd")
	return 4
}
