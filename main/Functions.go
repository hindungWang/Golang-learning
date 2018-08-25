package main

import "fmt"

func plus(a int, b int)   int {
	return a + b
}
func plusPlus(a, b, c int) int {
	return a + b + c
}

func values() (int ,int ) {
	return 0,3
}

func sum(nums ...int) int {
	fmt.Println("nums: ",nums)
	total := 0
	for _, num := range nums{
		total += num
	}
	return total
}
func main()  {
	res := plus(1,2)
	fmt.Println(res)

	fmt.Println(plusPlus(0, 0,'o'))

	fmt.Println(values())
	_, c := values()
	fmt.Println(c)

	fmt.Println(sum(1, 3, 2))
}
