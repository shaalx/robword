package main

import (
	// "crawl_json/gojson"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/httplib"
	// _ "github.com/go-sql-driver/mysql"
	// "github.com/go-xorm/xorm"
	"labix.org/v2/mgo"
	// "labix.org/v2/mgo/bson"
	"os"
	"time"
)

type Design struct {
	Stuff string
}

// var (
// 	IsDrop = true
// )

func main() {
	request := httplib.Get("https://itunes.apple.com/cn/app/you-zi-xiang-ji-ba-shou-ji/id839452750?mt=8")

	request.Header("Host", "itunes.apple.com")
	request.Header("X-Apple-Store-Front", "143465-19,21 t:native")
	request.Header("Accept", "*/*")
	request.Header("Accept-Language", "zh-cn")
	request.Header("X-Dsid", "1458643138")
	request.Header("Connection", "keep-alive")
	request.Header("Proxy-Connection", "keep-alive")
	request.Header("Design-Agent", "AppStore/2.0 iOS/7.1.1 model/iPod5,1 build/11D201 (4; dt:81)")

	str, err := request.String()
	checkerr(err)
	// fmt.Println(str)
	storeIntoMongoDB(str)

	// req_json := gojson.Json(str)
	// urls := []string{"storePlatformData", "product-dv-product", "results", "839452750"}
	// // s := req_json.Getpath(urls...).Get("artistUrl").Tostring()
	// s := req_json.Getpath(urls...).Get("artwork").Tostring()
	// fmt.Println(req_json)
	// storeIntoDB(s)
}

// func storeIntoDB(s string) {
// 	// 连接数据库
// 	Orm, err := xorm.NewEngine("mysql", "root:root@/appstore?charset=utf8")
// 	checkerr(err)
// 	defer Orm.Close()

// 	// 测试连接
// 	err = Orm.Ping()
// 	checkerr(err)
// 	err = Orm.CreateTables(&Design{}) //创建表
// 	checkerr(err)

// 	// 插入数据
// 	n, err := Orm.Insert(&[]Design{Design{s}})
// 	checkerr(err)
// 	// fmt.Println(s)
// 	fmt.Println(n, " affected.")
// }

func checkerr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}

type Data struct {
	Value interface{} `json:"value"`
	Date  time.Time   `json:"date"`
}

func storeIntoMongoDB(s string) {
	session, err := mgo.Dial("192.168.199.240:27017")
	// session, err := mgo.Dial("127.0.0.0:27017")
	checkerr(err)
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	// Collection
	c := session.DB("appstore").C("app")
	c.Count()
	var data Data
	// data.Value = s
	json.Unmarshal([]byte(s), &data.Value)
	data.Date = time.Now()
	fmt.Println(data)
	d, err := json.Marshal(data)
	checkerr(err)
	var v interface{}
	json.Unmarshal(d, &v)
	fmt.Println(v)
	c.Insert(v)

	// fmt.Println("inter All.")
}
