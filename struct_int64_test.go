package main

import (
	"testing"
)

type int64Struct struct{ num int64 }

var int64StructNumber int64Struct = int64Struct{num: 5}

func BenchmarkAbs_structInt64(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = int64(Abs(float64(int64StructNumber.num)))
	}
}

func BenchmarkGenericAbsReturnFloat_structInt64(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = int64(GenericAbs(int64StructNumber.num))
	}
}

func BenchmarkGenericAbsReturnT_structInt64(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = GenericAbsT(int64StructNumber.num)
	}
}
