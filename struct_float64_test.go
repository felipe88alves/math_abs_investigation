package main

import (
	"testing"
)

type float64Struct struct{ num float64 }

var float64StructNumber float64Struct = float64Struct{num: 5}

func BenchmarkAbs_structFloat64(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = Abs(float64StructNumber.num)

	}
}

func BenchmarkGenericAbsReturnFloat_structFloat64(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = float64(GenericAbs(float64StructNumber.num))
	}
}

func BenchmarkGenericAbsReturnT_structFloat64(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = GenericAbsT(float64StructNumber.num)
	}
}
