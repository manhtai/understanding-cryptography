package main

import "fmt"

// gcd compute gcd of two numbers using Euclidean Algorithm
func gcd(m, n int) int {
	if m < n {
		return gcd(n, m)
	}
	if n == 0 {
		return m
	}
	return gcd(m%n, n)
}

func gcdPrint(m, n int) {
	fmt.Printf("gcd(%d, %d) = %d\n", m, n, gcd(m, n))
}

func uc65() {
	gcdPrint(7469, 2464)
	gcdPrint(2689, 4001)
}
