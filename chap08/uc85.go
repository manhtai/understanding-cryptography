package main

import (
	"fmt"

	"github.com/manhtai/understanding-cryptography/pkg"
)

func dhke(p, alpha, a, b int) (A, B, KAB int) {
	A = pkg.Snm(alpha, a, p)
	B = pkg.Snm(alpha, b, p)
	KAB = pkg.Snm(B, a, p)
	return
}

func dhkePrint(p, alpha, a, b int) {
	A, B, KAB := dhke(p, alpha, a, b)
	fmt.Printf("A: %d, B: %d, K_AB: %d\n", A, B, KAB)
}

func uc85() {
	fmt.Println("======= (8.5) =======")
	dhkePrint(467, 2, 3, 5)
	dhkePrint(467, 2, 400, 134)
	dhkePrint(467, 2, 228, 57)
}
