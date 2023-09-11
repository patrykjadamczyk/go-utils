package env_test

import (
    "testing"
    . "github.com/patrykjadamczyk/go-utils/env"
)

func TestEnv(t *testing.T) {
    SetEnv("TEST", "1")
    if GetEnv("TEST") != "1" {
        t.Error("GetEnv and SetEnv should work")
    }
}
