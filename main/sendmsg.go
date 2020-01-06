package main

import (
	"bytes"
	"github.com/ugorji/go/codec"
	"net/http"
	"github.com/vinkdong/gox/log"
)

type Message struct {
	Cpu            string
	Mem            string
	K8sVersion     string
	Area           string
	InstallSucceed bool
}

func main()  {
	msg := Message{
		Cpu:            "1000m",
		Mem:            "16Gi",
		K8sVersion:     "1.8.5",
		Area:           "shanghai/shanghai",
		InstallSucceed: true,
	}
	log.Info("msg ready")
	var buffer = new(bytes.Buffer)
	var handle codec.Handle = new(codec.MsgpackHandle)
	var enc = codec.NewEncoder(buffer, handle)
	var err = enc.Encode(msg)
	if err != nil {
		log.Info(err)
	}
	req, err := http.NewRequest("POST", "http://127.0.0.1:9090", bytes.NewBuffer(buffer.Bytes()))
	if err != nil {
		log.Error(err)
		return
	}
	log.Info("do")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	log.Info(resp.Status)
	defer resp.Body.Close()
}
