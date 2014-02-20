package main

import (
	. "./cell" //引用cell包，前边加上 . 在使用cell包变量时可以不写包名
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
	"unicode"
)

// 过滤单词，从in中读入，将英语单词写入一个通道中（通道异步读取）
func word_filter(in string, ch chan string) {
	ifile, _ := os.Open(in)     //打开in文件，返回文件和错误
	defer ifile.Close()         //函数返回或异常时关闭文件
	r := bufio.NewReader(ifile) //返回一个可以读文件的接口
	for {
		line, _, err := r.ReadLine() //逐行读取，返回读取的数据([]byte)，，错误
		if err != nil {
			break
		}
		s := string(line) //将一行数据转为字符串，判断长度
		l := len(s)
		if l < 3 || l > 16 {
			goto H
		}

		// 遍历该字符串，若不是字母，跳过；最后若是单词，将其写入管道
		for _, i := range s {
			if unicode.IsLetter(i) == false {
				goto H
			}
		}
		ch <- s
	H:
	}
}

// 从管道中读取字符串，main函数中将一部调用本函数
// 这里有个坑，返回的word（cell包中定义的数据结构 map[string]bool）为空，在循环体内不为空
// 所以直接json化会出问题，只能将其写入out文件中，再做后期处理
func read_chan(out string, ch chan string, quit chan bool) *Word {
	word := make(Word, 9900) //创建一个9900容量的Word
	ofile, _ := os.Create(out)
	writer := bufio.NewWriter(ofile)
	defer ofile.Close()
	var i int32 //用于计数，每N个字母一组，写入out文件中
	for {
		select {
		case s := <-ch:
			// fmt.Print("\t", s)
			i++
			writer.WriteString(s + " ")
			if i%15 == 0 {
				writer.WriteByte('\n')
				writer.Flush() //写入out文件
			}
			word[s] = true //将s存入word
		case <-quit: //当quit中有数据时，退出循环
			fmt.Println(word)

			break
		}

	}
	// b, err := json.Marshal(word)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// ioutil.WriteFile(out, b, 0644)
	// fmt.Println(string(b), word)
	// fmt.Println(word)
	// fmt.Println(string(b))
	return &word
}

// 读取in（readchan函数写出的.txt文件），json化后写出到out中(.json)
func word_to_json(in, out string) {
	ifile, _ := os.Open(in)
	defer ifile.Close()
	word := make(Word, 9999)
	reader := bufio.NewReader(ifile)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		ss := strings.Split(string(line), " ")
		for _, i := range ss {
			word[i] = true
		}
	}
	/*将word转换为json格式，将结果b写出到out*/
	b, _ := json.Marshal(word)
	fmt.Print(len(word), " words, ")
	ioutil.WriteFile(out, b, 0644)
}

func end_word_filter() {
	word_chan := make(chan string, 20) //存放word的通道
	quit := make(chan bool, 1)         //结束读取单词源，告知退出for循环
	word_source_file := "./source/word_source.txt"
	word_txt := "./source/word_txt.txt"
	word_json := "./source/word_json.json"
	t1 := time.Now()
	go func() {
		read_chan(word_txt, word_chan, quit)
	}()
	word_filter(word_source_file, word_chan)
	quit <- true //退出（通道）
	word_to_json(word_txt, word_json)
	t2 := time.Now()
	fmt.Println(t2.Sub(t1))
}
func main() {
	end_word_filter()
}
