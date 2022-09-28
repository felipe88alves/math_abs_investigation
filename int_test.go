package main

import (
	"testing"
)

var intNumber int = 5

func BenchmarkAbs_int(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {

		_ = int(Abs(float64(intNumber)))

	}
}

func BenchmarkGenericAbsReturnFloat_int(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = int(GenericAbs(intNumber))
	}
}

func BenchmarkGenericAbsReturnT_int(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = GenericAbsT(intNumber)
	}
}
