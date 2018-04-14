package main

import "fmt"

func shiftChar(s string, i int) string {
	result := ""
	for _, c := range s {
		r := c
		if 'a' <= c && c <= 'z' {
			r = rune((int(c)+i-int('a'))%26 + int('a'))
		}
		result += string(r)
	}
	return result
}

func tryKey(cipher string) {
	for i := 0; i < 26; i++ {
		if i == 11 {
			fmt.Println(i, shiftChar(cipher, i))
		}
	}
}

func uc12() {
	const cipher = "xultpaajcxitltlxaarpjhtiwtgxktghidhipxciwtvgtpilpitghlxiwiwtxgqadds."
	tryKey(cipher)
}
