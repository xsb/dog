package types

import "time"

type ExecutionEvent interface{}

type TaskStartEvent struct {
	Task      string
	StartTime time.Time
}

type OutputEvent struct {
	Task string
	Body []byte
}

type TaskEndEvent struct {
	Task       string
	EndTime    time.Time
	StatusCode int
}
