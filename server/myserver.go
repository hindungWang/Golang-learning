package main

import (
	"encoding/json"
	"fmt"
	"github.com/vinkdong/gox/log"
	"io/ioutil"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/util/yaml"
	"net/http"
	"os"
	"syscall"
	"time"
)

type pvDefine struct {
	Version  string
	Metadata Metadata
	Spec     v1.PersistentVolumeSpec
}

type Metadata struct {
	Name string
}

type DiskStatus struct {
	All  uint64 `json:"all"`
	Used uint64 `json:"used"`
	Free uint64 `json:"free"`
}

type Result struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type PortStatus struct {
	Status string `json:"status"`
	Port   int64  `json:"port"`
}

type Port struct {
	PortNum int64  `json:"portnum"`
	Status  string `json:"status"`
}
type NetRespon struct {
	Ip        string `json:"ip"`
	Status    string `json:"status"`
	PortOne   Port   `json:"portOne"`
	PortTwo   Port   `json:"portTwo"`
	PortThree Port   `json:"portThree"`
}
type IpList struct {
	Iplist []NetRespon
}

var (
	storage   resource.Quantity
	diskFree  int64
	pvFile    string
	pvConfig  = &pvDefine{}
	mountInfo = &v1.NFSVolumeSource{}
)

func (netRespon *NetRespon) netTest(c *http.Client) {
	_, err := c.Get(fmt.Sprint("http://", netRespon.Ip, ":", netRespon.PortOne.PortNum))
	var flag bool
	if err != nil {
		netRespon.PortOne.Status = "error"
		flag = true
	} else {
		netRespon.PortOne.Status = "success"
	}

	_, err = c.Get(fmt.Sprint("http://", netRespon.Ip, ":", netRespon.PortTwo.PortNum))
	if err != nil {
		flag = true
		netRespon.PortTwo.Status = "error"
	} else {
		netRespon.PortTwo.Status = "success"
	}

	_, err = c.Get(fmt.Sprint("http://", netRespon.Ip, ":", netRespon.PortThree.PortNum))
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
	c := &http.Client{
		Timeout: 30 * time.Millisecond,
	}
	for i, _ := range list.Iplist {
		list.Iplist[i].netTest(c)
	}
	b, _ = json.Marshal(list)
	return
}
func DiskUsage(path string) (disk DiskStatus) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		return
	}
	disk.All = fs.Blocks * uint64(fs.Bsize)
	disk.Free = fs.Bfree * uint64(fs.Bsize)
	disk.Used = disk.All - disk.Free
	return
}
func getPvInfo(pvFile string) {
	var (
		data []byte
		err  error
	)
	if pvFile != "" {
		data, err = ioutil.ReadFile(pvFile)
		if err != nil {
			log.Error("read install file error")
			os.Exit(127)
		}
		data2, err := yaml.ToJSON(data)
		if err != nil {
			panic(err)
		}
		json.Unmarshal(data2, pvConfig)
	}
	//pv storage 大小
	request := pvConfig.Spec.Capacity
	storage = request["storage"]
	//NFS服务器信息
	mountInfo = pvConfig.Spec.NFS
}
func checkPermission() (re int) {
	re = 0
	file := fmt.Sprint(mountInfo.Path, "/test_file")
	f, err := os.Create(file)
	f.Close()
	if err != nil {
		log.Error("create file error")
		re = 1 //没有创建权限
	}
	errs := os.Remove(file)
	if errs != nil {
		log.Error("delete file error")
		re = 2 //没有删除权限
	}
	return
}
func portResponOne(w http.ResponseWriter, r *http.Request) {
	portStatus := PortStatus{
		Status: "success",
		Port:   80,
	}
	b, _ := json.Marshal(portStatus)
	w.Write([]byte(b))
}
func portResponTwo(w http.ResponseWriter, r *http.Request) {
	portStatus := PortStatus{
		Status: "success",
		Port:   53,
	}
	b, _ := json.Marshal(portStatus)
	w.Write([]byte(b))
}
func portResponThree(w http.ResponseWriter, r *http.Request) {
	portStatus := PortStatus{
		Status: "success",
		Port:   8800,
	}
	b, _ := json.Marshal(portStatus)
	w.Write([]byte(b))
}
func checkStorage(w http.ResponseWriter, r *http.Request) {
	diskFree = int64(DiskUsage(mountInfo.Path).Free)
	result := Result{"success", ""}
	if checkPermission() == 1 {
		result.Status = "error"
		result.Message = "No create permission"
	} else if checkPermission() == 2 {
		result.Status = "error"
		result.Message = "No delete permission"
	} else if storage.CmpInt64(diskFree) > 0 {
		result.Status = "error"
		result.Message = fmt.Sprintf("mounted volume storage too small,need %s", storage.String())
	}
	b, _ := json.Marshal(result)
	w.Write([]byte(b))
}
func checkNetwork(w http.ResponseWriter, r *http.Request) {
	var iplist IpList
	list := []string{
		"192.168.99.100",
		"192.168.99.101",
		"192.168.99.102",
	}
	for _, ip := range list {
		iplist.Iplist = append(iplist.Iplist, NetRespon{
			Ip: ip,
			PortOne: Port{
				PortNum: 53,
			},
			PortTwo: Port{
				PortNum: 80,
			},
			PortThree: Port{
				PortNum: 8800,
			},
		})
	}
	b := ipListNet(iplist)
	w.Write(b)
}
func main() {
	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/", portResponOne)
		http.ListenAndServe(":80", mux)
	}()
	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/", portResponTwo)
		http.ListenAndServe(":53", mux)
	}()
	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/", portResponThree)
		http.ListenAndServe(":8800", mux)
	}()
	getPvInfo("pv.yaml")
	http.HandleFunc("/net", checkNetwork)
	http.HandleFunc("/do", checkStorage)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Error("ListenAndServe")
	}
}
