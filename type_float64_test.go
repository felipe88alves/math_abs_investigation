package main

import (
	"testing"
)

type float64Type float64

var float64TypeNumber float64Type = 5

func BenchmarkAbs_typeFloat64(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = float64Type(Abs(float64(float64TypeNumber)))

	}
}

func BenchmarkGenericAbsReturnFloat_typeFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = float64Type(GenericAbs(float64TypeNumber))
	}
}

func BenchmarkGenericAbsReturnT_typeFloat64(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = GenericAbsT(float64TypeNumber)
	}
}
