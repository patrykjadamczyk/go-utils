package compareenum_test

import (
	"testing"

	compare "github.com/patrykjadamczyk/go-utils/base/CompareEnum"
	"github.com/patrykjadamczyk/go-utils/math"
)

func BenchmarkToCompareEnum(b *testing.B) {
	b.Run("stdlib", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			i1, i2 := math.GetRandomNumber[int](), math.GetRandomNumber[int]()
			b.StartTimer()
			if i1 > i2 {
				return
			}
			if i1 < i2 {
				return
			}
			if i1 == i2 {
				return
			}
		}
	})
	b.Run("ToCompareEnum", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			i1, i2 := math.GetRandomNumber[int](), math.GetRandomNumber[int]()
			b.StartTimer()
			comp := compare.ToCompareEnum(i1, i2)
			if comp == compare.Greater {
				return
			}
			if comp == compare.Less {
				return
			}
			if comp == compare.Equal {
				return
			}
		}
	})
	b.Run("stdlib-complex128", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			c1, c2 := math.GetRandomNumber[complex128](), math.GetRandomNumber[complex128]()
			b.StartTimer()
			if real(c1) > real(c2) && imag(c1) > imag(c2) {
				return
			}
			if real(c1) == real(c2) && imag(c1) == imag(c2) {
				return
			}
			if real(c1) < real(c2) && imag(c1) < imag(c2) {
				return
			}
			if real(c1) < real(c2) && imag(c1) >= imag(c2) {
				return
			}
		}
	})
	b.Run("ToCompareEnum-complex128", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			c1, c2 := math.GetRandomNumber[complex128](), math.GetRandomNumber[complex128]()
			b.StartTimer()
			comp := compare.ComplexToCompareEnum(c1, c2)
			if comp == compare.Greater {
				return
			}
			if comp == compare.Equal {
				return
			}
			if comp == compare.Less {
				return
			}
			if comp == compare.Uncomparable {
				return
			}
		}
	})
}
