package vm

const (
	// 16 bit address, giving us 65,536 locations.
	MAXMEMORY = (1 << 16)
	// Program counter starts at address 0x3000 on lc3 by convention.
	PC_START uint16 = 0x3000
)

// List of 8 available general purpose registers.
const (
	R0 uint16 = iota
	R1
	R2
	R3
	R4
	R5
	R6
	R7
)

// List of instructions supported by our toy VM.
const (
	OpBr   uint16 = iota // branch
	OpAdd                // add
	OpLd                 // load
	OpSt                 // store
	OpJsr                // jump register
	OpAnd                // bitwise and
	OpLdr                // load register
	OpStr                // store register
	OpRti                // Unused
	OpNot                // bitwise not
	OpLdi                // load indirect
	OpSti                // store indirect
	OpJmp                // jump
	OpRes                // reserved (unused)
	OpLea                // load effective address
	OpTrap               // execute trap
)

type condFlags uint16

// Condition flags that our cpu accepts.
const (
	P condFlags = 1 << 0 // Positive flag
	Z condFlags = 1 << 1 // Zero flag
	N condFlags = 1 << 2 // Negative flag
)

type VM struct {
	Reg [8]uint16         // Registers
	Mem [MAXMEMORY]uint16 // Memory (RAM)

	R_PC   uint16    // Program counter
	R_COND condFlags // Condition flags
}
