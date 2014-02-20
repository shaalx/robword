/*


@date 2014/02/17
@author husd
@email hu.shengdong.h@gmail.com



游戏的服务端，提供了游戏的后台运算功能入口


*/

package server

import(
	"fmt"
)

main(){
	loadConfig()
	startServer()
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
func startServer() {
	fmt.Println("gameserver start")
	//TODO 初始化系统环境，加载数据
	//
	//监听来自客户端的连接


}
