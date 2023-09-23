package unit

import (
	. "github.com/patrykjadamczyk/go-utils/base"
)

type ConvertionFuncTyped[T any] func(any) Result[UnitValue[T]]
type DoubleMap[T any] map[string]map[string]T

type Converters struct {
    FloatConverters DoubleMap[ConvertionFuncTyped[float64]]
    IntConverters   DoubleMap[ConvertionFuncTyped[int64]]
    UintConverters   DoubleMap[ConvertionFuncTyped[uint64]]
}

func (c *Converters) AddFloatConverter(
    unitFrom string,
    unitTo string,
    converter func(any) Result[UnitValue[float64]],
) {
    if c.FloatConverters == nil {
        c.FloatConverters = make(DoubleMap[ConvertionFuncTyped[float64]])
    }
    if c.FloatConverters[unitFrom] == nil {
        c.FloatConverters[unitFrom] = make(map[string]ConvertionFuncTyped[float64])
    }
    c.FloatConverters[unitFrom][unitTo] = converter
}

func (c *Converters) AddIntConverter(
    unitFrom string,
    unitTo string,
    converter func(any) Result[UnitValue[int64]],
) {
    if c.IntConverters == nil {
        c.IntConverters = make(DoubleMap[ConvertionFuncTyped[int64]])
    }
    if c.IntConverters[unitFrom] == nil {
        c.IntConverters[unitFrom] = make(map[string]ConvertionFuncTyped[int64])
    }
    c.IntConverters[unitFrom][unitTo] = converter
}

func (c *Converters) AddUintConverter(
    unitFrom string,
    unitTo string,
    converter func(any) Result[UnitValue[uint64]],
) {
    if c.UintConverters == nil {
        c.UintConverters = make(DoubleMap[ConvertionFuncTyped[uint64]])
    }
    if c.UintConverters[unitFrom] == nil {
        c.UintConverters[unitFrom] = make(map[string]ConvertionFuncTyped[uint64])
    }
    c.UintConverters[unitFrom][unitTo] = converter
}

var ConverterNotFoundError = NewError("Converter not found")

func (c *Converters) GetFloatConverter(
    unitFrom string,
    unitTo string,
) Result[ConvertionFuncTyped[float64]] {
    if c.FloatConverters == nil {
        return Err[ConvertionFuncTyped[float64]](ConverterNotFoundError)
    }
    if c.FloatConverters[unitFrom] == nil {
        return Err[ConvertionFuncTyped[float64]](ConverterNotFoundError)
    }
    return Ok(c.FloatConverters[unitFrom][unitTo])
}

func (c *Converters) GetIntConverter(
    unitFrom string,
    unitTo string,
) Result[ConvertionFuncTyped[int64]] {
    if c.IntConverters == nil {
        return Err[ConvertionFuncTyped[int64]](ConverterNotFoundError)
    }
    if c.IntConverters[unitFrom] == nil {
        return Err[ConvertionFuncTyped[int64]](ConverterNotFoundError)
    }
    return Ok(c.IntConverters[unitFrom][unitTo])
}

func (c *Converters) GetUintConverter(
    unitFrom string,
    unitTo string,
) Result[ConvertionFuncTyped[uint64]] {
    if c.UintConverters == nil {
        return Err[ConvertionFuncTyped[uint64]](ConverterNotFoundError)
    }
    if c.UintConverters[unitFrom] == nil {
        return Err[ConvertionFuncTyped[uint64]](ConverterNotFoundError)
    }
    return Ok(c.UintConverters[unitFrom][unitTo])
}

var Convertions = Converters{}

