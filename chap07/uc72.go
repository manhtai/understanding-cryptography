package main

import (
	"fmt"

	"github.com/manhtai/understanding-cryptography/pkg"
)

// snmPrint compute snm & print result
func snmPrint(x, e, m int) {
	r := pkg.Snm(x, e, m)
	fmt.Printf("%d^%d mod %d = %d\n", x, e, m, r)
}

func uc72() {
	fmt.Println("======= (7.2) =======")
	snmPrint(2, 79, 101)
	snmPrint(3, 197, 101)
}
