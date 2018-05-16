package main

import "fmt"

func uc99() {
	fmt.Println("======= (9.9) =======")
	ec := EC{a: 1, b: 6, p: 11}
	a := 6
	xB, yB := 5, 9
	xAB, yAB := ec.dna(xB, yB, a)
	fmt.Println(xAB, yAB)
}
