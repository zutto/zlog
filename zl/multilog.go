package zl

/*
Log to multiple engines simultaneously
*/

type MultiLog struct {
	LogEngines []Log
}


func (m MultiLog) Write(f string, values ...interface{}) {
	for _,engine := range m.LogEngines {
		engine.Write(f, values...)
	}
}