package main

import (
	"testing"
)

var int64Number int64 = 5

func BenchmarkAbs_int64(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = int64(Abs(float64(int64Number)))

	}
}

func BenchmarkGenericAbsReturnFloat_int64(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = int64(GenericAbs(int64Number))
	}
}

func BenchmarkGenericAbsReturnT_int64(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = GenericAbsT(int64Number)
	}
}
