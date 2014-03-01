package db

import (
	"testing"
)

// Benchmarking
func Benchmark_Excute(b *testing.B) {
	b.StopTimer()
	command := make([]interface{}, 3)
	command[0] = "INSERT into wordmining.word_in_stuff values('insert',5),('dowo',6)"
	b.StartTimer()
	// for i := 0; i < b.N; i++ {
	// 	Query(command[0])
	// }
	/* ERROR 1040：too many connections
	http://www.2cto.com/database/201112/112738.html*/
	Excute(command[0])
	b.Log(" Excute PASS.")
}

func Benchmark_Prepare_excute(b *testing.B) {
	b.StopTimer()
	command := make([]interface{}, 6)
	command[0] = "update wordmining.word_in_stuff set word=? where id=?"
	command[1] = "update"
	command[2] = "3"
	command[3] = "delete from wordmining.word_in_stuff where id>? and id<?"
	command[4] = "1"
	command[5] = "2"
	b.StartTimer()
	Prepare_excute(2, command[3:6]...) //update data
	// Prepare_excute(command[3:5]...) //delete data
	b.Log(" Prepare_excute PASS.")
}

// go test -v -bench=".*"
