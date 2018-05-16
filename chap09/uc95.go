package main

import (
	"fmt"

	"github.com/manhtai/understanding-cryptography/pkg"
)

// EC represents Elliptic Curves function y^2 = x^3 + ax^2 + b mod p
type EC struct {
	a, b, p int
}

func (ec *EC) assert() {
	if (4*ec.a*ec.a*ec.a+27*ec.b*ec.b)%ec.p == 0 {
		panic("4a^3 + 27b^2 = 0")
	}
	fmt.Println("4a^3 + 27b^2 != 0")
}

func modPositive(x, p int) (y int) {
	y = x % p
	if y < 0 {
		y += p
	}
	return y
}

func (ec *EC) add(x1, y1, x2, y2 int) (x3, y3 int) {
	s := modPositive((y2-y1)*pkg.Inverse(x2-x1, ec.p), ec.p)
	x3 = modPositive(s*s-x1-x2, ec.p)
	y3 = modPositive(s*(x1-x3)-y1, ec.p)
	return
}

func (ec *EC) double(x1, y1 int) (x3, y3 int) {
	s := modPositive((3*x1*x1+ec.a)*pkg.Inverse(2*y1, ec.p), ec.p)
	x3 = modPositive(s*s-2*x1, ec.p)
	y3 = modPositive(s*(x1-x3)-y1, ec.p)
	return
}

func uc95() {
	fmt.Println("======= (9.5) =======")
	ec := EC{a: 3, b: 2, p: 7}
	x1, y1 := 0, 3

	x2, y2 := ec.double(x1, y1)
	fmt.Println(x2, y2)

	for {
		x2, y2 = ec.add(x2, y2, x1, y1)
		fmt.Println(x2, y2)

		// Neutral element
		if (x1 == x2) && (y1+y2)%ec.p == 0 {
			fmt.Println("ø")
			break
		}
	}
}
