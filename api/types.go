package api

import "fmt"

type Processor interface {
	GetName() string
	Exec() error
	Init() error
	Uninit() error
}

type Source interface {
	GetOutput() <-chan interface{}
}

type LogSource interface {
	GetLogs() <-chan interface{}
}

type Sink interface {
	GetInput() <-chan interface{}
	Done() <-chan struct{}
}

type Collector interface {
	GetInputs() []<-chan interface{}
}

type Emitter interface {
	GetOutputs() []<-chan interface{}
}

type ProcError struct {
	Err      error
	ProcName string
}

func (e ProcError) Error() string {
	if e.ProcName != "" {
		return fmt.Sprintf("[%s] %v", e.ProcName, e.Err)
	}
	return e.Err.Error()
}
