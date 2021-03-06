package main

import (
	"crypto/cipher"
	"encoding/binary"
	// "fmt"
	"strconv"
)

// BlockSize of PRESENT is 64 bits
const BlockSize = 8

// KeySizeError is used to raise invalid key length, i.e. not 80 bits length
type KeySizeError int

func (k KeySizeError) Error() string {
	return "present80: invalid key size " + strconv.Itoa(int(k))
}

type presentCipher struct {
	subkeys [32]uint64
}

// NewCipher creates and returns a new cipher.Block.
func NewCipher(key []byte) (cipher.Block, error) {
	if len(key) != 10 {
		return nil, KeySizeError(len(key))
	}

	c := new(presentCipher)
	c.generateSubkeys(key)

	return c, nil
}

func (c *presentCipher) BlockSize() int { return BlockSize }

func (c *presentCipher) Encrypt(dst, src []byte) { encryptBlock(c.subkeys[:], dst, src) }

func (c *presentCipher) Decrypt(dst, src []byte) { decryptBlock(c.subkeys[:], dst, src) }

// creates 31 80-bit subkeys from the original key
func (c *presentCipher) generateSubkeys(keyBytes []byte) {
	// We seperate the 80 bits key: 80 bits = 64 bits + 16 bits
	key1 := binary.BigEndian.Uint64(keyBytes[0:8])
	key2 := binary.BigEndian.Uint16(keyBytes[8:10])

	key1RotateMask := (^uint64(0) >> 61)
	key1SBoxMask := (^uint64(0) >> 4)
	key2RotateMask := ((^uint64(0) >> 47) << 3)
	counterMask := ^uint64(0) >> 63

	c.subkeys[0] = key1
	// fmt.Printf("Initial key:         %08X %02X\n", key1, key2)

	// generate subkeys
	for i := 1; i < 32; i++ {
		// Rotate left 61
		key1 = ((key1 & key1RotateMask) << 61) | (uint64(key2) << 45) | (key1 >> 19)
		key2 = uint16((c.subkeys[i-1] & key2RotateMask) >> 3)
		// fmt.Printf("Key after rotation %2d: %08X %02X\n", i, key1, key2)

		// Left most bit passed to sBox
		s4 := runSBoxLayer(key1>>60, sBox[:])
		key1 = (key1 & key1SBoxMask) | (uint64(s4) << 56)
		// fmt.Printf("Key after S-box    %2d: %08X %02X\n", i, key1, key2)

		// XOR to rount_counter
		key1 ^= (uint64(i) >> 1)
		key2 ^= uint16(uint64(i) & counterMask)
		// fmt.Printf("Key after XOR      %2d: %08X %02X\n", i, key1, key2)

		c.subkeys[i] = key1
	}
}

func cryptBlock(subkeys []uint64, dst, src []byte, decrypt bool) {
	output := binary.BigEndian.Uint64(src)
	if decrypt {
		output ^= subkeys[31]
		for i := 30; i >= 0; i-- {
			output = sp(output, subkeys[i], true)
		}
	} else {
		for i := 0; i < 31; i++ {
			output = sp(output, subkeys[i], false)
		}
		output ^= subkeys[31]
	}
	binary.BigEndian.PutUint64(dst, output)
}

// Encrypt one block from src into dst, using the subkeys.
func encryptBlock(subkeys []uint64, dst, src []byte) {
	cryptBlock(subkeys, dst, src, false)
}

// Decrypt one block from src into dst, using the subkeys.
func decryptBlock(subkeys []uint64, dst, src []byte) {
	cryptBlock(subkeys, dst, src, true)
}

// Run a substitution-permutation network block
func sp(input, key uint64, decrypt bool) uint64 {
	if decrypt {
		input = runPLayer(input, reverseBox(pBox[:]))
		input = runSBoxLayer(input, reverseBox(sBox[:]))
		input ^= key
	} else {
		input ^= key
		input = runSBoxLayer(input, sBox[:])
		input = runPLayer(input, pBox[:])
	}
	return input
}

var sBox = [16]uint8{12, 5, 6, 11, 9, 0, 10, 13, 3, 14, 15, 8, 4, 7, 1, 2}
var pBox = [64]uint8{
	0, 16, 32, 48, 1, 17, 33, 49, 2, 18, 34, 50, 3, 19, 35, 51,
	4, 20, 36, 52, 5, 21, 37, 53, 6, 22, 38, 54, 7, 23, 39, 55,
	8, 24, 40, 56, 9, 25, 41, 57, 10, 26, 42, 58, 11, 27, 43, 59,
	12, 28, 44, 60, 13, 29, 45, 61, 14, 30, 46, 62, 15, 31, 47, 63,
}

func runSBoxLayer(input uint64, sBox []uint8) (output uint64) {
	var preOutput uint8
	for i := 0; i < 16; i++ {
		mask := (^uint64(0) >> uint(4*i)) & (^uint64(0) << (4 * uint(15-i)))
		preOutput = uint8((input & mask) >> (4 * uint(15-i)))
		output |= uint64(sBox[preOutput]) << (4 * uint64(15-i))
	}
	return
}

func runPLayer(src uint64, permutation []uint8) (block uint64) {
	for n, position := range permutation {
		bit := (src >> uint(n)) & 1
		block |= bit << uint(len(permutation)-1-int(position))
	}
	return
}

func reverseBox(box []uint8) (reverse []uint8) {
	reverse = make([]uint8, len(box))
	for n, position := range box {
		reverse[position] = uint8(n)
	}
	return
}
