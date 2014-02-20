/*


@date 2014/02/17
@author husd
@email hu.shengdong.h@gmail.com



游戏的客户端，提供了游戏的主界面


*/
package client

import(
	"fmt"
)

func main() {
	loadConfig()
	startClient()
}


/*

加载配置文件
*/
func loadConfig() {
	fmt.Println("loadConfig")
}

/*

开始游戏
*/
func startClient() {
	fmt.Println("game start ,welcome ")
	//检查环境，初始化页面
	//连接服务器
	//游戏开始
}

