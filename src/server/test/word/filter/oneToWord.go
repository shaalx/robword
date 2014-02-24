package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"unicode"
)

func Read(in string, ch chan string) bool {
	file, _ := os.Open(in)
	defer file.Close()
	r := bufio.NewReader(file)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			// fmt.Println(err.Error())
			break
		}
		GetWord(b, ch)
	}
	return true
}

func GetWord(b []byte, ch chan string) {
	s := string(b)
	// rp := strings.NewReplacer("v.", " ")
	// s = rp.Replace(s)
	ss := strings.Split(s, " ")
	for _, it := range ss {
		l := len(it)
		if l < 3 || l > 15 {
			continue
		}
		// for i := 0; i < l; i++ {
		for _, sub := range it {
			if !unicode.IsLetter(sub) {
				goto H
			}
		}
		// i := l / 2
		// j := it[i]
		// if unicode.IsLetter(j) {
		ch <- it

		// }
		// 	unicode.IsLetter(r)
		// 	if j <= 122 && j >= 97 || j <= 90 && j >= 65 { //65 90 97 122
		// 		ch <- it
		// 		// fmt.Println(it)
		// 	}
		// }
	H:
	}
}

func Write(out string, ch chan string, qch chan bool) {
	file, err := os.Create(out)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()
	w := bufio.NewWriter(file)
	for {
		//s := ""
		for it := range ch {
			//s = strings.Join([]string{s, it}, "\t")
			w.WriteString(it + " ")
			w.Flush()
		}
		// l := len(ch)
		// for i := 0; i < l; i++ {
		// 	ch_str := <-ch
		// 	s = strings.Join([]string{s, ch_str}, "\t")
		// }
		// w.WriteString(s)
		w.Flush()
		select {
		case <-qch:
			return
		}
	}
}
func Filter(dst, src string) {

	ch := make(chan string, 50)
	qch := make(chan bool, 1)

	t1 := time.Now()

	go Write(dst, ch, qch)
	Read(src, ch)

	qch <- true
	time.Sleep(1)

	t2 := time.Now()

	fmt.Print("END filter to wordTxt /filter")
	fmt.Println((float32)(t2.UnixNano()-t1.UnixNano())/(1e9), "s")
}

func main() {
	Filter("./word_crash.txt", "./1.txt")
}
