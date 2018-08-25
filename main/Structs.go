package main

import "fmt"

type person struct {
	name string
	age int
}

func main()  {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name :"Alice", age : 19})
	fmt.Println(person{name : "Fred"})
	fmt.Println(&person{name : "Ann", age : 23})

	s := person{"sen", 24}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 90

	fmt.Println(s.age)

}