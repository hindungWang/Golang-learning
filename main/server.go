package main

import (
	"net/http"
	"github.com/vinkdong/gox/log"
	"github.com/ugorji/go/codec"
	"io/ioutil"
	_ "github.com/go-sql-driver/mysql"
	"sync"
	"fmt"
	"database/sql"
	"time"
	"github.com/prometheus/client_golang/prometheus"
	"strconv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Metrics struct {
	CPU        int64
	Memory     int64
	Province   string
	City       string
	Version    string
	Status     string
	ErrorMsg   []string
	CurrentApp string
	Mux        sync.Mutex
	Ip         string
}

const (
	dbDriverName = "mysql"
	dbUser       = "root"
	dbPassword   = "root"
	dbHost       = "192.168.56.121"
	dbPort       = "3306"
	dbName       = "test"
)

const (
	metricPrefix            = "install_status"
	installSuccessTotalName = metricPrefix + "_success_total"
	installFailTotalName    = metricPrefix + "_fail_total"
)

var (
	dbUrl = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPassword, dbHost, dbPort, dbName)
)

var (
	installSuccess = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Help: "install success counter",
			Name: installSuccessTotalName}, []string{"CPU", "Memory", "Province", "City", "Version", "Status", "CurrentApp"})
	installFail = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Help: "install fail counter",
			Name: installFailTotalName}, []string{"CPU", "Memory", "Province", "City", "Version", "Status", "CurrentApp"})
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}
	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Error(err)
		return
	}
	var handle = new(codec.MsgpackHandle)
	var dec = codec.NewDecoderBytes(data, handle)
	var resMsg Metrics
	err = dec.Decode(&resMsg)
	if err != nil {
		log.Error(err)
		return
	}
	db, err := sql.Open(dbDriverName, dbUrl)
	if err != nil {
		log.Error(err)
		return
	}
	defer db.Close()
	//insert info
	sql := fmt.Sprintf("insert into install_info (cpu, mem, province, city, version, status, current_app, insert_date) VALUES(%d, %d, '%s', '%s', '%s', '%s', '%s', '%s')",
		resMsg.CPU, resMsg.Memory, resMsg.Province, resMsg.City, resMsg.Version, resMsg.Status, resMsg.CurrentApp, time.Now().String())
	res, err := db.Exec(sql)
	if err != nil {
		log.Error(err)
		return
	}
	msg_id, err := res.LastInsertId()
	if err != nil {
		log.Error(err)
		return
	}
	//insert error_msg
	if resMsg.ErrorMsg[0] != "" {
		insertSql := "insert into error_msg (msg_id, msg_context) values "
		for _, str := range resMsg.ErrorMsg {
			insertSql = fmt.Sprintf("%s(%d, '%s'),", insertSql, msg_id, str)
		}
		insertSql = insertSql[0 : len(insertSql)-1]
		stmt, err := db.Prepare(insertSql)
		if err != nil {
			log.Error(err)
			return
		}
		_, err = stmt.Exec()
		if err != nil {
			log.Error(err)
			return
		}
	}
	//prometheus
	resMsg.Mux.Lock()
	if resMsg.Status == "success" {
		installSuccess.WithLabelValues(strconv.FormatInt(resMsg.CPU, 10), strconv.FormatInt(resMsg.Memory, 10), resMsg.Province, resMsg.City, resMsg.Version, resMsg.Status, resMsg.CurrentApp).Inc()
	} else {
		installFail.WithLabelValues(strconv.FormatInt(resMsg.CPU, 10), strconv.FormatInt(resMsg.Memory, 10), resMsg.Province, resMsg.City, resMsg.Version, resMsg.Status, resMsg.CurrentApp).Inc()
	}
	resMsg.Mux.Unlock()
}
func initPrometheus() {
	prometheus.Register(installSuccess)
	prometheus.Register(installFail)
}
func startPrometheus() {
	initPrometheus()
	//conect to mysql and init data
	db, err := sql.Open(dbDriverName, dbUrl)
	if err != nil {
		log.Error(err)
		return
	}
	defer db.Close()
	log.Info("init history data")
	sql := "SELECT cpu,mem,version,province,city,current_app,status,count(*) AS count FROM install_info GROUP BY cpu,mem, version,province,city,current_app,status"
	res, err := db.Query(sql)
	if err != nil {
		log.Error(err)
		return
	}
	defer res.Close()
	for res.Next() {
		var m Metrics
		var count int64
		err := res.Scan(&m.CPU, &m.Memory, &m.Version, &m.Province, &m.City, &m.CurrentApp, &m.Status, &count)
		if err != nil {
			log.Error(err)
			return
		}
		if m.Status == "success" {
			installSuccess.WithLabelValues(strconv.FormatInt(m.CPU, 10), strconv.FormatInt(m.Memory, 10), m.Province, m.City, m.Version, m.Status, m.CurrentApp).Set(float64(count))
		} else {
			installFail.WithLabelValues(strconv.FormatInt(m.CPU, 10), strconv.FormatInt(m.Memory, 10), m.Province, m.City, m.Version, m.Status, m.CurrentApp).Set(float64(count))
		}
	}
	log.Info("start prometheus metrics providing")
}
func main() {
	go startPrometheus()
	log.Info("start server listening")
	http.HandleFunc("/api/v1/metrics", handler)
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Error(err)
	}
}
