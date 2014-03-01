package db

import (
	. "../protocal"
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

// 测试随机选取模块，返回一个字母序列，以及其包含的单词
func Test_Selecting(t *testing.T) {
	m := Selecting()
	fmt.Println(*m)
	for str, _ := range *(*m).Words {
		fmt.Print(str, "\t")
	}
	t.Log(" test Selecting PASS.")
}

// 插入测试不好测试，需要运行集成模块来测试，经过模拟测试（随机生成一个字符串），该模块是正确的。
func Test_Inserting(t *testing.T) {
	var r Resultwords
	r.Sequence = rand_string()
	r.Words = &map[string]bool{"we": true, "are": true, "the": true, "world!": true}
	Inserting(r)
	t.Log(" test Inserting PASS.")
}

//
func rand_string() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	i := r.Int()
	return strconv.Itoa(i)
}
