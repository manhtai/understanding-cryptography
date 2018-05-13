package main

import (
	"fmt"

	"github.com/manhtai/understanding-cryptography/pkg"
)

// Elgamal represent Elgamal scheme
type Elgamal struct {
	p           int
	alpha, beta int
	d, i        int
}

// NewElgamal initialize a new Elgamal instance with p & alpha
func NewElgamal(p, alpha int) (elg *Elgamal) {
	elg = &Elgamal{p: p, alpha: alpha}
	return
}

// SetPrivate set private key d for Elgamal
func (elg *Elgamal) SetPrivate(d int) {
	elg.d = d
	elg.beta = pkg.Snm(elg.alpha, d, elg.p)
}

// Encrypt encrypts message x using random i
func (elg *Elgamal) Encrypt(x, i int) (y, ke int) {
	ke = pkg.Snm(elg.alpha, i, elg.p)
	km := pkg.Snm(elg.beta, i, elg.p)
	y = (x * km) % elg.p
	return
}

// Decrypt decrypts cipher text using private key
func (elg *Elgamal) Decrypt(y, ke int) (x int) {
	x = (y * pkg.Snm(ke, elg.p-1-elg.d, elg.p)) % elg.p
	return
}

func uc813() {
	fmt.Println("======= (8.13) =======")

	elg := NewElgamal(467, 2)

	elg.SetPrivate(105)
	x1 := 33
	y1, ke1 := elg.Encrypt(x1, 213)
	y2, ke2 := elg.Encrypt(x1, 123)
	fmt.Println(ke1, y1)
	fmt.Println(ke2, y2)

	fmt.Println(elg.Decrypt(y1, ke1), "=", x1)
	fmt.Println(elg.Decrypt(y2, ke2), "=", x1)

	elg.SetPrivate(300)
	x2 := 248
	y3, ke3 := elg.Encrypt(x2, 45)
	y4, ke4 := elg.Encrypt(x2, 47)
	fmt.Println(ke3, y3)
	fmt.Println(ke4, y4)

	fmt.Println(elg.Decrypt(y3, ke3), "=", x2)
	fmt.Println(elg.Decrypt(y4, ke4), "=", x2)
}
