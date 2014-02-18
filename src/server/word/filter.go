package main

import (
	. "./cell"
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
	"unicode"
)

// type Word map[string]bool

func word_filter(in string, ch chan string) {
	ifile, _ := os.Open(in)
	defer ifile.Close()
	r := bufio.NewReader(ifile)
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			break
		}
		s := string(line)
		l := len(s)
		if l < 3 || l > 16 {
			goto H
		}
		for _, i := range s {
			if unicode.IsLetter(i) == false {
				goto H
			}
		}
		ch <- s
	H:
	}
}

func read_chan(out string, ch chan string, quit chan bool) *Word {
	word := make(Word, 9900)

	ofile, _ := os.Create(out)
	writer := bufio.NewWriter(ofile)
	defer ofile.Close()
	var i int32
	for {
		select {
		case s := <-ch:
			// fmt.Print("\t", s)
			i++
			writer.WriteString(s + " ")
			if i%15 == 0 {
				writer.WriteByte('\n')
				writer.Flush()
			}
			word[s] = true
		case <-quit:
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
	b, _ := json.Marshal(word)
	fmt.Println(len(word))
	ioutil.WriteFile(out, b, 0644)
}

func end_word_filter() {
	ch := make(chan string, 1)
	quit := make(chan bool, 1)
	word_source_file := "./source/English_word_source.txt"
	word_txt := "./source/word_txt.txt"
	word_json := "./source/word_json.json"
	t1 := time.Now()
	go func() {
		read_chan(word_txt, ch, quit)
	}()
	word_filter(word_source_file, ch)
	quit <- true
	word_to_json(word_txt, word_json)
	t2 := time.Now()
	fmt.Println(t2.Sub(t1))
}
func main() {
	end_word_filter()
}
