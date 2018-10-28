package main

import (
	"os"
	"fmt"
)

func createFile(path string) (err error) {
	 _, err = os.Stat(path)
	if os.IsNotExist(err) {
		file, err := os.Create(path)
		defer file.Close()
		if err != nil {
			return err
		}
	}
	return
}
func main()  {
	err := createFile("tst")
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("create filed")
		return
	}
	err = os.Remove("tst")
	if err != nil {
	}
	fmt.Println("create success")
}
