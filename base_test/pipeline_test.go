package base_test

import (
	"errors"
	"testing"

	. "github.com/patrykjadamczyk/go-utils/base"
)

func TestPipeline(t *testing.T) {
	p := MakePipeline(func(v int) int {
		return v + 1
	})
	pd := p.Run(1)
	if len(pd) != 1 {
		t.Error("Pipeline result should have 1 element")
	}
	if pd[0] != 2 {
		t.Error("Pipeline result should be 2")
	}
	p2 := MakePipeline(func(v int) int {
		return v + 2
	})
	p2 = *p2.Evaluate(2).EvaluateOnData()
	pd2 := p2.Value()
	if len(pd2) != 1 {
		t.Error("Pipeline result 2 should have 1 element")
	}
	if pd2[0] != 6 {
		t.Error("Pipeline result 2 should be 6")
	}
	if p2.IsError() {
		t.Error("Pipeline 2 should not have error")
	}
	p3 := MakePipeline(func(v int) int {
		return v + 3
	})
	p3 = *p3.Evaluate(3)
	p3 = *p3.Then(func(v int) int {
		return v + 4
	}).Catch(func(err error) {
		t.Error("Pipeline 3 should have an error", err)
	}).Finally(func(v int) int {
		return v + 5
	})
	p3d := p3.Value()
	if p3.IsError() {
		t.Error("Pipeline 3 should not have error")
	}
	if len(p3d) != 1 {
		t.Error("Pipeline 3 result should have 1 element")
	}
	if p3d[0] != 15 {
		t.Error("Pipeline 3 result should be 15")
	}
}

func TestPipelineArray(t *testing.T) {
	p1 := MakePipeline(func(v int) int {
		return v + 2
	})
	p2 := MakePipeline(func(v int) int {
		return v + 3
	})
	p := NewPipelineArray(p1, p2)
	pd := p.RunAll([]any{2}, []any{2})
	if len(pd) != 2 {
		t.Error("Pipeline array result should have 2 elements")
	}
	if EnsureType[[]any](pd[0])[0] != 4 {
		t.Error("Pipeline array result 1 should be 4")
	}
	if EnsureType[[]any](pd[1])[0] != 5 {
		t.Error("Pipeline array result 2 should be 5")
	}
}

func TestPipelineArrayAll(t *testing.T) {
	p1 := MakePipeline(func(v int) int {
		return v + 2
	})
	p2 := MakePipeline(func(v int) (int, error) {
		return v + 3, errors.New("Hello")
	})
	p := NewPipelineArray(p1, p2)
	pr := p.All([]any{2}, []any{2})
	pd := pr.Run()
	if len(pd) != 2 {
		t.Error("Pipeline array result should have 2 elements")
	}
	if pr.IsError() != true {
		t.Error("Pipeline array should have error")
	}
}

func TestPipelineArrayAny(t *testing.T) {
	p1 := MakePipeline(func(v int) int {
		return v + 2
	})
	p2 := MakePipeline(func(v int) (int, error) {
		return v + 3, errors.New("Hello")
	})
	p := NewPipelineArray(p1, p2)
	pr := p.Any([]any{2}, []any{2})
	pd := pr.Run()
	if len(pd) != 3 {
		t.Error("Pipeline array result should have 3 elements")
	}
	if pr.IsError() != false {
		t.Error("Pipeline array should not have error")
	}
	if EnsureType[[]any](pd[0])[0] != 4 {
		t.Error("Pipeline array result 1 should be 4")
	}
}

func TestPipelineFunc(t *testing.T) {
	add1 := func(t int) int {
		return t + 1
	}
	pt1 := Pipe(1, add1, add1)
	if pt1 != any(3) {
		t.Error("Pipeline result 1 should be 3")
	}
	pt2 := PipeTypedReturn[int](1, add1, add1)
	if pt2 != 3 {
		t.Error("Pipeline result 2 should be 3")
	}
	pt3 := PipeTypedArgReturn[int, int](1, add1, add1)
	if pt3 != 3 {
		t.Error("Pipeline result 3 should be 3")
	}
	pt4 := PipeTyped[int](1, add1, add1)
	if pt4 != 3 {
		t.Error("Pipeline result 4 should be 3")
	}
}
