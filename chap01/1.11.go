package main

import "fmt"

func inverse(a int) int {
	for i := 0; i < 26; i++ {
		if (a*i)%26 == 1 {
			return i
		}
	}
	panic("Inverse not found!")
}

func affine(c, a, b int) string {
	cb := c - b
	if cb < 0 {
		cb += 26
	}
	r := (inverse(a) * cb) % 26
	return string('a' + r)
}

func decrypt(s string, a, b int) string {
	plain := ""
	for _, c := range s {
		plain += affine(int(c-'a'), a, b)
	}
	return plain
}

func uc111() {
	const cipher = "falszztysyjzyjkywjrztyjztyynaryjkyswarztyegyyj"
	fmt.Print(decrypt(cipher, 7, 22))
}
