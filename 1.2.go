package uc1_2

import "fmt"

const cipher = "xultpaajcxitltlxaarpjhtiwtgxktghidhipxciwtvgtpilpitghlxiwiwtxgqadds."

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

func tryKey() {
	for i := 0; i < 26; i++ {
		if i == 11 {
			fmt.Println(i, shiftChar(cipher, i))
		}
	}
}

func main() {
	tryKey()
}
