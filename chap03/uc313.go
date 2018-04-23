package main

import (
	"encoding/hex"
	"log"
)

func uc313() {
	plainText, _ := hex.DecodeString("0000000000000000")
	key, _ := hex.DecodeString("BBBB55555555EEEEFFFF")
	present, err := NewCipher(key)

	if err != nil {
		log.Fatal(err)
	}

	var cipherText []byte
	cipherText = make([]byte, 16)
	present.Encrypt(cipherText, plainText)
}
