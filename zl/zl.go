package zl

import (
	"time"
	"fmt"
)
/*
Package for generic logging
*/

var ticker int64 = -1


//public interface that all loggers should follow
type Log interface {
	Write(f string, values ...interface{})
}



//public tick to poll time behind the scenes
func Tick() {
	ticker = time.Now().UnixNano() / int64(time.Millisecond)
}

//tick measures time since last call to log actions, can be manually overridden with zl.Tick()
func tick() string {
	var since int64 = 0
	var unix int64 = time.Now().UnixNano() / int64(time.Millisecond)
	if ticker > 0 {
		since = unix - ticker
	}
	ticker = unix
	return fmt.Sprintf("%d", since)
}


//generic prefix for all logged events
func formatPrefix() string {
	time := time.Now().Format("2006.01.02 15:04:05")
	lastTick := tick()
	return fmt.Sprintf("[%s]\t[%s]: \t", lastTick, time)
}


func SpacerTo(i Log) {
	i.Write("-----------------------------");
}