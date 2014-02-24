package word

import (
	// "bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Ciku map[string]bool

func FromTxt(in string) (*Ciku, int) {
	ifile, _ := os.Open(in)
	defer ifile.Close()
	b, err := ioutil.ReadAll(ifile)
	if err != nil {
		return nil, 0
	}
	s := string(b)
	ss := strings.Split(s, " ")
	ciku := make(Ciku, 5000)
	// fmt.Println(len(ss))
	l := len(ss)
	for _, item := range ss {
		// if i >= 1959 {
		// 	break
		// }
		// ciku[i].M = make(map[string]bool, 1)
		ciku[item] = true
	}

	//fmt.Println(ciku)
	return &ciku, l
}

func test() {
	ciku, l := FromTxt("./filter/word.txt")
	out := "./filter/ciku.json"
	c := (*ciku)
	fmt.Println(c, *ciku, l)
	bu, _ := json.Marshal(c)
	// json.Unmarshal(bu, &ciku)
	ofile, _ := os.Create(out)
	defer ofile.Close()
	ioutil.WriteFile(out, bu, 0644)
	fmt.Println("end")
}

func FromJson(in string) (*Ciku, int) {
	// in := "./filter/ciku.json"
	b, _ := ioutil.ReadFile(in)
	var ciku Ciku
	json.Unmarshal(b, &ciku)
	return &ciku, len(ciku)
}

// func main() {
// 	test()
// 	// testing()
// }
