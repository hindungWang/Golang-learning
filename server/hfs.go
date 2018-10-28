package main

import (
	"net/http"
	"log"
)

func main()  {
	http.Handle("/", http.FileServer(http.Dir("/u01/kubeadm-offline-package/yum/kubernetes")))
	err := http.ListenAndServe(":9080", nil)
	if err != nil {
		log.Println(err)
	}
}
/*
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build main.go
 */