package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Graph [4][4]Bee
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

func initGraph(in, src string) *Graph {
	var t Graph
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

func (t Graph) write(out string) {
	file, _ := os.Create(out)
	defer file.Close()
	b, _ := json.Marshal(t)
	ioutil.WriteFile(out, b, 0644)
}

func read(in string) *Graph {
	b, _ := ioutil.ReadFile(in)

	var v Graph
	json.Unmarshal(b, &v)
	return &v
}

func initAll() {
	N = 4
	vis = *initVis()

}

func once() {
	src := "hhelacabnrotfiti"
	in := "./graph.txt"
	t := initGraph(in, src)
	//fmt.Println(t)
	t.write("./graph.json")
	fmt.Println("END txtTube to jsonTube /Graph")
}

func initVis() *Vis {
	// var vis Vis
	return &vis
}

func (t Graph) dfs(e Enge, deep int, s string) {
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

	// t := read("./graph/graph.json")
	// fmt.Println(t)
	// vis := initVis()
	// vis[0][0] = false
	// var e Enge
	// var s string
	// t.dfs(e, 0, s)
}
