package db

import (
	. "../protocal"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func fmt_p() {
	fmt.Println("...")
}

func Selecting() *Resultwords {
	command := make([]interface{}, 10)
	command[0] = "select * from wordmining.sequence"
	m := Query(command...)
	length := len((*m)[0])
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	i := r.Intn(length)
	// fmt.Println(*m)
	// fmt.Println(i)
	var result Resultwords
	j := 0
	var id_target int
	for stuff, id := range (*m)[0] {
		// fmt.Println(stuff, id)
		if j == i {
			result.Sequence = stuff
			id_target = id
			break
		}
		j++
	}
	command[0] = "select * from wordmining.word_in_stuff where id=" + strconv.Itoa(id_target)
	m = Query(command...)
	// fmt.Println(command)
	w := (make(Words, 100))
	result.Words = &w
	for stuff, _ := range (*m)[0] {
		(*result.Words)[stuff] = 30
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
