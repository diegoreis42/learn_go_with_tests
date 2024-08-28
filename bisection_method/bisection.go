package bisection

type fn func(float64) float64

func Bisection(f fn, a, b, e float64) [2]float64 {
	for {
		if b-a < e {
			return [2]float64{a, b}
		} else {
			m := f(a)
			x := (a + b) / 2

			if m*f(x) > 0 {
				a = x
			} else {
				b = x
			}
		}
	}
}
