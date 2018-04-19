package main

import "fmt"

// Register is interface for LFSR and Trivium register
type Register interface {
	getOutputBit()
}

// LFSR represents an LFSR element in our Trivium
type LFSR struct {
	registers   []int64
	feedbackBit int
	forwardBit  int
}

// Trivium combines 3 LRFSs into 1
type Trivium struct {
	A, B, C *LFSR
}

func (l *LFSR) getOutputBit() int64 {
	return (l.registers[len(l.registers)-1] + l.forwardOutput() + l.andOutput()) % 2
}

func (l *LFSR) feedInputBit(input int64) {
	l.registers = append([]int64{input}, l.registers...)
	l.registers = l.registers[:len(l.registers)-1]
}

func (l *LFSR) forwardOutput() int64 {
	return l.registers[l.forwardBit]
}

func (l *LFSR) feedbackOutput() int64 {
	return l.registers[l.feedbackBit]
}

func (l *LFSR) andOutput() int64 {
	return l.registers[len(l.registers)-3] + l.registers[len(l.registers)-2]
}

func (t *Trivium) getOutputBit() int64 {
	return (t.A.getOutputBit() + t.B.getOutputBit() + t.C.getOutputBit()) % 2
}

func (t *Trivium) feedInputBits() {
	aOutput := t.A.getOutputBit()
	bOutput := t.B.getOutputBit()
	cOutput := t.C.getOutputBit()
	t.A.feedInputBit(cOutput)
	t.B.feedInputBit(aOutput)
	t.C.feedInputBit(bOutput)
}

func uc212() {
	cRegisters := make([]int64, 111)
	cRegisters[len(cRegisters)-1] = 1
	cRegisters[len(cRegisters)-2] = 1
	cRegisters[len(cRegisters)-3] = 1

	t := Trivium{
		A: &LFSR{
			registers:   make([]int64, 93),
			feedbackBit: 68,
			forwardBit:  65,
		},
		B: &LFSR{
			registers:   make([]int64, 84),
			feedbackBit: 77,
			forwardBit:  68,
		},
		C: &LFSR{
			registers:   cRegisters,
			feedbackBit: 86,
			forwardBit:  65,
		},
	}
	output := ""
	for i := 0; i < 70; i++ {
		o := t.getOutputBit()
		output += fmt.Sprintf("%b", o)
		// fmt.Printf("A: %+v\nB: %+v\nC: %+v\n===\n", *t.A, *t.B, *t.C)
		t.feedInputBits()
	}
	fmt.Println(output)
}
