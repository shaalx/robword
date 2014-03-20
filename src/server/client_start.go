package main

import (
	"./db"
	. "./tick"
	"fmt"
	"net"
	"os"
)

func data_exchange() {
	conn, e := net.Dial("tcp", "127.0.0.1:1234")
	defer conn.Close()
	checkerr(e)
	for {
		fmt.Println("-------------------------------------------------")
		fmt.Println("...加入游戏...")
		bb := read(&conn)
		Tick(6e9) //60秒游戏时间
		write(&conn, bb)
		fmt.Println("-------------------------------------------------")
		fmt.Println("...本局排名结果...")
		read(&conn)
		Tick(4e9) //查看结果，准备下一次
	}

}

// 读取游戏原始数据
func read(conn *net.Conn) *([]byte) {
	b := make([]byte, 10024)
	n, er := (*conn).Read(b)
	checkerr(er)
	fmt.Println("show", n, "bytes:")
	fmt.Println(string(b[:n]))
	bb := b[:n]
	return &bb
}

/*模拟客户端
从数据库中选择一条记录，作为该玩家的结果，发送至服务器*/
func write(conn *net.Conn, bb *[]byte) (connected bool) {
	(*bb)[0] = 1 //暂时没有用到
	r := db.Selecting()
	bys := r.To_json()
	_, e := (*conn).Write(*bys)
	if e != nil {
		fmt.Println(e)
		return false
	}
	fmt.Println("游戏结果发送至服务器。")
	return true
}

func checkerr(e error) {
	if e != nil {
		fmt.Println(e.Error())
		os.Exit(-1)
	}
}

func main() {
	data_exchange()
}
