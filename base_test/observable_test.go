package base_test

import (
	. "github.com/patrykjadamczyk/go-utils/base"
	"testing"
)

func TestObservable(t *testing.T) {
	obs := MakeObservable[bool](false)
	if obs.Value != false {
		t.Error("Value should be false")
	}
	obs.AddGetHandler(func(obs Observable[bool]) {
		if obs.Value != true {
			t.Error("Value should be true")
		}
	})
	obs.Value = true
	if obs.GetValue() != true {
		t.Error("Value should be true")
	}
	obs.AddSetHandler(func(obs Observable[bool]) {
		if obs.Value != false {
			t.Error("Value should be false")
		}
	})
	obs.SetValue(false)
	if obs.Value != false {
		t.Error("Value should be false")
	}
}
