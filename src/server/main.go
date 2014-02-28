package main

import (
	"./db"
	. "./protocal"
	"fmt"
)

func example() {
	var r Resultwords
	r.Sequence = "xxoo"
	r.Words = &map[string]bool{"wo": true, "ai": true, "ni": true}
	db.Inserting(r)
}

func example2() {
	r := db.Selecting()
	fmt.Println(r.Sequence, *r.Words)
}
func main() {
	example2()
}
