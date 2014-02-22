package player

import(
	"fmt"
)

type Player struct{
	name string   //玩家名称
	email string  
	score int     //得分
	isR bool      //是否是游客
	seq int       //排名
	words         //拼出的单词,怎么存储?
}