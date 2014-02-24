/*

@date 2014/02/123
@author husd
@email hu.shengdong.h@gmail.com

包含了问题和答案,以及对它的各种处理方法
数据的可见性都设置了可见，安全性不好，再考虑怎么改
所有的函数都没有添加错误处理 ---测试中补上

*/
package question

import(
	"fmt"
)

type Ques struct{
	Seq string          //16个字母
	Words map[string]int  //答案 [单词]分值
}

//显示问题和答案
func ( q * Ques) Display() {
	for i:=0; i<16;i++ {
		fmt.Printf(" %c",q.Seq[i])
		if i>0 && (i+1)%4 == 0 {
			fmt.Println()
		}
	}
	fmt.Println(q.Words)
}

//判断玩家生成的单词是否合法
func (qu * Ques) IswordLegal(word string)bool {
	return qu.Words[word] > 0
}

//得到某个单词的分数
func (qu * Ques) GetScore(word string)int{
	return qu.Words[word]
}

//显示所有的正确单词
func (qu * Ques) DisplayWords() {
	for word,_ := range qu.Words {
		fmt.Printf(" %v \n",word)
	}
}

//显示正确单词总数
func (qu * Ques) GetWordsCount()int{
	return len(qu.Words)
}

func m() {
	//qu := QUES{"hello",{"hel":1}}
	m := make(map[string]int) //使用make创建一个空的map      
	 m["one"] = 1    
	 m["two"] = 2     
	 m["three"] = 3 
	qu := Ques{"hellsjielxmdisen",m}
	qu.Display()
}
