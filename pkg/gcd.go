package pkg

import "fmt"

func gcdeInternal(m, n, s0, t0, s1, t1 int) (s, t, g int) {
	if m < n {
		return gcdeInternal(n, m, t0, s0, t1, s1)
	}
	if n == 0 {
		s = s0
		t = t0
		g = m
		return
	}
	r := m % n
	q := (m - r) / n
	return gcdeInternal(r, n, s1, t1, s0-q*s1, t0-q*t1)
}

// Gcde compute gcd of two numbers using Extended Euclidean Algorithm
func Gcde(m, n int) (s, t, g int) {
	if n < 0 {
		n += m
	}
	s, t, g = gcdeInternal(m, n, 1, 0, 0, 1)
	// gcdeInternal switch m, n when recursive finding gcd, so we reorder s, t here
	// to correct order of parameters
	if m*s+n*t != g {
		s, t = t, s
	}

	// Make t not negative
	if t < 0 {
		t += m
	}
	return
}

// Inverse calculate a^-1 mod m
func Inverse(a, m int) (i int) {
	_, i, _ = Gcde(m, a)
	return
}

// GcdePrint calculate Gcde & print it out
func GcdePrint(m, n int) {
	s, t, g := Gcde(m, n)
	fmt.Printf("[%d]x%d + [%d]x%d = %d = %d\n", s, m, t, n, s*m+t*n, g)
}
