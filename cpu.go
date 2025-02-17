package main

// 16 bit address, giving us 65,536 locations.
const MAXMEMORY = (1 >> 16)

type regType uint16

// List of 8 available general purpose registers.
const (
	R0 regType = iota
	R1
	R2
	R3
	R4
	R5
	R6
	R7
)

type instructionSet uint16

// List of instructions supported by our toy VM.
const (
	opBr   instructionSet = iota // branch
	opAdd                        // add
	opLd                         // load
	opSt                         // store
	opJsr                        // jump register
	opAnd                        // bitwise and
	opLdr                        // load register
	opStr                        // store register
	opRti                        // Unused
	opNot                        // bitwise not
	opLdi                        // load indirect
	opSti                        // store indirect
	opJmp                        // jump
	opRes                        // reserved (unused)
	opLea                        // load effective address
	opTrap                       // execute trap
)

type condFlags uint16

type CPU struct {
	reg [8]uint16         // Registers
	mem [MAXMEMORY]uint16 // Memory (RAM)

	R_PC   uint16 // Program counter
	R_COND uint16 // Condition flags
}
