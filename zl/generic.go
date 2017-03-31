package zl

import "fmt"

/*
generic STDOUT logging
*/

type GenericLog struct {}


func (l GenericLog) Write(f string, values ...interface{}) {
	fmt.Printf(formatPrefix()+f+"\n", values...)
}