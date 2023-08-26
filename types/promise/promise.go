package promise

import (
	"sync"

	"github.com/patrykjadamczyk/go-utils/types/pipeline"
	"github.com/patrykjadamczyk/go-utils/utils"
)

type Promise struct {
	wg sync.WaitGroup
    Function any
    Result any
    Error error
}

func MakePromise(f any) Promise {
    if utils.IsFunc(f) == false {
        panic("Value is not a function")
    }
    return Promise{Function: f}
}

func NewPromise(f any) *Promise {
    p := MakePromise(f)
    return &p
}

func (p *Promise) Call(args ...any) pipeline.Pipeline {
    p.wg.Add(1)
    go func() {
        defer p.wg.Done()
        pl := pipeline.Value(p.Function)
        pl.Evaluate(args...)
        p.Result = pl.Data
        p.Error = pl.GetError()
    }()
    rpl := pipeline.Value(func() (any, error) {
        return p.Result, p.Error
    })
    return rpl
}

func (p *Promise) Wait() {
    p.wg.Wait()
}
