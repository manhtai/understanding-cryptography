package main

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

// gcde compute gcd of two numbers using Extended Euclidean Algorithm
func gcde(m, n int) (s, t, g int) {
	s, t, g = gcdeInternal(m, n, 1, 0, 0, 1)
	// gcdeInternal switch m, n when recursive finding gcd, so we reorder s, t here
	// to correct order of parameters
	if m*s+n*t != g {
		s, t = t, s
	}
	return
}

func gcdePrint(m, n int) {
	s, t, g := gcde(m, n)
	fmt.Printf("[%d]x%d + [%d]x%d = %d = %d\n", s, m, t, n, s*m+t*n, g)
}

func uc66() {
	gcdePrint(198, 243)
	gcdePrint(1819, 3587)
}
