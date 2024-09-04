package bisection

import (
	"math"
	"testing"
)

func TestBisection(t *testing.T) {
	t.Run("find root for x³ - 9x + 3", func(t *testing.T) {
		function := func(x float64) float64 {
			return math.Pow(x, 3) - 9*x + 3
		}

		got := Bisection(function, 1.0, 3.0, 0.02)
		givenPrecision := got[1] - got[0]
		want := 0.02

		if givenPrecision > want {
			t.Errorf("%v want %f, got %f", got, want, givenPrecision)
		}
	})
}
