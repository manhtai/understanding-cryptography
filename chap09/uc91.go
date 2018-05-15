package main

import "fmt"

func uc91() {
	fmt.Println("======= (9.1) =======")
	ec := EC{a: 2, b: 2, p: 17}
	ec.assert()
}
