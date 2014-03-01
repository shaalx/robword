package db

import (
	// "fmt"
	"testing"
)

func Test_Query(t *testing.T) {
	command := make([]interface{}, 3)
	command[0] = "SELECT * FROM wordmining.sequence;"
	command[1] = "SELECT * FROM wordmining.word_in_stuff"
	Query(command...)
	// fmt.Println(*m)
	// for i, stuff := range *m {
	// 	if stuff == nil {
	// 		continue
	// 	}
	// 	fmt.Println(i, stuff)
	// }
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
	command := make([]interface{}, 8)
	command[0] = "INSERT wordmining.word_in_stuff set word='mini',id=?"
	command[1] = "67"
	command[2] = "67"
	command[3] = "67"
	command[4] = "67"
	command[5] = "67"
	command[6] = "67"
	Prepare_excute(1, command[:]...) //insert data
	t.Log(" test Prepare_excute PASS.")
}

// go test -v
