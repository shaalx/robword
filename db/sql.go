package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"reflect"
)

var (
	ch_db chan *sql.DB
)

func init() {
	ch_db = make(chan *sql.DB, 1)
	// db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/men?charset=utf8&autocommit=true")
	db, err := sql.Open("mysql", "root:1234@/wordmining")
	if err != nil {
		panic(err.Error())
	}
	defer func() {
		ch_db <- db
	}()
}

// 针对特定数据库的查询，不具有通用型，仅限于查询结果返回2个Columns的情况
// 将多次Query的值分别存入一个map数组中
func Query(command ...interface{}) *[]map[string]int {
	// command := make([]interface{}, 3)
	// command[0] = "SELECT * FROM wordmining.sequence;"
	// command[1] = "SELECT * FROM wordmining.word_in_stuff"
	// m := Query(command...)
	// fmt.Println(*m)
	// for i, stuff := range *m {
	// 	if stuff == nil {
	// 		continue
	// 	}
	// 	fmt.Println(i, stuff)
	// }
	db := <-ch_db
	defer func() {
		ch_db <- db
	}()
	Map := make([]map[string]int, 5, 50)
	var word string
	var id int
	for i, cmd := range command {
		if cmd == nil {
			continue
		}
		rows, err := db.Query(interface_to_string(cmd))
		checkErr(err)
		Map[i] = make(map[string]int, 100)
		// fmt.Println(".........")
		for rows.Next() {
			rows.Scan(&word, &id)
			// fmt.Println(word, id)
			Map[i][word] = id
		}
	}

	// fmt.Println(Map)
	return &Map
}

// 分为两部分的sql执行，一次prepare，多次excute
/*参照网上教程：http://www.cnblogs.com/yjf512/archive/2013/01/23/2872577.html
但是，stmt并没有说的NumInput参数，只好自己设计这个参数，即n，表示要绑定的参数个数*/
func Prepare_excute(n int, command ...interface{}) {
	// command := make([]interface{}, 3)
	// command[0] = "INSERT wordmining.word_in_stuff SET word=?,id=?"
	// command[1] = "word--------"
	// command[2] = "3"
	// Prepare_excute(command...)
	db := <-ch_db
	defer func() {
		ch_db <- db
	}()
	stmt, err := db.Prepare(interface_to_string(command[0]))
	checkErr(err)
	for i := 0; i < (len(command)-1)/n; i++ {
		for j := i*n + 1; j < n*(i+1)+1; j++ {
			if command[j] == nil {
				goto H
			}
		}
		_, err = stmt.Exec(command[i*n+1 : n*(i+1)+1]...) //.Query(command[1:]...)
		checkErr(err)
	H:
	}
}

// query 与excute的区别，query可以直接提交一个完整的命令，excute还要有后续的变量（暂时这么理解）
func Excute(command ...interface{}) {
	// command := make([]interface{}, 3)
	// command[0] = "INSERT into wordmining.word_in_stuff values('minings',8),('wordmine',5)"
	// command[1] = "INSERT into wordmining.word_in_stuff values('&&&&&&&',7),('wordmine',6)"
	// Query(command...)
	db := <-ch_db
	defer func() {
		ch_db <- db
	}()
	for _, cmd := range command {
		if cmd == nil {
			continue
		}
		_, er := db.Exec(interface_to_string(cmd))
		checkErr(er)
	}
}

// interface转为字符串类型
// 原因：用到了第三方的库（mysql），传递的参数有要求是字符串类型
func interface_to_string(i interface{}) string {
	return reflect.ValueOf(i).String()
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
}
