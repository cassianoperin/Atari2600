package VGS

import	"fmt"

//-------------------------------------------------- Processor Flags --------------------------------------------------//

// Zero Flag
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

// Negative Flag
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

// Carry Flag
func flags_C(value1, value2 byte) {
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

// Carry Flag for Subtractions (SBC and CMP)
func flags_C_Subtraction(originalValue, newValue byte) {
	if Debug {
		fmt.Printf("\tFlag C: %d ->", P[0])
	}

	// If the new value is bigger than the original clear the flag
	// if originalValue < newValue {
	if newValue > originalValue {
		P[0] = 0
	} else {
		P[0] = 1
	}

	if Debug {
		fmt.Printf(" %d (SBC)\n" , P[0])
	}
}

// oVerflow Flag for ADC
// value1 = Accumulator, value2 = memory
func Flags_V_ADC(value1, value2 byte) {
	var (
		carry_bit		[8]byte
		carry_OUT 	byte = 0
	)
	// fmt.Printf("\n  %08b\t%d",value1,value1)
	// fmt.Printf("\n  %08b\t%d",value2,value2)

	if Debug {
		fmt.Printf("\tFlag V: %d ->", P[6])
	}

	// Set the carry flag on bit 0 of carry_bit Array to bring the carry if exists
	carry_bit[0] = P[0]

	// Make the magic
	for i:=0 ; i <= 7 ; i++{
		// sum the bit from value one + bit from value 2 + carry value
		tmp := (value1 >> byte(i) & 0x01) + (value2 >> byte(i) & 0x01) + carry_bit[i]
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

// oVerflow Flag for SBC
func Flags_V_SBC(value1, value2 byte) {
	var (
		carry_bit		= [8]byte{}
		carry_OUT 	byte = 0
	)

	// fmt.Printf("\n  \t %d\t(carry IN)",P[0])
	// fmt.Printf("\n  %08b\tDecimal: %d",value1,value1)
	// Since internall subtraction is just addition of the ones-complement
	// N can simply be replaced by 255-N in the formulas of Flags_V_ADC
	value2 = 255 - value2
	// fmt.Printf("\n  %08b\tDecimal: %d",value2,value2)

	// Set the carry flag on bit 0 of carry_bit Array to bring the carry if exists
	carry_bit[0] = P[0]

	if Debug {
		fmt.Printf("\tFlag V: %d ->", P[6])
	}

	// Set the carry flag on bit 0 of carry_bit Array to bring the carry if exists
	carry_bit[0] = P[0]

	// Make the magic
	for i:=0 ; i <= 7 ; i++{
		// sum the bit from value one + bit from value 2 + carry value
		tmp := (value1 >> byte(i) & 0x01) + (value2 >> byte(i) & 0x01) + carry_bit[i]
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
