package zl

import (
	"fmt"
	"os"
)

/*
generic append-to-file logging
*/

type FileLog struct {
	Location string
}

func (l FileLog) Write(f string, values ...interface{}) {

	if (len(l.Location) <= 0) {
		panic("No file defined")
	}

	file, err := os.OpenFile(l.Location, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic("failed to open or create file")
	}

	defer func() {
		if error := recover(); error != nil {
			fmt.Printf("recovered")

			//panic continues
		}

		//guessing that these will make a fatal error
		file.Sync()
		file.Close()

	}()
	file.WriteString(fmt.Sprintf(formatPrefix()+f+"\n", values...))
}
