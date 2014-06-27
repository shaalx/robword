// test project main.go
package main

import (
	"./gojson"
	"fmt"
	"github.com/astaxie/beego/httplib"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"os"
)

var (
	// req_json *gojson.Js
	str  string
	data MAP
)

type User struct {
	Name      string //路径
	Attribute interface{}
	Children  interface{} //下一个子目录
}

type MAP map[string]string

func init() {
	data = make(MAP, 20)
}

func settings() interface{} {
	kid := User{"839452750", []string{"artistId", "artistUrl", "artistName"}, nil}
	result := User{"results", nil, kid}
	children := User{"product-dv-product", []string{"isAuthenticated", "version"}, result}
	user := User{"storePlatformData", nil, children}
	fmt.Println(user)
	return user
}

func travel(url []string, u interface{}) {
	user, ok := u.(User)
	if ok {
		url = append(url, user.Name)
		if user.Attribute != nil {
			strs, ok := (user.Attribute).([]string) //属性
			if ok {
				for _, i := range strs {
					//fmt.Println(i)
					fmt.Println(url, " 's ", i)
					s := gojson.Json(str).Getpath(url...).Get(i).Tostring()
					// fmt.Println(gojson.Json(str).Getpath("storePlatformData", "product-dv-product", "results", "839452750").Get("artistId").Tostring())
					// s := req_json.Getpath("storePlatformData", "product-dv-product", "results", "839452750", "artwork").Get("bundleId").Tostring()   //数据抓取有问题
					// fmt.Println("data is ", s)
					data[i] = s
				}
			}
		}
		if user.Children != nil {
			travel(url, user.Children)
		}
		fmt.Println(url)
	}

}

func checkerr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}

func run() {
	user := settings()
	url := make([]string, 0)
	travel(url, user)
	fmt.Println(data)
}
func main() {
	request := httplib.Get("https://itunes.apple.com/cn/app/you-zi-xiang-ji-ba-shou-ji/id839452750?mt=8")

	request.Header("Host", "itunes.apple.com")
	request.Header("X-Apple-Store-Front", "143465-19,21 t:native")
	request.Header("Accept", "*/*")
	request.Header("Accept-Language", "zh-cn")
	request.Header("X-Dsid", "1458643138")
	request.Header("Connection", "keep-alive")
	request.Header("Proxy-Connection", "keep-alive")
	request.Header("User-Agent", "AppStore/2.0 iOS/7.1.1 model/iPod5,1 build/11D201 (4; dt:81)")

	//var v interface{}
	//err := (&request).ToJson(&v) //////////加上&
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//fmt.Println(v)

	str, _ = request.String()
	//checkerr(err)
	//fmt.Println(str)

	// req_json = gojson.Json(str)
	// //gojson.Json(str)
	// //fmt.Println(req_json)
	// // urls := []string{"storePlatformData", "product-dv-product", "results", "839452750"}
	// // s := req_json.Getpath(urls...).Get("artistUrl").Tostring()
	// s, v := req_json.Getpath("storePlatformData", "product-dv-product", "results", "839452750", "artwork").ToArray() //数据抓取有问题
	// fmt.Println(s, v)
	// fmt.Println(req_json, str)
	// fmt.Println("...")

	// fmt.Println(req_json)
	// testDB(s)
	run()
}

func testDB(s string) {
	// 连接数据库
	Orm, err := xorm.NewEngine("mysql", "root:root@/appstore?charset=utf8")
	checkerr(err)
	defer Orm.Close()

	// 测试连接
	err = Orm.Ping()
	checkerr(err)
	err = Orm.CreateTables(&User{}) //创建表
	checkerr(err)

	// 插入数据
	// _, err = Orm.Insert(&[]User{User{s}, User{s}})
	// checkerr(err)

	// // 获得数据
	// u := User{}
	// _, err = Orm.Id("416048308").Get(&u)
	// checkerr(err)
	// fmt.Println(u)

	// // 查询
	// us := make([]User, 0)
	// Orm.Sql("select * from user limit 3").Find(&us)
	// fmt.Println(us)

	// _, err = Orm.Count(&u)
	// checkerr(err)
	// fmt.Println(u)
}
