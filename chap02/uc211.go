package main

import "fmt"
import "strings"
import "strconv"

func buildKeyDict() map[string]int64 {
	keyDict := map[string]int64{}
	for i := 0; i < 26; i++ {
		keyDict[string('A'+i)] = int64(i)
	}
	for i := 0; i < 6; i++ {
		k := fmt.Sprintf("%d", i)
		keyDict[k] = int64(26 + i)
	}
	return keyDict
}

func stringToBinary(s string) string {
	r := ""
	d := buildKeyDict()
	for _, c := range s {
		k := strings.ToUpper(string(c))
		r += fmt.Sprintf("%05b", d[k])
	}
	return r
}

func xor(s1, s2 string) string {
	if len(s1) != len(s2) {
		panic("s1, s2 must be same length" + s1 + "+" + s2)
	}
	r := ""
	for i, c := range s1 {
		c1, _ := strconv.ParseInt(string(c), 2, 64)
		c2, _ := strconv.ParseInt(string(s2[i]), 2, 64)
		r += fmt.Sprintf("%b", (c1+c2)%2)
	}
	return r
}

func lfsr(iv, co string, length int) string {
	if len(iv) != len(co) {
		panic("initial values & feedback coefficients must be same length!")
	}

	loop := length - len(iv)
	for i := 0; i < loop; i++ {
		var m int64
		for j := 0; j < len(co); j++ {
			ivij, _ := strconv.ParseInt(string(iv[i+j]), 2, 64)
			coj, _ := strconv.ParseInt(string(co[j]), 2, 64)
			m += ivij * coj
		}
		iv += fmt.Sprintf("%b", m%2)
	}
	return iv
}

func binaryToString(b string) string {
	d := buildKeyDict()
	m := map[int64]string{}
	for k, v := range d {
		m[v] = k
	}
	n, _ := strconv.ParseInt(b, 2, 64)
	return m[n]
}

func decrypt(c string) string {
	s := ""
	for i := 0; i < len(c); i += 5 {
		s += binaryToString(c[i : i+5])
	}
	return s
}

func uc211() {
	cipher := "j5a0edj2b"
	p03 := stringToBinary("WPI")
	c03 := stringToBinary(cipher[0:3])
	fmt.Println(p03)
	fmt.Println(c03)
	fmt.Println(xor(p03, c03))

	// After solving equation
	// TODO: Use Go to solve that
	iv := xor(p03, c03)[0:6]
	f := "110000"
	c := stringToBinary(cipher)
	s := lfsr(iv, f, len(c))
	p := xor(s, c)
	plain := decrypt(p)
	fmt.Println(plain)
}
