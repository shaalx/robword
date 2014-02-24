/*
生成单词序列

方法：从一个字符序列中随机取得16次字母，组合成一个序列
*/
package sequence

// package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	// "unicode"
)

//从一个单词集合中得到一个字母集合
func get_letter_set(set ...string) string {
	var result string

	// 如果有非字母字符，需要执行以下语句
	// for _, sub_string := range set {
	// 	str_spilt := strings.Split(sub_string, " ")
	// 	str_join := strings.Join(str_spilt, "")
	// 	result = strings.Join([]string{result, str_join}, "")
	// }

	// 单词集合中每一项不能有空格
	result = strings.Join(set, "")
	return result
}

// 得到一个随机数数组，每一项对应字母序列的索引
func rand_nums(len int) *([]int) {
	nums := make([]int, 16)
	rander := rand.New(rand.NewSource(time.Now().UnixNano()))
	perm_nums := rander.Perm(len)
	if len < 16 {
		// nums = perm_nums[:]
		for i := 0; i < len; i++ {
			nums[i] = perm_nums[i]
		}
		for i := len; i < 16; i++ {
			nums[i] = rander.Intn(len)
		}
	} else {
		nums = perm_nums[:16]
	}
	// fmt.Println(nums)
	return &nums
}

// 字母随机序列
// 从letter_set中抽取16个字母组成一个序列
func letter_sequence(letter_set string) string {
	set_len := len(letter_set)
	nums := rand_nums(set_len)
	// fmt.Println(*nums)
	str := make([]string, 16)
	for i, n := range *nums {
		str[i] = string(letter_set[n])
	}
	result := strings.Join(str, "")
	return result
}

// 封装，对外提供借口
func Sequence(wordset ...string) string {
	s := get_letter_set(wordset...)
	sequence := letter_sequence(s)
	// fmt.Println("................")
	//fmt.Println(sequence)
	for i:=0; i<16; i++ {
		fmt.Print(" %v ",sequence[i])
		if(i%4 == 0){
			fmt.Println()
		}
	// fmt.Println("END generate sequence.")
	return sequence
}
func test() {
	wordset := []string{"worsd", "ssecond"}
	s := get_letter_set(wordset[0:]...)
	squence := letter_sequence(s)
	fmt.Println("................")
	fmt.Println(squence)

}

// func main() {
// 	// test()
// 	wordset := []string{"worsd", "ssecond"}
// 	Sequence(wordset...)
// }
