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
