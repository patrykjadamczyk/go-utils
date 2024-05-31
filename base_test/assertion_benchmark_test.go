package base_test

import (
	"testing"

	"github.com/patrykjadamczyk/go-utils/math"
	. "github.com/patrykjadamczyk/go-utils/base"
)

func BenchmarkAssertions(b *testing.B) {
	b.Run("stdlib", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			i := math.GetRandomNumber[float32]()
			if i != i {
				panic("NaN")
			}
		}
	})
	b.Run("Assert", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			i := math.GetRandomNumber[float32]()
			Assert(i == i)
		}
	})
	b.Run("Assert10", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			i := math.GetRandomNumber[float32]()
			Assert(i == i)
			Assert(i == i)
			Assert(i == i)
			Assert(i == i)
			Assert(i == i)
			Assert(i == i)
			Assert(i == i)
			Assert(i == i)
			Assert(i == i)
			Assert(i == i)
		}
	})
}
