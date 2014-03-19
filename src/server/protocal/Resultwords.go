package protocal

import (
	"encoding/json"
	"sort"
)

type Words map[string]int

type Resultwords struct {
	Sequence string //'json:"sequence"
	Score    int
	Words    *Words //'json:"words"'
}

func (r Resultwords) To_json() *([]byte) {
	b, e := json.Marshal(r)
	if e != nil {
		return nil
	}
	return &b
}

type Slice []Resultwords

func (s Slice) Len() int           { return len(s) }
func (s Slice) Less(i, j int) bool { return s[i].Score > s[j].Score }
func (s Slice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func (s *Slice) Sort() {
	sort.Sort(s)
}
