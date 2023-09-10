package rust_types

import "math"

type f32 float32
type f64 float64
type i8 int8
type i16 int16
type i32 int32
type i64 int64
type str string
type u8 uint8
type u16 uint16
type u32 uint32
type u64 uint64

const f32Max = float32(math.MaxFloat32)
const f64Max = float64(math.MaxFloat64)

const i8Max = int8(math.MaxInt8)
const i16Max = int16(math.MaxInt16)
const i32Max = int32(math.MaxInt32)
const i64Max = int64(math.MaxInt64)

const u8Max = uint8(math.MaxUint8)
const u16Max = uint16(math.MaxUint16)
const u32Max = uint32(math.MaxUint32)
const u64Max = uint64(math.MaxUint64)

const i8Min = int8(math.MinInt8)
const i16Min = int16(math.MinInt16)
const i32Min = int32(math.MinInt32)
const i64Min = int64(math.MinInt64)

const u8Min = uint8(0)
const u16Min = uint16(0)
const u32Min = uint32(0)
const u64Min = uint64(0)
