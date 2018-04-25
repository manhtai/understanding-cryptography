package main

import (
	"encoding/hex"
	"fmt"
	"log"
)

func uc313() {
	plainText, _ := hex.DecodeString("0000000000000000")
	key, _ := hex.DecodeString("BBBB55555555EEEEFFFF")
	present, err := NewCipher(key)

	if err != nil {
		log.Fatal(err)
	}

	var cipherText, plainText2 []byte
	cipherText = make([]byte, 8)
	present.Encrypt(cipherText, plainText)
	fmt.Printf("%X\n", cipherText)

	plainText2 = make([]byte, 8)
	present.Decrypt(plainText2, cipherText)
	fmt.Printf("%X\n%X\n", plainText, plainText2)
}
