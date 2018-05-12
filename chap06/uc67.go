package main

import "fmt"
import "github.com/manhtai/understanding-cryptography/pkg"

func inverse(a, m int) (i int) {
	_, i, _ = pkg.Gcde(m, a)
	if i < 0 {
		i += m
	}
	return
}

func inversePrint(a, m int) {
	i := inverse(a, m)
	fmt.Printf("(%d x %d) %% %d = %d\n", a, i, m, (a*i)%m)
}

func uc67() {
	inversePrint(7, 26)
	inversePrint(19, 999)
}
