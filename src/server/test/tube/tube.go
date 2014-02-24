package main

import (
	"../sequence"
	"bufio"
	"encoding/json"
	// "fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
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
	N   int
	vis Vis
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

func initAll() {
	N = 4
	vis = *initVis()

}

func once() {
	letters := []string{"abcdefg", "hijklmn", "opqrst", "uvwxyz"}
	src := sequence.Sequence(letters...)
	in := "./tube.txt"
	t := initTube(in, src)
	//fmt.Println(t)
	t.write("./tube.json")
	// fmt.Println("END init words.")
}

func initVis() *Vis {
	// var vis Vis
	return &vis
}

func (t Tube) dfs(e Enge, deep int, s string) {
	if vis[e.I][e.J] || deep > 10 {
		return
	}
	s = s + t[e.I][e.J].V
	// fmt.Println(s)
	vis[e.I][e.J] = true
	for _, item := range t[e.I][e.J].E {
		t.dfs(item, deep+1, s)
	}
	vis[e.I][e.J] = false
}
func main() {

	initAll()
	once()
	// ciku, _ = word.FromTxt("./word/filter/word.txt")

	// t := read("./tube/tube.json")
	// fmt.Println(t)
	// vis := initVis()
	// vis[0][0] = false
	// var e Enge
	// var s string
	// t.dfs(e, 0, s)
}
