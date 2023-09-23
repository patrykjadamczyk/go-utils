package convert_test

import (
	. "github.com/patrykjadamczyk/go-utils/base"
    . "github.com/patrykjadamczyk/go-utils/unit"
	"testing"
)

func TestConvert(t *testing.T) {
    Convertions.AddIntConverter(
        "meters",
        "kilometers",
        func(value any) Result[UnitValue[int64]] {
            v := NormalizeUnit[int64]("meters")(value)
            if v.IsError() {
                return Err[UnitValue[int64]](v.GetError())
            }
            return Ok[UnitValue[int64]](
                Unit[int64](v.Unwrap().GetValue() / 1000, "kilometers"),
            )
        },
    )
    Convertions.AddUintConverter(
        "meters",
        "kilometers",
        func(value any) Result[UnitValue[uint64]] {
            v := NormalizeUnit[uint64]("meters")(value)
            if v.IsError() {
                return Err[UnitValue[uint64]](v.GetError())
            }
            return Ok[UnitValue[uint64]](
                Unit[uint64](v.Unwrap().GetValue() / 1000, "kilometers"),
            )
        },
    )
    Convertions.AddFloatConverter(
        "meters",
        "kilometers",
        func(value any) Result[UnitValue[float64]] {
            v := NormalizeUnit[float64]("meters")(value)
            if v.IsError() {
                return Err[UnitValue[float64]](v.GetError())
            }
            return Ok[UnitValue[float64]](
                Unit[float64](v.Unwrap().GetValue() / 1000, "kilometers"),
            )
        },
    )

    ti := Unit[int64](2000, "meters")
    tu := Unit[uint64](2000, "meters")
    tf := Unit[float64](2000, "meters")
    t1 := ConvertInt(ti, MakeEmptyUnit[int64]("kilometers"))
    t2 := ConvertUint(tu, MakeEmptyUnit[uint64]("kilometers"))
    t3 := ConvertFloat(tf, MakeEmptyUnit[float64]("kilometers"))
    if t1.IsError() {
        t.Error(t1.GetError())
    }
    if t2.IsError() {
        t.Error(t2.GetError())
    }
    if t3.IsError() {
        t.Error(t3.GetError())
    }
    t1r := t1.Unwrap()
    t2r := t2.Unwrap()
    t3r := t3.Unwrap()
    if t1r.GetValue() != 2 {
        t.Error("t1 should be 2km")
    }
    if t2r.GetValue() != 2 {
        t.Error("t2 should be 2km")
    }
    if t3r.GetValue() != 2 {
        t.Error("t3 should be 2km")
    }
}
