package vm

func extendSign(val uint16, bitCount int) uint16 {
	if val>>(bitCount-1)&1 == 1 {
		val |= (0xFFFF << bitCount)
	}

	return val
}

// Check if val is pos, neg or zero after every write to a register.
// Set the condition flags accordingly.
func (v *VM) update_flags(address uint16) {
	if v.Reg[address] == 0 {
		v.R_COND = Z
	}

	// A 1 in the leftmost bit indicates a negative signed value
	if v.Reg[address]>>15 == 1 {
		v.R_COND = N
	} else {
		v.R_COND = P
	}
}
