package main

import (
	_ "github.com/go-sql-driver/mysql"
	"os"
	"fmt"
)

func main()  {
	mysqlHost    := os.Getenv("mysqlHost")
	mysqlPost    := os.Getenv("mysqlPost")
	mysqlName    := os.Getenv("mysqlName")
	mysqlPwd     := os.Getenv("mysqlpwd")
	//databaseName := os.Getenv("databaseName")
	ConnetInfo   := fmt.Sprint( mysqlName,":",mysqlPwd,"@tcp(",mysqlHost,":",mysqlPost,")/?charset=utf8&timeout=3s")
	fmt.Println(ConnetInfo)

	/*fmt.Println("FOO:", os.Getenv("mysqlHost"))
	fmt.Println("BAR:", os.Getenv("BAR"))
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}
	db, err := sql.Open("mysql", "root:123456@tcp(192.168.99.100:3306)/?charset=utf8&timeout=3s")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	err = db.Ping()

	if err != nil {
	}

	_,err = db.Exec("DROP DATABASE IF EXISTS T")
	_,err = db.Exec("CREATE DATABASE T")
	if err != nil {
		panic(err)
	}*/
}
