package main

import "fmt"

func main()  {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map", m)

	//遍历
	for k, v := range m{
		fmt.Printf("%s -> %d \n", k, v)
	}

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len: ", len(m))

	delete(m , "k2")
	fmt.Println("map: ", m)

	k,prs := m["k2"]
	fmt.Println("k: ", k)
	fmt.Println("prs: ", prs)

}