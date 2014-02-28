package db

import (
	"fmt"
	"sync"
	"testing"
)

// Database's connection nums is limitted,so use the Once to run the inits(return a connection to DB) only once.
func init() {
	var once sync.Once
	once.Do(Inits())
}

func Test_Query(t *testing.T) {
	command := make([]interface{}, 3)
	command[0] = "SELECT * FROM wordmining.sequence;"
	command[1] = "SELECT * FROM wordmining.word_in_stuff"
	m := Query(command...)
	// fmt.Println(*m)
	for i, stuff := range *m {
		if stuff == nil {
			continue
		}
		fmt.Println(i, stuff)
	}
	t.Log(" test Query PASS.")
}

func Test_Excute(t *testing.T) {
	command := make([]interface{}, 3)
	command[0] = "INSERT into wordmining.word_in_stuff values('Query_test1',8),('Query_test2',5)"
	command[1] = "INSERT into wordmining.word_in_stuff values('-----',7),('++++++',6)"
	Excute(command...)
	t.Log(" test Excute PASS.")
}

func Test_Prepare_excute(t *testing.T) {
	command := make([]interface{}, 5)
	command[0] = "INSERT wordmining.word_in_stuff set word=?,id=?"
	command[1] = "Prepare_moon"
	command[2] = "4"
	command[3] = "Prepare_like"
	command[4] = "4"
	Prepare_excute(2, command[:]...) //insert data
	t.Log(" test Prepare_excute PASS.")
}

// go test -v
