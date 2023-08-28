package base

import (
	"sync"

	"github.com/patrykjadamczyk/go-utils/utils"
)

type Promise struct {
	wg sync.WaitGroup
    Function any
    Result any
    Error error
}

// Make Promise
func MakePromise(f any) Promise {
    if utils.IsFunc(f) == false {
        panic("Value is not a function")
    }
    return Promise{Function: f}
}

// Make Promise Pointer
func MakePromisePointer(f any) *Promise {
    p := MakePromise(f)
    return &p
}

// Call Promise
func (p *Promise) Call(args ...any) Pipeline {
    p.wg.Add(1)
    go func() {
        defer p.wg.Done()
        pl := MakePipeline(p.Function)
        pl.Evaluate(args...)
        p.Result = pl.Data
        p.Error = pl.GetError()
    }()
    rpl := MakePipeline(func() (any, error) {
        return p.Result, p.Error
    })
    return rpl
}

// Wait until Promise is finished
func (p *Promise) Wait() {
    p.wg.Wait()
}
