package main

import (
	. "./db"
	. "./protocol"
	. "./tick"
	// "bytes"
	"fmt"
	"net"
	"os"
	"time"
)

var (
	data_chan        chan *([]byte) //游戏原始数据通道
	result_chan      chan *([]byte) //游戏排名结果数据通道
	palyer_data_chan chan *([]byte) //玩家游戏结果数据通道
	player_num       int64          //玩家个数，暂时没有用到，留着提高性能：见下面解释
	player_data      Slice          //所有玩家结果数据，暂时用到的是动态数组，效率不高，若要求高性能，需要做改进。
)

//服务器侦听，同时启动一个gorutine,实现整体的时间同步（在特定的时刻，将数据放入通道中，待取）。
func listener() {
	listener, e := net.Listen("tcp", ":1234")
	system_error(e)
	go sync() //时间同步
	for {
		conn, er := listener.Accept()
		system_error(er)
		go data_exchange(&conn) //数据交换
	}
}

func checkerr(e error) {
	if e != nil {
		fmt.Println(e.Error())
	}
}

// 系统错误，即服务器的系统错误，而非某个客户端的连接引起的错误
func system_error(e error) {
	if e != nil {
		fmt.Print(e.Error())
		os.Exit(-1)
	}
}

/*数据交换
连接了客户端，读取玩家基本信息；从data_chan中读取游戏数据，并发送；读取玩家游戏结果数据；发送游戏排名数据*/
func data_exchange(conn *net.Conn) {
	defer (*conn).Close()
	for {
		read_info(conn)
		// t1 := time.Now()
		if !write_data(conn) {
			return
		}
		// t2 := time.Now()
		// fmt.Println(t2.Sub(t1))
		player_num++
		fmt.Println("-------------------------------------------------")
		fmt.Println("1分钟后收集所有玩家游戏结果.")
		read_result(conn)
		statistics()
		if !write_reslut(conn) {
			return
		}
	}
}

/*读取游戏原始数据，并发送。
读取玩之后，还需把该信息放回，供其他玩家读取*/
func write_data(conn *net.Conn) (connected bool) {
	b := <-data_chan
	data_chan <- b
	_, e := (*conn).Write(*b)
	if e != nil {
		fmt.Println(e)
		return false
	}
	fmt.Println("-------------------------------------------------")
	fmt.Println("数据已经发送至玩家")
	return true
}

/*读取游戏排名数据，并发送，其他同write_data()*/
func write_reslut(conn *net.Conn) (connected bool) {
	b := <-result_chan
	result_chan <- b
	_, e := (*conn).Write(*b)
	if e != nil {
		fmt.Println(e)
		return false
	}
	// fmt.Println(string(*b))
	fmt.Println("-------------------------------------------------")
	fmt.Println("排名结果已经发送至玩家")
	return true
}

/*读取玩家基本信息(感觉这个模块没有用，服务器不必维护一个玩家列表，只需要维护一个统计玩家结果数据的数组，每一组需要有玩家的信息，暂时这么考虑)*/
func read_info(conn *net.Conn) (connected bool) {
	fmt.Println("-------------------------------------------------")
	fmt.Println("获取玩家基本信息", (*conn).RemoteAddr())
	return true
}

/*读取玩家游戏结果数据，读取的数据放入player_data_chan*/
func read_result(conn *net.Conn) (connected bool) {
	b := make([]byte, 1024)
	n, er := (*conn).Read(b)
	if er != nil {
		fmt.Println(er)
		return false
	}
	fmt.Println("-------------------------------------------------")
	fmt.Println(n, "bytes，数据来自玩家：", (*conn).RemoteAddr())
	fmt.Println(string(b[:n]))
	bb := b[:n]
	palyer_data_chan <- &bb
	return true
}

/*统计玩家游戏结果数据（从player_data_chan中读取，反序列化后存入player_data数组(等待排序)）*/
func statistics() {
	data := <-palyer_data_chan
	ok, r := From_json(data)
	if !ok {
		return
	}
	r.Set_Score()
	// fmt.Println("======from players=========")
	// fmt.Println(r)
	if player_data == nil {
		player_data = make(Slice, 1)
		player_data[0] = r
	} else {
		player_data = append(player_data, r)
	}
	// fmt.Println(player_data)
}

/*游戏初始数据的初始化，从数据库中选择一条记录，序列化，存入data_chan*/
func data_init() {
	// a := "游戏数据"
	// bs := bytes.NewBufferString(a)
	// b := bs.Bytes()
	clear_data()
	// data_chan <- b
	var r *Resultwords
	r = Selecting()
	bys := r.To_json()
	data_chan <- bys
}

/*游戏排名结果序列化，结果放入result_chan*/
func result_formate() {
	// a := "游戏结果排名数据：*****"
	// bs := bytes.NewBufferString(a)
	// b := bs.Bytes()

	// statistics()
	// player_data.Show()

	player_data.Sort()
	bys := player_data.To_json()
	// fmt.Println(string(*bys))
	// fmt.Println("......result@ server(will sent to client)......")
	// player_data.Show()
	result_chan <- bys
}

/*清空通道*/
func clear_player_data() {
	player_num = 0
	player_data = nil
}

//清空通道//
func clear_data() {
	if len(data_chan) == 1 {
		<-data_chan
	}
}

/*清空通道*/
func clear_result() {
	if len(result_chan) == 1 {
		<-result_chan
	}
}

/*时间的同步
等待若干秒（玩家加入游戏），
 [同步]发送游戏原始数据，[同步]发送游戏排名结果数据；*/
func sync() {
	tick := time.NewTicker(2e9)
	for {
		<-tick.C
		// <-tick.C
		// <-tick.C
		// var t int
		// t = time.Now().Second()
		// fmt.Println(t)
		// if t%10 == 0 {
		clear_player_data()
		data_init()
		Tick(1e6) //需要保证在1s完成发送，发送完毕后，新加入的玩家需要等待下一次开局
		// 其含义：允许客户端的时间延迟。1s后若还不能加入游戏的玩家，只能等待下一次开局
		// 同时，也是为了保证加入的玩家都能接收到游戏数据（需在1s内完成接收）
		clear_data()
		clear_result()
		Tick(10e9)
		result_formate()
		// }
	}
}

func init() {
	data_chan = make(chan *[]byte, 1)
	result_chan = make(chan *[]byte, 1)
	palyer_data_chan = make(chan *[]byte, 100)
}

/**/
func main() {

	listener()
}
