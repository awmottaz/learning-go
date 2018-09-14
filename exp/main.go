package exp

// Exp efficiently calculates a^p, where a and p are both integers.
func Exp(a, p int) float64 {
	if a == 0 || a == 1 {
		return float64(a)
	}
	if p == 0 {
		return 1
	}
	if p < 0 {
		den := Exp(a, -1*p)
		return 1 / den
	}
	powA := a
	ans := 1
	for i := uint(0); ; i++ {
		if (p>>i)&1 == 1 {
			ans *= powA
		}
		if 1<<i > p {
			break
		}
		powA *= powA
	}
	return float64(ans)
}

func expSlow(a, p int) float64 {
	if a == 0 || a == 1 {
		return float64(a)
	}
	if p == 0 {
		return 1
	}
	if p < 0 {
		den := expSlow(a, -1*p)
		return 1 / den
	}

	ans := 1
	for i := 0; i < p; i++ {
		ans *= a
	}
	return float64(ans)
}
