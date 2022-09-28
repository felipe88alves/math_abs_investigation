package main

import (
	"testing"
)

var float64Number float64 = 5

func BenchmarkAbs_float64(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {

		_ = Abs(float64Number)

	}
}

func BenchmarkGenericAbsReturnFloat_float64(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = float64(GenericAbs(float64Number))
	}
}

func BenchmarkGenericAbsReturnT_float64(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = GenericAbsT(float64Number)
	}
}
