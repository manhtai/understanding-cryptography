package main

import (
	"fmt"

	"github.com/manhtai/understanding-cryptography/pkg"
)

func verifyRSADS(n, e, x, s int) (t bool) {
	xPrime := pkg.Snm(s, e, n)
	if xPrime == x {
		t = true
	}
	return
}

func uc105() {
	fmt.Println("======= (10.5) =======")
	fmt.Println(verifyRSADS(9797, 131, 123, 6292))
	fmt.Println(verifyRSADS(9797, 131, 4333, 4768))
	fmt.Println(verifyRSADS(9797, 131, 4333, 1424))
}
