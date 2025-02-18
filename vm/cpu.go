package vm

import "fmt"

func NewVm() *VM {
	return &VM{}
}

func (vm *VM) Execute() {
	// Init the condition flag to zero (Z) condition.
	vm.R_COND = Z
	vm.R_PC = PC_START

	running := true

	for i := 0; ; i++ {

		if !running {
			break
		}

		instruction := vm.Mem[vm.R_PC]

		if instruction == 0 {
			fmt.Println("Execution completed.")
			break
		}

		vm.R_PC++
		// Get the first 4 bits as the op code by bit shifting to the right by 12 bits.
		opCode := instruction >> 12

		switch opCode {
		case OpAdd:
			vm.add(instruction)
			fmt.Println("Added value: ", vm.Reg[R3])
		default:
			fmt.Println("Wrong optype manh, sorry")
			running = false
		}

	}
}
