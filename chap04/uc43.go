package main

import (
	"fmt"
	"strings"
)

// GF2 represents order of P(x) from right to left
// i.e. let P(x) = x^2 then GF2 = []uint8{0, 0, 1}
type GF2 []uint8

// gfx generates multiplication table for GF(2^m) with irreducible polynomial p
func gfx(m uint8, p GF2) (t [][]GF2) {
	s := 1 << m

	// Init slice
	t = make([][]GF2, s)
	for i := 0; i < s; i++ {
		t[i] = make([]GF2, s)
	}

	// Calculate
	int2gf2 := int2GF2(s)
	var tij, k GF2
	for i := 0; i < s; i++ {
		for j := i; j < s; j++ {
			tij = mulGF2(int2gf2(i), int2gf2(j))
			for degGF2(tij) > int(m)-1 {
				k = divGF2(tij, p)
				tij = addGF2(tij, mulGF2(p, k))
			}
			t[i][j] = tij
			t[j][i] = tij
		}
	}

	return
}

// int2GF2 convert integer to corresponding GF2 representation,
// i.e. 1 => [1],  2 => [0, 1],  3 => [1, 1]
// it uses function closure to cache results
func int2GF2(s int) func(int) GF2 {
	cache := make([]GF2, s)

	return func(n int) (gf2 GF2) {
		if gf2 = cache[n]; gf2 != nil {
			return gf2
		}
		gf2 = make(GF2, s)
		for i := 0; i < s; i++ {
			gf2[i] = uint8((n >> uint(i)) & 1)
		}
		cache[n] = gf2
		return gf2
	}
}

func mulGF2(g, h GF2) (k GF2) {
	k = GF2{}

	for i := 0; i < len(g); i++ {
		if g[i] == 1 {
			hLen := len(h)
			hPrime := make(GF2, i+hLen)
			for j := 0; j < i; j++ {
				hPrime[j] = 0
			}
			for j := i; j < i+hLen; j++ {
				hPrime[j] = h[j-i]
			}
			k = addGF2(k, hPrime)
		}
	}

	return k
}

func divGF2(g, h GF2) (k GF2) {
	gDeg := degGF2(g)
	hDeg := degGF2(h)
	var d int
	if gDeg >= hDeg {
		d = gDeg - hDeg
	}
	k = make(GF2, d+1)
	k[d] = 1
	return k
}

func degGF2(g GF2) (d int) {
	for i := len(g) - 1; i >= 0; i-- {
		if g[i] == 1 {
			return i
		}
	}
	return 0
}

func addGF2(g, h GF2) (k GF2) {
	max := len(h)

	if len(g) > max {
		max = len(g)
	}

	k = make(GF2, max)

	for i := 0; i < max; i++ {
		var gi, hi uint8
		if i < len(g) {
			gi = g[i]
		}
		if i < len(h) {
			hi = h[i]
		}

		k[i] = (gi + hi) % 2
	}

	return k
}

func (g GF2) String() string {
	s := []string{}
	for i := range g {
		if g[i] == 0 {
			continue
		}

		if i == 0 {
			s = append([]string{"1"}, s...)
		} else if i == 1 {
			s = append([]string{"x"}, s...)
		} else {
			s = append([]string{fmt.Sprintf("x^%d", i)}, s...)
		}
	}

	if len(s) == 0 {
		s = []string{"0"}
	}

	return strings.Join(s, " + ")
}

func printGF2(gf [][]GF2) {
	for i := range gf {
		for j := range gf[i] {
			fmt.Printf("%d x %d = %v\n", i, j, gf[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func uc43() {
	// m = 3
	// P(x) = x^3 + x + 1
	printGF2(gfx(3, []uint8{1, 1, 0, 1}))
}
