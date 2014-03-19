package main

import (
	// "bytes"
	. "./tick"
	"bytes"
	"fmt"
	"net"
	"os"
)

func chant() {
	conn, e := net.Dial("tcp", "127.0.0.1:1234")
	defer conn.Close()
	checkerr(e)
	read(&conn)
	Tick(10e9)
	write(&conn)
	read(&conn)
	Tick(6e9)

}

func read(conn *net.Conn) {
	b := make([]byte, 1024)
	n, er := (*conn).Read(b)
	checkerr(er)
	fmt.Println(n, "bytes:")
	fmt.Println(string(b[:n]))
}

func write(conn *net.Conn) (connected bool) {
	a := "我的游戏结果。"
	bs := bytes.NewBufferString(a)
	b := bs.Bytes()
	_, e := (*conn).Write(b)
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
	chant()
}
