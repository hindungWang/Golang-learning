package main

import (
	"fmt"
)

func qsort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	mid := arr[0]
	begin, end := 0, len(arr)-1
	for i := 1; i <= end; {
		if (arr[i] < mid) {
			arr[i], arr[begin] = arr[begin], arr[i]
			begin++
			i++
		} else {
			arr[i], arr[end] = arr[end], arr[i]
			end--
		}
	}
	fmt.Println("** ", arr, "**\n")
	qsort(arr[:begin])
	qsort(arr[begin+1:])
}

func Qsort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	mid := arr[0]
	begin, end := 0, len(arr)-1

	for i := 1; i <= end; {
		if mid > arr[i] {
			arr[i], arr[begin] = arr[begin], arr[i]
			i++
			begin++
		} else {
			arr[i], arr[end] = arr[end], arr[i]
			end--
		}
	}
	Qsort(arr[:begin])
	Qsort(arr[begin+1:])

}
func main() {
	arr := []int{1, 2, 4, 5, 2, 3, 5, 34, 33, -1}

	//sort.Ints(arr)
	Qsort(arr)
	fmt.Println(arr)
	a := []int{1, 3}
	a[0], a[1] = a[1], a[0]
	fmt.Println(a)
}
