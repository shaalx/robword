/*
根据字母在单词中出现的频率，随机生成字母序列
*/
package percent

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

// 将频率表放在percentage.json中,Percent 为数据结构
type Percent map[string]float64

// 从.json中读取频率表，返回Percent的一个实例指针
func fromJson() *Percent {
	in := "percentage.json"
	var per Percent
	data, _ := ioutil.ReadFile(in)
	json.Unmarshal(data, &per)
	return &per
}

// 生成字符序列集合,如：
// e的概率为0.12**，则e对应的项为 Ceil(100*0.12**)=13 个e
func Generate_set() *([]string) {
	per := fromJson()
	result := make([]string, 26)
	var index int
	for letter, p := range *per {
		i := (int)(math.Ceil(p * 100))
		result[index] = strings.Repeat(letter, i)
		index++
	}
	// fmt.Println(result)
	return &result
}

func example() {
	fmt.Println(Generate_set())
}

// func main() {
// 	example()
// }
