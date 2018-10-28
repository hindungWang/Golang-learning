package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"time"
)
type Port struct {
	PortNum int64 `json:"portnum"`
	Status string `json:"status"`
}
type NetRespon struct {
	Ip      string `json:"ip"`
	Status  string `json:"status"`
	PortOne Port   `json:"portOne"`
	PortTwo Port   `json:"portTwo"`
	PortThree Port `json:"portThree"`
}
type IpList struct {
	Iplist  []NetRespon
}
func (netRespon *NetRespon)netTest(){

	c := &http.Client{
		Timeout: 30 * time.Millisecond,
	}
	_, err := c.Get(fmt.Sprint("http://",netRespon.Ip,":",netRespon.PortOne.PortNum))
	var flag bool
	if err != nil {
		netRespon.PortOne.Status = "error"
		flag = true
	} else {
		netRespon.PortOne.Status = "success"
	}

	_, err = c.Get(fmt.Sprint("http://",netRespon.Ip,":",netRespon.PortTwo.PortNum))
	if err != nil {
		flag = true
		netRespon.PortTwo.Status = "error"
	} else {
		netRespon.PortTwo.Status = "success"
	}

	_, err = c.Get(fmt.Sprint("http://",netRespon.Ip,":",netRespon.PortThree.PortNum))
	if err != nil {
		flag = true
		netRespon.PortThree.Status = "error"
	} else {
		netRespon.PortThree.Status = "success"
	}
	if flag == true {
		netRespon.Status = "error"
	} else {
		netRespon.Status = "success"
	}
}

func ipListNet(list IpList) (b []byte) {
	for i, _ := range list.Iplist {
		list.Iplist[i].netTest()
	}
	b, _ = json.Marshal(list)
	return
}
func main()  {
	var list IpList
	iplist := []string{
		"192.168.99.100",
		"192.168.99.101",
		"192.168.99.102",
	}
	for _,ip := range iplist {
		list.Iplist = append(list.Iplist, NetRespon{
			Ip:ip,
			PortOne: Port{
				PortNum:53,
			},
			PortTwo: Port{
				PortNum:80,
			},
			PortThree: Port{
				PortNum:8800,
			},
		})
	}
	b := ipListNet(list)
	fmt.Println(string(b))
}
