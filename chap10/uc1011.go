package main

import (
	"fmt"

	"github.com/manhtai/understanding-cryptography/pkg"
)

type elgamalDS struct {
	d              int
	p, alpha, beta int
}

func (e *elgamalDS) verify(x, r, s int) (t bool) {
	t1 := (pkg.Snm(e.beta, r, e.p) * pkg.Snm(r, s, e.p)) % e.p
	t2 := pkg.Snm(e.alpha, x, e.p)
	if t1 == t2 {
		t = true
	}
	return
}

func (e *elgamalDS) sign(x, kE int) (r, s int) {
	pMinus1 := e.p - 1

	r = pkg.Snm(e.alpha, kE, e.p)
	s = ((x - e.d*r) * pkg.Inverse(kE, pMinus1)) % pMinus1

	if s <= 0 {
		s += pMinus1
	}
	return
}

func (e *elgamalDS) getBeta() (beta int) {
	beta = pkg.Snm(e.alpha, e.d, e.p)
	return
}

func uc1011() {
	fmt.Println("======= (10.11) ======")
	e := elgamalDS{p: 31, alpha: 3, beta: 6}
	x := 10
	fmt.Println(e.verify(x, 17, 5))
	fmt.Println(e.verify(x, 13, 15))

	for d := 2; d < e.p-1; d++ {
		e.d = d
		if e.getBeta() != e.beta {
			continue
		}

		for kE := 2; kE < e.p-1; kE++ {
			_, _, g := pkg.Gcde(e.p-1, kE)
			if g != 1 {
				continue
			}

			r, s := e.sign(x, kE)
			if !e.verify(x, r, s) {
				panic("Not valid!")
			}
			fmt.Println(r, s)
		}
	}
}
