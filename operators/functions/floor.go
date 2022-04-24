package functions

import (
	"math"

	"github.com/alfredxing/calc/operators"
)

var (
	floor = &operators.Operator{
		Name:          "floor",
		Precedence:    0,
		Associativity: operators.L,
		Args:          1,
		Operation: func(args []float64) float64 {
			return math.Floor(args[0])
		},
	}
	ceil = &operators.Operator{
		Name:          "ceil",
		Precedence:    0,
		Associativity: operators.L,
		Args:          1,
		Operation: func(args []float64) float64 {
			return math.Ceil(args[0])
		},
	}
	trunc = &operators.Operator{
		Name:          "trunc",
		Precedence:    0,
		Associativity: operators.L,
		Args:          1,
		Operation: func(args []float64) float64 {
			return math.Trunc(args[0])
		},
	}
	round = &operators.Operator{
		Name:          "round",
		Precedence:    0,
		Associativity: operators.L,
		Args:          1,
		Operation: func(args []float64) float64 {
			return math.Round(args[0])
		},
	}
)

func init() {
	Register(round)
	Register(floor)
	Register(ceil)
	Register(trunc)
}
