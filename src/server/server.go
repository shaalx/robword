/*


@date 2014/02/17
@author husd
@email hu.shengdong.h@gmail.com



游戏的服务端，提供了游戏的后台运算功能入口


*/

package main

import(
	"fmt"
	"./question"
)

/*

加载配置文件
*/
func loadConfig() {
	fmt.Println("loadConfig")
}

func main() {
	startServer()
}
/*

开始游戏
*/
func startServer() {
	fmt.Println("gameserver start")
	loadConfig()
	//TODO 初始化系统环境，加载数据
	//
	//监听来自客户端的连接
	//从数据库中随机取一个
	testQues()
}

func testQues() {
	m := make(map[string]int) //使用make创建一个空的map      
	m["one"] = 1    
	m["two"] = 2     
	m["three"] = 3 

	//需要从sequence中的函数返回一个结构，
	//实际运行的时候是客户端从服务端得到一个json，
	//解析这个json并转换为Ques结构
	//这里我自己简单初始化了一个

	qu := question.Ques{"hellsjielxmdisen",m}
	qu.Display()
	fmt.Printf("is word legal:%v\n",qu.IswordLegal("one"))
	fmt.Printf("word one's goal is:%v\n",qu.GetScore("two"))
	fmt.Printf("right words count is:%v they are :\n",qu.GetWordsCount())
	qu.DisplayWords()
}
