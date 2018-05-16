package main

import (
	"fmt"

	"github.com/manhtai/understanding-cryptography/pkg"
)

func (ec *EC) dna(x1, y1, k int) (x2, y2 int) {
	x2, y2 = x1, y1
	l := pkg.BitLen(k)
	for i := 1; i < l; i++ {
		x2, y2 = ec.double(x2, y2)
		if (k>>uint(l-i-1))&1 == 1 {
			x2, y2 = ec.add(x2, y2, x1, y1)
		}
	}
	return
}

func uc97() {
	fmt.Println("======= (9.7) =======")
	ec := EC{a: 4, b: 20, p: 29}
	x1, y1 := 8, 10

	fmt.Println(ec.dna(x1, y1, 9))
	fmt.Println(ec.dna(x1, y1, 20))
}
