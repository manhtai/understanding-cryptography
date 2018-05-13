package main

import (
	"fmt"

	"github.com/manhtai/understanding-cryptography/pkg"
)

// RSA represent the "book" RSA algorithm
type RSA struct {
	p, q, e, n, phi, d int
}

// NewRSA initialize an RSA cipher using p, q & e
func NewRSA(p, q, e int) *RSA {
	r := RSA{p: p, q: q}
	r.p, r.q, r.e = p, q, e
	r.phi = (p - 1) * (q - 1)
	r.n = p * q
	_, r.d, _ = pkg.Gcde(r.phi, e)
	return &r
}

// Pub return public key pair (n, e)
func (r *RSA) Pub() (int, int) {
	return r.n, r.e
}

// Prv return private key d
func (r *RSA) Prv() int {
	return r.d
}

// Encrypt encrypts plain text x using snm
func (r *RSA) Encrypt(x int) int {
	return pkg.Snm(x, r.e, r.n)
}

// Decrypt decrypts cipher text y using square-and-multiply algorithm
func (r *RSA) Decrypt(y int) int {
	return pkg.Snm(y, r.d, r.n)
}

// DecryptCRT decrypts cipher text y using Chinese Remainder Theorem
func (r *RSA) DecryptCRT(y int) int {
	dp := r.d % (r.p - 1)
	dq := r.d % (r.q - 1)
	xp := pkg.Snm(y, dp, r.p)
	xq := pkg.Snm(y, dq, r.q)
	_, cp, _ := pkg.Gcde(r.p, r.q)
	_, cq, _ := pkg.Gcde(r.q, r.p)
	return (r.q*cp*xp + r.p*cq*xq) % r.n
}

func uc77() {
	fmt.Println("======= (7.7) =======")
	rsa := NewRSA(31, 37, 17)
	y := 2
	x := rsa.Decrypt(y)
	x2 := rsa.DecryptCRT(y)
	fmt.Println("x, x2 =", x, x2)
}
