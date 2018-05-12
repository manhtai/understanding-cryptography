package main

import "fmt"

func bitLen(e int) (l int) {
	i := e
	for i > 0 {
		i >>= 1
		l++
	}
	return
}

// snm compute x^e % n using square-and-multiply algorithm
func snm(x, e, m int) (r int) {
	r = x
	l := bitLen(e)
	for i := 1; i < l; i++ {
		r = (r * r) % m
		if (e>>uint(l-i-1))&1 == 1 {
			r = (r * x) % m
		}
	}
	return
}

// snmPrint compute snm & print result
func snmPrint(x, e, m int) {
	r := snm(x, e, m)
	fmt.Printf("%d^%d mod %d = %d\n", x, e, m, r)
}

func uc72() {
	fmt.Println("======= (7.2) =======")
	snmPrint(2, 79, 101)
	snmPrint(3, 197, 101)
}
