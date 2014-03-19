package protocal

import (
	"encoding/json"
)

type Words map[string]int

type Resultwords struct {
	Sequence string //'json:"sequence"'
	Words    *Words //'json:"words"'
}

func (r Resultwords) To_json() *([]byte) {
	b, e := json.Marshal(r)
	if e != nil {
		return nil
	}
	return &b
}
