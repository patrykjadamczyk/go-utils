package math

import "math"

type AnyFloat interface {
	~float32 | ~float64
}

func ensureType[T any](value any) T {
	switch value := value.(type) {
	case T:
		return value
	default:
		panic("Invalid Type")
	}
}

func IsInf[F AnyFloat](f F, sign int) bool {
	var maxFloat F = 0
	var v F
	value := any(v)
	switch value.(type) {
	case float64:
		maxFloat = ensureType[F](float64(math.MaxFloat64))
	case float32:
		maxFloat = ensureType[F](float32(math.MaxFloat32))
	}
	return sign >= 0 && f > maxFloat || sign <= 0 && f < -maxFloat
}

func Modf[F AnyFloat](f F) (int F, frac F) {
	var v F
	value := any(v)
	switch value.(type) {
	case float64:
		i, f := math.Modf(ensureType[float64](f))
		return ensureType[F](i), ensureType[F](f)
	case float32:
		i, f := math.Modf(ensureType[float64](float64(f)))
		return ensureType[F](float32(i)), ensureType[F](float32(f))
	}
	return
}

func Floor[F AnyFloat](x F) F {
	if x == 0 || IsNaN(x) || IsInf(x, 0) {
		return x
	}
	if x < 0 {
		d, fract := Modf(-x)
		if fract != 0.0 {
			d = d + 1
		}
		return -d
	}
	d, _ := Modf(x)
	return d
}

func Fmod[T ~float32 | ~float64](x, y T) T {
	if y == 0 {
		return NaN[T]()
	}

	q := x / y
	r := x - (Floor(q) * y)

	if r >= y {
		r -= y
	}
	if r < 0 {
		r += y
	}

	return r
}
