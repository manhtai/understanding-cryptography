package main

import "fmt"

func uc21() {
	cipher := "bsaspp kkuosp"
	key := "rsidpy dkawoa"
	text := ""
	for i, c := range cipher {
		x := int(c - 'a')
		s := int(key[i] - 'a')
		y := (x - s) % 26
		if y < 0 {
			y += 26
		}
		text += string('a' + y)
	}
	fmt.Println(text)
}
