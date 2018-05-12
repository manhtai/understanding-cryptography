package main

import (
	"fmt"

	"github.com/manhtai/understanding-cryptography/pkg"
)

func uc71() {
	pkg.GcdePrint(32, 640)
	pkg.GcdePrint(49, 640)
	p := 41
	q := 17
	e := 49
	_, d, _ := pkg.Gcde((p-1)*(q-1), e)
	fmt.Println("(p, q, d) = ", p, q, d)
}
