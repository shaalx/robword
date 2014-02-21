package main

import (
	"./word"
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

type Tube [4][4]Bee
type Vis [4][4]bool
type Bee struct {
	V string
	E []Enge
}

type Enge struct {
	I, J int
}

var (
	N      int
	vis    Vis
	ciku   *word.Ciku
	result chan string
)

func initTube(in, src string) *Tube {
	var t Tube
	file, e := os.Open(in)
	if e != nil {
		return nil
	}
	defer file.Close()

	r := bufio.NewReader(file)
	for i := 0; i < N; i++ {

		for j := 0; j < N; j++ {
			(t[i][j]).V = string(src[i*N+j])
			b, _, er := r.ReadLine()
			if er != nil {
				break
			}
			s := string(b)
			ss := strings.Split(s, " ")
			(t[i][j]).E = make([]Enge, len(ss)-1)
			k := 0
			for _, it := range ss {
				item, err := strconv.Atoi(it)
				if err != nil {
					break
				}
				if item >= N*N {
					break
				}
				(t[i][j]).E[k] = Enge{I: item / N, J: item % N}
				k++
			}
		}
	}
	return &t
}

func (t Tube) write(out string) {
	file, _ := os.Create(out)
	defer file.Close()
	b, _ := json.Marshal(t)
	ioutil.WriteFile(out, b, 0644)
}

func read(in string) *Tube {
	b, _ := ioutil.ReadFile(in)

	var v Tube
	json.Unmarshal(b, &v)
	return &v
}

func initTubeFromJson(in string /*, src string*/) *Tube {
	t := read(in)
	// 	var src string
	// 	fmt.Println("................")
	// 	fmt.Scanf("%s", &src)
	// 	l := len(src)
	// 	for i := 0; i < N; i++ {
	// 		for j := 0; j < N; j++ {
	// 			index := i*N + j
	// 			if index >= l {
	// 				goto H
	// 			}
	// 			t[i][j].V = string(src[index])
	// 		}
	// 	}
	// 	for i := 0; i < N; i++ {
	// 		for j := 0; j < N; j++ {
	// 			index := i*N + j
	// 			if index >= l {
	// 				goto H
	// 			}
	// 			t[i][j].V = string(src[index])
	// 		}
	// 	}
	// H:
	return t
}

func initAll() {
	N = 4
	vis = *initVis()
	result = make(chan string, 1)
}

func once() {
	src := "abidfrtegshudrew"
	in := "./tube/tube.txt"
	t := initTube(in, src)
	fmt.Println(t)
	t.write("./tube/tube.json")
}

func initVis() *Vis {
	// var vis Vis
	return &vis
}

func (t Tube) dfs(e Enge, deep int, s string) {
	if vis[e.I][e.J] || deep > 15 {
		return
	}
	s = s + t[e.I][e.J].V
	if (*ciku)[s] {
		// os.Stdout.WriteString(s + "\n")
		result <- s
	}
	vis[e.I][e.J] = true
	for _, item := range t[e.I][e.J].E {
		t.dfs(item, deep+1, s)
	}
	vis[e.I][e.J] = false
}

func readChan(q chan bool) *(map[string]bool) {
	m := make(map[string]bool, 500)
	quit := false
	for {
		select {
		case <-q:
			quit = true
		case r := <-result:
			fmt.Printf("%16s", r)
			m[r] = true
		}
		if quit {
			break
		}
	}
	// for item := range result {
	// 	// fmt.Println(item, "---o")
	// 	m[item] = true
	// 	// fmt.Println(m)
	// }
	// fmt.Println(m)
	// fmt.Println("END...")
	return &m
}

func resultToJson(out string, m *map[string]bool) {
	of, _ := os.Create(out)
	defer of.Close()
	b, _ := json.Marshal(m)
	ioutil.WriteFile(out, b, 0644)
}

func resultFromJson(in string) {
	var m map[string]bool
	b, _ := ioutil.ReadFile(in)
	json.Unmarshal(b, &m)
	for item := range m {
		fmt.Println(item)
	}
}
func start() {
	// ciku, _ = word.FromTxt("./word/filter/word.txt")
	ciku, _ = word.FromJson("./word/filter/ciku.json")
	// fmt.Println(len(*ciku))
	// t := read("./tube/tube.json")
	t := initTubeFromJson("./tube/tube.json")
	var e Enge
	var s string
	qb := make(chan bool, 1)
	t1 := time.Now()
	var m map[string]bool
	go func(ma *map[string]bool) {
		ma = readChan(qb)
		// fmt.Println(*ma)
		resultToJson("result.json", ma)
	}(&m)
	// go readChan(qb)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			e.I = i
			e.J = j
			s = ""
			t.dfs(e, 0, s)
		}
	}
	time.Sleep(1)
	qb <- true
	time.Sleep(1)
	// resultFromJson("result.json")
	t2 := time.Now()
	fmt.Println("\nEND wordFound,", t2.Sub(t1))
}
func main() {
	initAll()
	// once()
	start()
}
