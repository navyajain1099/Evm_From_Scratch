// Package evm is an **incomplete** implementation of the Ethereum Virtual
// Machine for the "EVM From Scratch" course:
// https://github.com/w1nt3r-eth/evm-from-scratch
//
// To work on EVM From Scratch In Go:
//
// - Install Golang: https://golang.org/doc/install
// - Go to the `go` directory: `cd go`
// - Edit `evm.go` (this file!), see TODO below
// - Run `go test ./...` to run the tests
package evm

import (
	"math/big"
)

// Run runs the EVM code and returns the stack and a success indicator.
func Evm(code []byte) ([]*big.Int, bool) {
	var stack []*big.Int

	for pc := 0; pc < len(code); pc++ {
		op := code[pc]

		if op >= 0x60 && op <= 0x7f {
			n := int(op - 0x5f)
			if pc+n >= len(code) {
				return stack, false // Not enough bytes for PUSHn
			}
			value := new(big.Int).SetBytes(code[pc+1 : pc+1+n])
			stack = append([]*big.Int{value}, stack...)
			pc += n // Increment pc to skip the next n bytes
		}

		switch op {
		case 0x00: // STOP
			return stack, true
		case 0x5f: //PUSH0
			stack = append([]*big.Int{big.NewInt(0)}, stack...)
			// case 0x60: //PUSH1

			// 	stack = append(stack, big.NewInt(int64(code[pc+1])))
			// 	pc++ // Increment pc to get the next byte

			// case 0x61: //PUSH2

			// 	if pc+2 >= len(code) {
			// 		return stack, false // Not enough bytes for PUSH2
			// 	}
			// 	value := (int64(code[pc+1]) << 8) | int64(code[pc+2])
			// 	stack = append(stack, big.NewInt(value))
			// 	pc += 2 // Increment pc to skip the next two bytes

			// case 0x63: //PUSH4
			// 	if pc+4 >= len(code) {
			// 		return stack, false // Not enough bytes for PUSH3
			// 	}
			// 	value := (int64(code[pc+1]) << 24) | (int64(code[pc+2]) << 16) | int64(code[pc+3])<<8 | int64(code[pc+4])
			// 	stack = append(stack, big.NewInt(value))
			// 	pc += 4 // Increment pc to skip the next four bytes

		}
	}
	return stack, true
}
