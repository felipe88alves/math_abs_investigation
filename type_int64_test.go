package main

import (
	"testing"
)

type int64Type int64

var int64TypeNumber int64Type = 5

func BenchmarkAbs_typeInt64(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = int64Type(Abs(float64(int64TypeNumber)))

	}
}

func BenchmarkGenericAbsReturnFloat_typeInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = int64Type(GenericAbs(int64TypeNumber))
	}
}

func BenchmarkGenericAbsReturnT_typeInt64(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = GenericAbsT(int64TypeNumber)
	}
}
