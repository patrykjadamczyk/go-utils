package base_test

import (
	. "github.com/patrykjadamczyk/go-utils/base"
	"testing"
)

func TestFeatureFlags(t *testing.T) {
	GetFeatureFlags().SetIntFlag("test", 1)
	if GetFeatureFlags().GetIntFlag("test") != 1 {
		t.Error("GetIntFlag on test should be 1")
	}
	res := GetFeatureFlags().Dev()
	if res.IsError() {
		t.Error("Result 1 should not be error", res.DataValue, res.ErrorValue)
	}
}
