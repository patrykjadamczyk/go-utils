package pipeline

import (
	"github.com/patrykjadamczyk/go-utils/errors"
)

type PipelineArray struct {
	Pipelines []Pipeline
}

func NewPipelineArray(pipelines ...Pipeline) PipelineArray {
	return PipelineArray{Pipelines: pipelines}
}

func (p PipelineArray) RunAll(args ...[]any) []any {
	result := make([]any, len(p.Pipelines))
	for i, p := range p.Pipelines {
		result[i] = p.Run(args[i]...)
	}
	return result
}

func (p PipelineArray) All(args ...[]any) Pipeline {
	return Value(func() ([]any, error) {
		result := make([]any, len(p.Pipelines))
		for i, p := range p.Pipelines {
			p.Evaluate(args[i]...)
			if p.IsError() {
				return []any{i}, p.GetErrorSpecial()
			}
			result[i] = p.Value()
		}
		return result, errors.NilErrorErr
	})
}

type PipelineArrayAnyError struct {
}

func (PipelineArrayAnyError) Error() string {
	return "Pipeline Any Error"
}

func (p PipelineArray) Any(args ...[]any) Pipeline {
	return Value(func() ([]any, []error, error) {
		result := make([]any, len(p.Pipelines))
		errs := make([]error, len(p.Pipelines))
		for i, p := range p.Pipelines {
			p.Evaluate(args[i]...)
			if p.IsError() == false {
				return p.Value(), []error{}, errors.NilErrorErr
			}
			result[i] = p.Value()
			errs[i] = p.GetError()
		}
		return result, errs, PipelineArrayAnyError{}
	})
}
