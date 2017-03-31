package main

import (
	"fmt"
	"strings"
	"github.com/zutto/zlog/zl"
)

const
(
	spacerLen    = 20
	spacerString = "-"
)

func main() {

	var multi zl.Log = zl.MultiLog{[]zl.Log{zl.GenericLog{}, zl.FileLog{Location: "/tmp/golang-test.log" }}}

	multi.Write("%s", spacers())
	multi.Write("this is a test %v %v %v", 123, "string", func() { /* dummy */ })
	multi.Write(spacers())

	multi.Write("second test", nil)
	multi.Write("%v", spacers())

	defer func() {
		if error := recover(); error != nil {
			fmt.Printf("recovered from crash")
		}
	}()

}

func spacers() string {
	return strings.Repeat(spacerString, spacerLen)
}
