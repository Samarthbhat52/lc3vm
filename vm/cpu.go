package vm

import "fmt"

func NewVm() *VM {
	return &VM{}
}

func (c *VM) Execute() {
	vm := NewVm()

	// Init the condition flag to zero (Z) condition.
	vm.R_COND = Z
	vm.R_PC = PC_START

	running := true

	for i := 0; ; i++ {
		instruction := vm.Mem[vm.R_PC]
		instruction++
		// Get the first 4 bits as the op code by bit shifting to the right by 12 bits.
		opCode := instruction >> 12

		switch opCode {
		case OpAdd:
			fmt.Println(opCode)
		default:
			return
		}

		if !running {
			break
		}

	}
}
