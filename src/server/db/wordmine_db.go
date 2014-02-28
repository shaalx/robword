package db

import (
	. "../protocal"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func init() {
	var once sync.Once
	once.Do(Inits())
}

func fmt_p() {
	fmt.Println("...")
}

// func main() {
// 	Select()
// }

func Selecting() *Resultwords {
	command := make([]interface{}, 10)
	command[0] = "select * from wordmining.sequence"
	m := Query(command...)
	length := len((*m)[0])
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	i := r.Intn(length) + 1
	// fmt.Println(*m)
	// fmt.Println(i)
	var result Resultwords
	for stuff, id := range (*m)[0] {
		// fmt.Println(stuff, id)
		if id == i {
			result.Sequence = stuff
		}
	}
	command[0] = "select * from wordmining.word_in_stuff where id=" + strconv.Itoa(i)
	m = Query(command...)
	w := (make(map[string]bool, 100))
	result.Words = &w
	for stuff, _ := range (*m)[0] {
		(*result.Words)[stuff] = true
	}
	// fmt.Println(result, *result.Words)
	return &result
}

func Inserting(r Resultwords) {
	command := make([]interface{}, 100)
	command[0] = "INSERT wordmining.sequence SET stuff=?,id=''"
	command[1] = r.Sequence
	Prepare_excute(1, command...)

	command[0] = "select * from wordmining.sequence where stuff='" + r.Sequence + "'"
	// fmt.Println(command[0:1])
	m := Query(command[0:1]...)

	command[0] = "INSERT wordmining.word_in_stuff SET word=?,id=" + strconv.Itoa((*m)[0][r.Sequence])
	i := 1
	for stf, _ := range *r.Words {
		// fmt.Println(stf)
		command[i] = stf
		i++
	}
	// fmt.Println(command)
	Prepare_excute(1, command...)
}
