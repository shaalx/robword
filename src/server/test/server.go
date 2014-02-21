package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Result map[string]bool

func resultFromJson(in string) *Result {
	b, _ := ioutil.ReadFile(in)
	var r Result
	json.Unmarshal(b, &r)
	return &r
}

func main() {
	r := resultFromJson("result.json")
	fmt.Println(len(*r))
	fmt.Println("................")
	for item := range *r {
		fmt.Printf("%16s", item)
	}
	fmt.Println()
}
