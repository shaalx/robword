package protocol

type Resultwords struct {
	Sequence string //'json:"sequence"'
	Nun      int
	Words    *map[string]bool //'json:"words"'
}
