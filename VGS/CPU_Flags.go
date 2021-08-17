package VGS

import "fmt"

//-------------------------------------------------- Processor Flags --------------------------------------------------//

// ------------------------------ Zero Flag ------------------------------ //
func flags_Z(value byte) {
	if Debug {
		fmt.Printf("\tFlag Z: %d ->", P[1])
	}
	// Check if final value is 0
	if value == 0 {
		P[1] = 1
	} else {
		P[1] = 0
	}
	if Debug {
		fmt.Printf(" %d\n", P[1])
	}
}

// ---------------------------- Negative Flag ---------------------------- //
func flags_N(value byte) {
	if Debug {
		fmt.Printf("\tFlag N: %d ->", P[7])
	}

	// Set Negtive flag to the the MSB of the value
	P[7] = value >> 7

	if Debug {
		fmt.Printf(" %d | Value = %08b\n", P[7], value)
	}
}

func flags_SBC_DECIMAL(value int) {
	if Debug {
		fmt.Printf("\tFlag N: %d ->", P[7])
	}

	if value < 0x00 {
		P[7] = 1
	} else {
		P[7] = 0
	}

	if Debug {
		fmt.Printf(" %d | Value = %08b\n", P[7], value)
	}
}

// ----------------------------- Carry Flag ------------------------------ //

// Used by CPX, CPY, CMP
func flags_C_CPX_CPY_CMP(value1, value2 byte) {
	if Debug {
		fmt.Printf("\tFlag C: %d ->", P[0])
	}

	// Check if final value is 0
	if value1 >= value2 {
		P[0] = 1
	} else {
		P[0] = 0
	}

	if Debug {
		fmt.Printf(" %d\n", P[0])
	}
}

// Used by ASL
func flags_C(value byte) {
	if Debug {
		fmt.Printf("\tFlag C: %d ->", P[0])
	}

	P[0] = value

	if Debug {
		fmt.Printf(" %d\n", P[0])
	}
}

// Used by ADC and SBC
func flags_C_ADC_SBC(value_A, value_Mem, value_P0 byte) {

	var A_16bit uint16 // 16 bit variable to detect carry

	if Debug {
		fmt.Printf("\tFlag C: %d ->", P[0])
	}

	A_16bit = uint16(value_A) + uint16(value_Mem) + uint16(value_P0)

	// Set if overflow in bit 7 (the sum of values are smaller than original A)
	if A_16bit > 255 {
		P[0] = 1
	} else {
		P[0] = 0
	}

	if Debug {
		fmt.Printf(" %d\n", P[0])
	}
}

// Used by ADC in Decimal mode
func flags_C_ADC_DECIMAL(value int64) {

	if Debug {
		fmt.Printf("\tFlag C: %d ->", P[0])
	}

	// Update the carry flag value
	if value > 0x99 {
		P[0] = 1
	} else {
		P[0] = 0
	}

	if Debug {
		fmt.Printf(" %d\n", P[0])
	}
}

// Used by SBC in Decimal mode
func flags_C_SBC_DECIMAL(value int) {

	if Debug {
		fmt.Printf("\tFlag C: %d ->", P[0])
	}

	// Update the carry flag value
	if value >= 0x00 {
		P[0] = 1
	} else {
		P[0] = 0
	}

	if Debug {
		fmt.Printf(" %d\n", P[0])
	}
}

// ---------------------------- oVerflow Flag ---------------------------- //

// oVerflow Flag for ADC
// oVerflow Flag for SBC (receiving one's complement of Memory value)
func flags_V(value_A, value_Mem, value_P0 byte) {
	var (
		carry_bit [8]byte
		carry_OUT byte = 0
	)
	// fmt.Printf("\n  %08b\t%d",value1,value1)
	// fmt.Printf("\n  %08b\t%d",value2,value2)

	if Debug {
		fmt.Printf("\tFlag V: %d ->", P[6])
	}

	// Set the carry flag on bit 0 of carry_bit Array to bring the carry if exists
	carry_bit[0] = value_P0

	// Make the magic
	for i := 0; i <= 7; i++ {
		// sum the bit from value one + bit from value 2 + carry value
		tmp := (value_A >> byte(i) & 0x01) + (value_Mem >> byte(i) & 0x01) + carry_bit[i]
		if tmp >= 2 {
			// set the carry out
			if i == 7 {
				carry_OUT = 1
			} else {
				carry_bit[i+1] = 1
			}
		}
	}

	// Formula to calculate: V = C6 xor C7
	P[6] = carry_bit[7] ^ carry_OUT
	// fmt.Printf("\nV: %d", P[6])

	if Debug {
		fmt.Printf(" %d\n", P[6])
	}
}

// Used by ASL
func flags_V_BIT(value byte) {
	// Memory Address bit 6 -> V (oVerflow)
	if Debug {
		fmt.Printf("\tFlag V: %d -> ", P[6])
	}
	P[6] = value >> 6 & 0x01 // Keep only the 6th bit

	if Debug {
		fmt.Printf("%d\n", P[6])
	}
}

// --------------------------- IRQ Disable Flag -------------------------- //
func flags_I(value byte) {
	if Debug {
		fmt.Printf("\tFlag I: %d ->", P[2])
	}

	P[2] = value

	if Debug {
		fmt.Printf(" %d\n", P[2])
	}
}

// --------------------------- Break Flag Flag --------------------------- //
func flags_B(value byte) {
	if Debug {
		fmt.Printf("\tFlag B: %d ->", P[4])
	}

	P[4] = value

	if Debug {
		fmt.Printf(" %d\n", P[4])
	}
}
