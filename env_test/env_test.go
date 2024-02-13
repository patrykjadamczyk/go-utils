package env_test

import (
	. "github.com/patrykjadamczyk/go-utils/env"
	"testing"
)

func TestEnv(t *testing.T) {
	SetEnv("TEST", "1")
	if GetEnv("TEST") != "1" {
		t.Error("GetEnv and SetEnv should work")
	}
}
