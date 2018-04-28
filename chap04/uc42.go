package main

import "fmt"

type ops func(int, int) int

func gf(p int, o ops) (m [][]int) {
	// Init slice
	m = make([][]int, p)
	for i := 0; i < p; i++ {
		m[i] = make([]int, p)
	}

	// Calculate
	for i := 0; i < p; i++ {
		for j := i; j < p; j++ {
			m[i][j] = o(i, j)
			m[j][i] = m[i][j]
		}
	}

	return
}

func printGf(gf [][]int) {
	for i := range gf {
		fmt.Println(gf[i])
	}
	fmt.Println()
}

func uc42() {
	p := 7
	printGf(gf(p, func(a, b int) int { return (a + b) % p }))
	printGf(gf(p, func(a, b int) int { return (a * b) % p }))
}
