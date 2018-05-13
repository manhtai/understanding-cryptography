package main

import (
	"fmt"

	"github.com/manhtai/understanding-cryptography/pkg"
)

// kary compute x^e % n using k-ary exponentiation algorithm,
// and handle k bits at a time
func kary(x, e, m, k int) (r int) {
	l := pkg.BitLen(e) / k

	tableLen := 1 << uint(k)
	table := make([]int, tableLen)

	// Build lookup table
	table[0] = 1
	table[1] = x
	for i := 2; i < tableLen; i++ {
		table[i] = (table[i-1] * x) % m
	}

	// Main iteration
	r = table[e>>uint(k*l)]
	for i := 0; i < l; i++ {
		// SQ step
		for j := 0; j < k; j++ {
			r = (r * r) % m
		}

		// MUL step
		idx := (e >> uint(k*(l-i-1))) & (tableLen - 1)
		if idx != 0 {
			r = (r * table[idx]) % m
		}
	}

	return
}

// karyPrint compute kary with k = 3 & print result
func karyPrint(x, e, m int) {
	r := kary(x, e, m, 3)
	fmt.Printf("%d^%d mod %d = %d\n", x, e, m, r)
}

func uc715() {
	fmt.Println("======= (7.15) ======")
	karyPrint(2, 79, 101)
	karyPrint(3, 197, 101)
	karyPrint(12, 145, 163)
}
