package protocol

// 协议，protocol，定义了服务器与客户端交互的数据结构
import (
	"encoding/json"
	"fmt"
	"sort"
)

type Words map[string]int

type Resultwords struct { //玩家游戏结果数据
	Sequence string //字母序列
	Score    int    //得分
	Num      int    //单词个数
	ID       string //ID
	Words    *Words //找到的单词
}

// 单词的分数,???怎么做，再给出个计分策略
func (w *Words) score() int {
	return 10 * len(*w)
}

// 计算得分
func (r *Resultwords) Set_Score() {
	r.Score = 0
	for _, each := range *(r.Words) {
		r.Score += each
	}
}

// 序列化
func (r Resultwords) To_json() *([]byte) {
	b, e := json.Marshal(r)
	if e != nil {
		return nil
	}
	return &b
}

// 反序列化
func From_json(data *[]byte) (ok bool, r *Resultwords) {
	var res Resultwords
	e := json.Unmarshal(*data, &res)
	if e != nil {
		fmt.Println(e.Error())
		return false, nil
	}
	return true, &res
}

type Slice [](*Resultwords) //各个玩家的结果集

// 结果集的排序，继承自sort接口
func (s Slice) Len() int           { return len(s) }
func (s Slice) Less(i, j int) bool { return (*s[i]).Score > (*s[j]).Score }
func (s Slice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func (s *Slice) Sort() {
	sort.Sort(s)
}

// 结果集的序列化
func (s Slice) To_json() *([]byte) {
	b, e := json.Marshal(s)
	if e != nil {
		return nil
	}
	return &b
}

func (s *Slice) Show() {
	for i, each := range *s {
		fmt.Print(i)
		fmt.Println(((*each).Sequence))
	}
}
