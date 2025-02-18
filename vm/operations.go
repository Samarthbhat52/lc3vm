package vm

func (vm *VM) add(instr uint16) {
	// Get the register values.
	r0 := instr >> 9 & 0x7
	r1 := instr >> 6 & 0x7

	// Get the 'immediate mode' value(either a 0 or a 1)
	imm := instr >> 5 & 0x1

	if imm == 1 {
		// Get the value coded in the instruction and sign extend it to match our 16 bit registers.
		imm_addr := instr & 0x1F
		val2 := extendSign(imm_addr, 5)

		// Add the decoded value with the value held by r1
		vm.Reg[r0] = vm.Reg[r1] + val2
	} else {
		// Get the second operand
		r2 := instr & 0x7

		// add the values
		vm.Reg[r0] = vm.Reg[r1] + vm.Reg[r2]
	}
}
