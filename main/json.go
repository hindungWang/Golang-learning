package main

import (
	"fmt"
	"encoding/json"
)

type Result struct {

	Status string `json:"status"`
	message string `json:"message"`
}

type Account struct {
	Email string
	password string
	Money float64
}

func main() {
	account := Account{
		Email: "rsj217@gmail.com",
		password: "123456",
		Money: 100.5,
	}
	re := Result{
		Status:"ds",
		message:"dasd",
	}
	fmt.Println(account)
	rs, err := json.Marshal(re)
	if err != nil{

	}
	fmt.Println(string(rs))
}

