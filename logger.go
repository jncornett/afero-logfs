package logfs

import "time"

type Record struct {
	Op   string
	Args []interface{}
	Err  error
	Time time.Duration
}

type Logger interface {
	Log(Record)
}

type LoggerFunc func(Record)

func (f LoggerFunc) Log(r Record) { f(r) }
