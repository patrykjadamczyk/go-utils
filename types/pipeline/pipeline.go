package pipeline

import (
	"reflect"

	"github.com/patrykjadamczyk/go-utils/utils"
)

type Pipeline struct {
	Function any
    Data []any
}

func Value(value any) Pipeline {
	if utils.IsFunc(value) == false {
		panic("Value is not a function")
	}
	return Pipeline{Function: value}
}

func ValueFromPointer(value *any) Pipeline {
	if utils.IsFunc(*value) == false {
		panic("Value is not a function")
	}
	return Value(*value)
}

func (p *Pipeline) Run(args ...any) []any {
    return p.Evaluate(args...).Data
}

func (p *Pipeline) Evaluate(args ...any) *Pipeline {
	p.Data = []any{}
	var returns []reflect.Value
	if len(args) == 0 {
		returns = utils.CallFuncWithoutArgs(p.Function)
	} else {
		returns = utils.CallFunc(p.Function, args...)
	}
    p.Data = utils.ReflectValueToValue(returns...)
	return p
}

func (p *Pipeline) Call(args ...any) *Pipeline {
	return p.Evaluate(args...)
}

func (p *Pipeline) EvaluateOnData() *Pipeline {
	oldData := p.Data
	var returns []reflect.Value
	if len(oldData) == 0 {
		returns = utils.CallFuncWithoutArgs(p.Function)
	} else {
		returns = utils.CallFunc(p.Function, oldData...)
	}
	p.Data = utils.ReflectValueToValue(returns...)
	return p
}

func (p *Pipeline) Value() []any {
	return p.Data
}

func (p *Pipeline) ValueRef(returns ...*any) {
	for i, v := range p.Data {
		*returns[i] = v
	}
}

func (p *Pipeline) IsError() bool {
	if len(p.Data) == 0 {
		return false
	}
	if len(p.Data) == 1 {
		switch p.Data[0].(type) {
		case utils.NilError:
			return false
		case error:
			return true
		default:
			return false
		}
	}
	switch p.Data[len(p.Data)-1].(type) {
	case utils.NilError:
		return false
	case error:
		return true
	default:
		return false
	}
}

func (p *Pipeline) GetError() error {
	if len(p.Data) == 0 {
		return nil
	}
	if len(p.Data) == 1 {
		switch p.Data[0].(type) {
		case utils.NilError:
			return nil
		case error:
			return utils.EnsureType[error](p.Data[0])
		default:
			return nil
		}
	}
	switch p.Data[len(p.Data)-1].(type) {
	case utils.NilError:
		return nil
	case error:
		return utils.EnsureType[error](p.Data[len(p.Data)-1])
	default:
		return nil
	}
}

func (p *Pipeline) GetErrorSpecial() error {
	err := p.GetError()
	if err == nil {
		return utils.NilErrorErr
	}
	return err
}

func (p *Pipeline) Then(f any) *Pipeline {
	if utils.IsFunc(f) == false {
		panic("Value is not a function")
	}
	if p.IsError() {
		return p
	}
	p.Function = f
	p.EvaluateOnData()
	return p
}

func (p *Pipeline) Catch(f any) *Pipeline {
	if utils.IsFunc(f) == false {
		panic("Value is not a function")
	}
	if p.IsError() == false {
		return p
	}
	p.Function = f
	returns := utils.CallFunc(p.Function, p.GetErrorSpecial())
	p.Data = utils.ReflectValueToValue(returns...)
	return p
}

func (p *Pipeline) Finally(f any) *Pipeline {
	if utils.IsFunc(f) == false {
		panic("Value is not a function")
	}
	p.Function = f
	p.EvaluateOnData()
	return p
}

