package main

import "fmt"
import "github.com/manhtai/understanding-cryptography/pkg"

func inversePrint(a, m int) {
	i := pkg.Inverse(a, m)
	fmt.Printf("(%d x %d) %% %d = %d\n", a, i, m, (a*i)%m)
}

func uc67() {
	inversePrint(7, 26)
	inversePrint(19, 999)
}
