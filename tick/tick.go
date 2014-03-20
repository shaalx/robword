package tick

import (
	"time"
)

func Tick(t time.Duration) {
	tick := time.Tick(t)
	<-tick
}
