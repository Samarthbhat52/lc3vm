package tests

import (
	"testing"

	"github.com/Samarthbhat52/lc3vm/vm"
)

func TestAddInstruction(t *testing.T) {
	v := vm.NewVm()
	v.Reg[vm.R1] = 10
	v.Reg[vm.R2] = 5

	instruction := uint16(0b0001011001000010)
	v.Mem[vm.PC_START] = instruction

	v.Execute()

	if v.Reg[3] != 15 {
		t.Errorf("Expected 15, got %d", v.Reg[3])
	}

	if v.R_COND != vm.P {
		t.Errorf("Expected COND = 0x1 (P), got %d", v.R_COND)
	}
}
