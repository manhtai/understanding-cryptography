package main

import (
	"crypto/cipher"
	"encoding/binary"
	"fmt"
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

func (c *presentCipher) Decrypt(dst, src []byte) { panic("To be implemented!") }

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
	fmt.Printf("Initial key: %08X %02X\n", key1, key2)

	// generate subkeys
	for i := 1; i < 32; i++ {
		// Rotate left 61
		key1 = ((key1 & key1RotateMask) << 61) | (uint64(key2) << 45) | (key1 >> 19)
		key2 = uint16((c.subkeys[i-1] & key2RotateMask) >> 3)
		// fmt.Printf("Key after rotation %2d: %08X %02X\n", i, key1, key2)

		// Left most bit passed to sBox
		s4 := runSBoxLayer(uint8(key1 >> 60))
		key1 = (key1 & key1SBoxMask) | (uint64(s4) << 56)
		// fmt.Printf("Key after S-box    %2d: %08X %02X\n", i, key1, key2)

		// XOR to rount_counter
		key1 ^= (uint64(i) >> 1)
		key2 ^= uint16(uint64(i) & counterMask)
		// fmt.Printf("Key after XOR      %2d: %08X %02X\n", i, key1, key2)

		c.subkeys[i] = key1
	}
}

func cryptBlock(subkeys []uint64, dst, src []byte) {
	output := binary.BigEndian.Uint64(src)
	for i := 0; i < 31; i++ {
		output = sp(output, subkeys[i])
	}
	output ^= subkeys[31]
	binary.BigEndian.PutUint64(dst, output)
}

// Encrypt one block from src into dst, using the subkeys.
func encryptBlock(subkeys []uint64, dst, src []byte) {
	cryptBlock(subkeys, dst, src)
}

// Run a substitution-permutation network block
func sp(input, key uint64) uint64 {
	input ^= key
	fmt.Printf("State after Add:     %8X\n", input)
	input = runSBoxLayers(input)
	fmt.Printf("State after S-Box:   %8X\n", input)
	input = runPLayers(input)
	fmt.Printf("State after P-Layer: %8X\n", input)
	return input
}

func runPLayers(input uint64) uint64 {
	var preOutput uint8
	var output uint64
	for i := 0; i < 16; i++ {
		preOutput = uint8((input << uint64(4*i)) >> (4 * uint64(15+i)))
		output |= uint64(runPLayer(preOutput)) << (4 * uint64(15-i))
	}
	return output
}

func runSBoxLayers(input uint64) uint64 {
	var preOutput uint8
	var output uint64
	for i := 0; i < 16; i++ {
		preOutput = uint8((input << uint64(4*i)) >> (4 * uint64(15+i)))
		output |= uint64(runSBoxLayer(preOutput)) << (4 * uint64(15-i))
	}
	return output
}

func runPLayer(i uint8) uint8 {
	if i >= 0 && i < 63 {
		return i * 16 % 63
	}
	return 63
}

var sBox = [16]uint8{12, 5, 6, 11, 9, 0, 10, 13, 3, 14, 15, 8, 4, 7, 1, 2}

func runSBoxLayer(input uint8) uint8 {
	return sBox[input]
}
