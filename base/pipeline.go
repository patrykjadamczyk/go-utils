package base

import (
	"reflect"

	"github.com/patrykjadamczyk/go-utils/errors"
	"github.com/patrykjadamczyk/go-utils/utils"
)

// Pipeline object
// You can think of this as promise like construct without any asynchronity or like |> operator
type Pipeline struct {
	Function any
	Data     []any
}

// Make New Pipeline from specified function
func MakePipeline(value any) Pipeline {
	if !utils.IsFunc(value) {
		panic("Value is not a function")
	}
	return Pipeline{Function: value}
}

// Make New Pipeline from specified pointer to function
func MakePipelineFromPointer(value *any) Pipeline {
	if !utils.IsFunc(*value) {
		panic("Value is not a function")
	}
	return MakePipeline(*value)
}

// Run Pipeline and Get underlying data
func (p *Pipeline) Run(args ...any) []any {
	return p.Evaluate(args...).Data
}

// Run Pipeline and chain next operations
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

// Run Pipeline and chain next operations (Same as Evaluate)
func (p *Pipeline) Call(args ...any) *Pipeline {
	return p.Evaluate(args...)
}

// Run Pipeline with arguments as data inside the pipeline and chain next operations
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

// Get underlying data of Pipeline
func (p *Pipeline) Value() []any {
	return p.Data
}

// Get underlying data of Pipeline inside variables specified as args
func (p *Pipeline) ValueRef(returns ...*any) {
	for i, v := range p.Data {
		*returns[i] = v
	}
}

// Check if Pipeline got error value in the data
func (p *Pipeline) IsError() bool {
	if len(p.Data) == 0 {
		return false
	}
	if len(p.Data) == 1 {
		switch p.Data[0].(type) {
		case ErrorableGenericResultInterface:
			return p.Data[0].(ErrorableGenericResultInterface).IsError()
		case errors.NilError:
			return false
		case error:
			return true
		default:
			return false
		}
	}
	switch p.Data[len(p.Data)-1].(type) {
	case ErrorableGenericResultInterface:
		return p.Data[len(p.Data)-1].(ErrorableGenericResultInterface).IsError()
	case errors.NilError:
		return false
	case error:
		return true
	default:
		return false
	}
}

// Get Error if Pipeline has error otherwise return nil
func (p *Pipeline) GetError() error {
	if len(p.Data) == 0 {
		return nil
	}
	if len(p.Data) == 1 {
		switch p.Data[0].(type) {
		case ErrorableGenericResultInterface:
			return p.Data[0].(ErrorableGenericResultInterface).GetError()
		case errors.NilError:
			return nil
		case error:
			return EnsureType[error](p.Data[0])
		default:
			return nil
		}
	}
	switch p.Data[len(p.Data)-1].(type) {
	case ErrorableGenericResultInterface:
		return p.Data[len(p.Data)-1].(ErrorableGenericResultInterface).GetError()
	case errors.NilError:
		return nil
	case error:
		return EnsureType[error](p.Data[len(p.Data)-1])
	default:
		return nil
	}
}

// Get Error if it's nil return NilErrorErr Structure
func (p *Pipeline) GetErrorSpecial() error {
	err := p.GetError()
	if err == nil {
		return errors.NilErrorErr
	}
	return err
}

// Do something if Pipeline doesn't have any error
func (p *Pipeline) Then(f any) *Pipeline {
	if !utils.IsFunc(f) {
		panic("Value is not a function")
	}
	if p.IsError() {
		return p
	}
	p.Function = f
	p.EvaluateOnData()
	return p
}

// Do something if Pipeline has error
func (p *Pipeline) Catch(f any) *Pipeline {
	if !utils.IsFunc(f) {
		panic("Value is not a function")
	}
	if !p.IsError() {
		return p
	}
	p.Function = f
	returns := utils.CallFunc(p.Function, p.GetErrorSpecial())
	p.Data = utils.ReflectValueToValue(returns...)
	return p
}

// Do something no matter if Pipeline has error or no
func (p *Pipeline) Finally(f any) *Pipeline {
	if !utils.IsFunc(f) {
		panic("Value is not a function")
	}
	p.Function = f
	p.EvaluateOnData()
	return p
}
