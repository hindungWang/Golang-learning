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

var (
	storage   resource.Quantity
	diskFree  int64
	pvFile    string
	pvConfig  = &pvDefine{}
	mountInfo = &v1.NFSVolumeSource{}
)

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
		result.Message = fmt.Sprintf("mounted volume storage to small,need %s", storage.String())
	}
	b, _ := json.Marshal(result)
	w.Write([]byte(b))
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
func checkNetwork(w http.ResponseWriter, r *http.Request) {
	//ping other pod ip

	w.Write([]byte(fmt.Sprintf("hhh %s", "io")))
}
func main() {
	getPvInfo("pv.yaml")
	http.HandleFunc("/do", checkStorage)
	http.HandleFunc("/net", checkNetwork)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Error("ListenAndServe: ")
	}
}
