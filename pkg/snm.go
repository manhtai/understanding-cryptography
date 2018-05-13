package pkg

// BitLen calculates bit length of an integer
func BitLen(e int) (l int) {
	i := e
	for i > 0 {
		i >>= 1
		l++
	}
	return
}

// Snm compute x^e % n using square-and-multiply algorithm
func Snm(x, e, m int) (r int) {
	r = x
	l := BitLen(e)
	for i := 1; i < l; i++ {
		r = (r * r) % m
		if (e>>uint(l-i-1))&1 == 1 {
			r = (r * x) % m
		}
	}
	return
}
