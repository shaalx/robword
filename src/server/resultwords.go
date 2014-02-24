package main

import (
	"./sequence"
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

// 无向图，4*4=16个节点。 即在这个序列图上生成字母组合
type Graph [4][4]Neighbour

// 对应无向图节点是否访问过
type Vis [4][4]bool

// 无向图节点，
// V：该节点的值（字母）
// E：邻接点（相邻的节点）
type Neighbour struct {
	V string
	E []Edge
}

// 邻接点信息，I行 J列
type Edge struct {
	I, J int
}

// 单词词库
type Ciku map[string]bool

type Resultwords struct {
	Sequence string //'json:"sequence"'
	Nun      int
	Words    *map[string]bool //'json:"words"'
}

// 全局变量
var (
	N           int         //N行N列（N=4）
	vis         Vis         // 记录是否访问过
	ciku        *Ciku       //单词库，预加载
	result      chan string // 生成单词后放入该通道
	resultwords Resultwords //结果
)

// 初始化无向图，in（.txt）文件中读入无向图的初始化信息，以src字母序列初始化
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
			(t[i][j]).E = make([]Edge, len(ss)-1)
			k := 0
			for _, it := range ss {
				item, err := strconv.Atoi(it)
				if err != nil {
					break
				}
				if item >= N*N {
					break
				}
				(t[i][j]).E[k] = Edge{I: item / N, J: item % N}
				k++
			}
		}
	}
	return &t
}

// 将无向图写出到out(.json) 中
func (t Graph) write(out string) {
	file, _ := os.Create(out)
	defer file.Close()
	b, _ := json.Marshal(t)
	ioutil.WriteFile(out, b, 0644)
}

// 从in(.json)读取无向图
func read(in string) *Graph {
	b, _ := ioutil.ReadFile(in)
	var v Graph
	json.Unmarshal(b, &v)
	return &v
}

// 从in(.json)初始化无向图：
func initGraphFromJson(in string) *Graph {
	t := read(in)
	return t
}

// 使用字母序列初始化无向图
func (t *Graph) inti_graph_with_sequence(sequence string) {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			index := i*N + j
			t[i][j].V = string(sequence[index])
		}
	}
}

// 词库
func FromJson(in string) (*Ciku, int) {
	b, _ := ioutil.ReadFile(in)
	var ciku Ciku
	json.Unmarshal(b, &ciku)
	return &ciku, len(ciku)
}

// 初始化全局变量
func initAll() {
	N = 4
	vis = *initVis()
	result = make(chan string, 1)
	ciku, _ = FromJson("./word/source/word_json.json")
}

// 初始化访问标志   （其实没有必要）
func initVis() *Vis {
	return &vis
}

// 该函数只调用一次，用于(从文本文件)初始化无向图，得到无向图的json结构，并写出到json
func once() {
	src := "abidfrtegshudrew"
	in := "./graph/graph.txt"
	t := initGraph(in, src)
	fmt.Println(t)
	t.write("./graph/graph.json")
}

// 核心代码，深度优先算法遍历该无向图，若单词合法，将其压入单词通道result中
// e：当前访问的位置
// deep:访问的深度，调节deep的值可以改变整体性能
// s：当前访问得到的字母组合
func (t Graph) dfs(e Edge, deep int, s string) {
	if vis[e.I][e.J] || deep > 15 {
		return
	}
	s = s + t[e.I][e.J].V
	if (*ciku)[s] {
		result <- s
	}
	vis[e.I][e.J] = true
	for _, item := range t[e.I][e.J].E {
		t.dfs(item, deep+1, s)
	}
	vis[e.I][e.J] = false
}

// 读取result单词通道，最后返回得到的单词
// q ：结束命令，若q中有数据，表示可以return了
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
	resultwords.Words = &m
	resultwords.Nun = len(m)
	if resultwords.Nun > 35 {
		resultwords.resultWords_to_json()
	}
	return &m
}

// 将结果out(map[string]bool)输出到out(.json)
func resultToJson(out string, m *map[string]bool) {
	of, _ := os.Create(out)
	defer of.Close()
	b, _ := json.Marshal(m)
	of.Write(b)
}

// 读取结果：测试用（可以删除）
func resultFromJson(in string) {
	var m map[string]bool
	b, _ := ioutil.ReadFile(in)
	json.Unmarshal(b, &m)
	for item := range m {
		fmt.Println(item)
	}
}

// 结果输出
func (r *Resultwords) resultWords_to_json() {
	of, _ := os.OpenFile("resutlwords.json", os.O_CREATE|os.O_APPEND, 0644)
	defer of.Close()
	b, _ := json.Marshal(resultwords)
	of.Write(b)
}

// 开始
func start() {
	graph := initGraphFromJson("./graph/graph.json")
	for {
		sequence := sequence.Sequence_2()
		resultwords.Sequence = sequence
		graph.inti_graph_with_sequence(sequence)
		one_finding(graph)
		var order string
		fmt.Scanf("%s", &order)
	}
}

// 做一次序列的查找
func one_finding(graph *Graph) {
	var e Edge
	var s string
	qb := make(chan bool, 1)
	t1 := time.Now()
	var m map[string]bool
	go func(m *map[string]bool) {
		m = readChan(qb)
		fmt.Printf("\nEND, %d word found , ", len(*m))
	}(&m) // go readChan(qb)
	fmt.Println("start...")
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			e.I = i
			e.J = j
			s = ""
			graph.dfs(e, 0, s)
		}
	}
	time.Sleep(1)
	qb <- true
	time.Sleep(1)
	t2 := time.Now()
	fmt.Println(t2.Sub(t1))
}
func main() {
	initAll()
	// once()
	start()
}
