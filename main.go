package main

import (
	"math"
)

func main() {}

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

func Abs(x float64) float64 {
	return math.Float64frombits(math.Float64bits(x) &^ (1 << 63))
}

func GenericAbs[T Number](x T) float64 {
	return math.Float64frombits(math.Float64bits(float64(x)) &^ (1 << 63))
}

func GenericAbsT[T Number](x T) T {
	return T(math.Float64frombits(math.Float64bits(float64(x)) &^ (1 << 63)))
}
