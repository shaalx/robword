package main

import (
	. "../db"
	. "../protocal"
	. "./tick"
	"bytes"
	"fmt"
	"net"
	"os"
	"time"
)

var (
	data_chan chan *([]byte)
	result    chan *([]byte)
)

func listener() {
	listener, e := net.Listen("tcp", ":1234")
	system_error(e)
	go sync() //时间同步
	for {
		conn, er := listener.Accept()
		system_error(er)
		go data_exchange(&conn)
	}
}

func checkerr(e error) {
	if e != nil {
		fmt.Println(e.Error())
	}
}

func system_error(e error) {
	if e != nil {
		fmt.Print(e.Error())
		os.Exit(-1)
	}
}

func data_exchange(conn *net.Conn) {
	defer (*conn).Close()
	read(conn)
	if !write_data(conn) {
		return
	}
	fmt.Println("1分钟后收集所有玩家游戏结果.")
	read2(conn)
	fmt.Println("统计结果。")
	if !write_reslut(conn) {
		return
	}
}

func write_data(conn *net.Conn) (connected bool) {
	b := <-data_chan
	data_chan <- b
	_, e := (*conn).Write(*b)
	if e != nil {
		fmt.Println(e)
		return false
	}
	fmt.Println("数据已经发送至玩家")
	return true
}

func write_reslut(conn *net.Conn) (connected bool) {
	b := <-result
	result <- b
	_, e := (*conn).Write(*b)
	if e != nil {
		fmt.Println(e)
		return false
	}
	fmt.Println("结果已经发送至玩家")
	return true
}

func read(conn *net.Conn) {
	fmt.Println("获取玩家基本信息", (*conn).LocalAddr())
}

func read2(conn *net.Conn) {
	fmt.Println("获取玩家游戏结果..数据：")
	b := make([]byte, 1024)
	n, er := (*conn).Read(b)
	checkerr(er)
	fmt.Println(n, "bytes:")
	fmt.Println(string(b[:n]))
}

func data_init() {
	// a := "游戏数据"
	// bs := bytes.NewBufferString(a)
	// b := bs.Bytes()
	clear_tran()
	// data_chan <- b
	var r *Resultwords
	r = Selecting()
	bys := r.To_json()
	data_chan <- bys
}

func result_formate() {
	a := "游戏结果排名数据：*****"
	bs := bytes.NewBufferString(a)
	b := bs.Bytes()
	clear_result()
	result <- &b
}

func clear_tran() {
	if len(data_chan) == 1 {
		<-data_chan
	}
}

func clear_result() {
	if len(result) == 1 {
		<-result
	}
}

func sync() {
	tick := time.NewTicker(1e9)
	var t int
	for {
		<-tick.C
		t = time.Now().Second()
		fmt.Println(t)
		if t%10 == 0 {
			data_init()
			Tick(1e9) //需要保证在1s完成发送，发送完毕后，新加入的玩家需要等待下一次开局
			clear_tran()
			Tick(11e9)
			result_formate()
		}
	}
}

func main() {
	data_chan = make(chan *[]byte, 1)
	result = make(chan *[]byte, 1)
	listener()
}
