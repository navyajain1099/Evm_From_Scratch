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
		case 0x50: // POP
			if len(stack) == 0 {
				return stack, false // Stack underflow
			}
			stack = stack[1:] // Remove the top element from the stack
		case 0x01: // ADD
			if len(stack) < 2 {
				return stack, false // Stack underflow
			}
			value1 := stack[0]
			value2 := stack[1]
			stack = stack[2:]                              // Remove the top two elements from the stack
			sumValue := new(big.Int).Add(value1, value2)   // Add the two values
			stack = append([]*big.Int{sumValue}, stack...) // Push the sum back onto the stack
		case 0x02: // MUL
			if len(stack) < 2 {
				return stack, false // Stack underflow
			}
			value1 := stack[0]
			value2 := stack[1]
			stack = stack[2:]
			productValue := new(big.Int).Mul(value1, value2)   // Multiply the two values
			stack = append([]*big.Int{productValue}, stack...) // Push the product back onto the stack
		case 0x03: // SUB
			if len(stack) < 2 {
				return stack, false // Stack underflow
			}
			value1 := stack[0]
			value2 := stack[1]
			stack = stack[2:]
			subtractValue := new(big.Int).Sub(value1, value2)   // Subtract the two values
			stack = append([]*big.Int{subtractValue}, stack...) // Push the difference back onto the stack
		case 0x04: // DIV
			if len(stack) < 2 {
				return stack, false // Stack underflow
			}

			if stack[1].Sign() == 0 {
				return stack, false // Division by zero
			}
			value1 := stack[0]
			value2 := stack[1]
			stack = stack[2:]
			divValue := new(big.Int).Div(value1, value2)   // Divide the two values
			stack = append([]*big.Int{divValue}, stack...) // Push the quotient back onto the stack

		}

	}
	return stack, true
}