func ConvertFloat(
    valueToConvert any,
    valueUnitToConvertTo UnitValue[float64],
) Result[UnitValue[float64]] {
    value := NormalizeUnit[float64](valueUnitToConvertTo.UnitName)(valueToConvert)
    if value.IsError() {
        if IsErrorEqual(value.GetError(), UnitInvalidUnitError) {
            if val, ok := valueToConvert.(UnitValue[float64]); ok {
                cfr := Convertions.GetFloatConverter(val.GetUnitName(), valueUnitToConvertTo.UnitName)
                if cfr.IsError() {
                    return Err[UnitValue[float64]](cfr.GetError())
                }
                cf := cfr.Unwrap()
                if cf != nil {
                    v := cf(val.Value)
                    if v.IsError() {
                        return Err[UnitValue[float64]](v.GetError())
                    }
                    return Ok[UnitValue[float64]](v.Unwrap())
                }
                return Err[UnitValue[float64]](NewError("Convertion not Found"))
            }
        }
        return Err[UnitValue[float64]](value.GetError())
    }
    val := value.Unwrap()
    cfr := Convertions.GetFloatConverter(val.GetUnitName(), valueUnitToConvertTo.UnitName)
    if cfr.IsError() {
        return Err[UnitValue[float64]](cfr.GetError())
    }
    cf := cfr.Unwrap()
    if cf != nil {
        v := cf(val.Value)
        if v.IsError() {
            return Err[UnitValue[float64]](v.GetError())
        }
        return Ok[UnitValue[float64]](
            EnsureType[UnitValue[float64]](v.Unwrap()),
        )
    }
    return Err[UnitValue[float64]](NewError("Convertion not Found"))
}

func ConvertInt(
    valueToConvert any,
    valueUnitToConvertTo UnitValue[int64],
) Result[UnitValue[int64]] {
    value := NormalizeUnit[int64](valueUnitToConvertTo.UnitName)(valueToConvert)
    if value.IsError() {
        if IsErrorEqual(value.GetError(), UnitInvalidUnitError) {
            if val, ok := valueToConvert.(UnitValue[int64]); ok {
                cfr := Convertions.GetIntConverter(val.GetUnitName(), valueUnitToConvertTo.UnitName)
                if cfr.IsError() {
                    return Err[UnitValue[int64]](cfr.GetError())
                }
                cf := cfr.Unwrap()
                if cf != nil {
                    v := cf(val.Value)
                    if v.IsError() {
                        return Err[UnitValue[int64]](v.GetError())
                    }
                    return Ok[UnitValue[int64]](v.Unwrap())
                }
                return Err[UnitValue[int64]](NewError("Convertion not Found"))
            }
        }
        return Err[UnitValue[int64]](value.GetError())
    }
    val := value.Unwrap()
    cfr := Convertions.GetIntConverter(val.GetUnitName(), valueUnitToConvertTo.UnitName)
    if cfr.IsError() {
        return Err[UnitValue[int64]](cfr.GetError())
    }
    cf := cfr.Unwrap()
    if cf != nil {
        v := cf(val.Value)
        if v.IsError() {
            return Err[UnitValue[int64]](v.GetError())
        }
        return Ok[UnitValue[int64]](
            EnsureType[UnitValue[int64]](v.Unwrap()),
        )
    }
    return Err[UnitValue[int64]](NewError("Convertion not Found"))
}

func ConvertUint(
    valueToConvert any,
    valueUnitToConvertTo UnitValue[uint64],
) Result[UnitValue[uint64]] {
    value := NormalizeUnit[uint64](valueUnitToConvertTo.UnitName)(valueToConvert)
    if value.IsError() {
        if IsErrorEqual(value.GetError(), UnitInvalidUnitError) {
            if val, ok := valueToConvert.(UnitValue[uint64]); ok {
                cfr := Convertions.GetUintConverter(val.GetUnitName(), valueUnitToConvertTo.UnitName)
                if cfr.IsError() {
                    return Err[UnitValue[uint64]](cfr.GetError())
                }
                cf := cfr.Unwrap()
                if cf != nil {
                    v := cf(val.Value)
                    if v.IsError() {
                        return Err[UnitValue[uint64]](v.GetError())
                    }
                    return Ok[UnitValue[uint64]](v.Unwrap())
                }
                return Err[UnitValue[uint64]](NewError("Convertion not Found"))
            }
        }
        return Err[UnitValue[uint64]](value.GetError())
    }
    val := value.Unwrap()
    cfr := Convertions.GetUintConverter(val.GetUnitName(), valueUnitToConvertTo.UnitName)
    if cfr.IsError() {
        return Err[UnitValue[uint64]](cfr.GetError())
    }
    cf := cfr.Unwrap()
    if cf != nil {
        v := cf(val.Value)
        if v.IsError() {
            return Err[UnitValue[uint64]](v.GetError())
        }
        return Ok[UnitValue[uint64]](
            EnsureType[UnitValue[uint64]](v.Unwrap()),
        )
    }
    return Err[UnitValue[uint64]](NewError("Convertion not Found"))
}
