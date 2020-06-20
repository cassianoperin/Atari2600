package CPU

import	"fmt"

// BMI  Branch on Result Minus (relative)
//
//      branch on N = 1                  N Z C I D V
//                                       - - - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      relative      BMI oper      30    2     2**
func opc_BMI(offset uint16) {
	// If Negative
	if P[7] == 1 {

		if Debug {
			fmt.Printf("\n\tOpcode %02X%02X [2 bytes]\tBMI  Branch on Result Minus (relative).\tCarry EQUAL 1, JUMP TO %04X\n", Opcode, Memory[PC+1], PC+2+uint16(offset) )
		}
		// Current PC (To detect page bounday cross)
		tmp := PC
		// fmt.Printf("\ntmp: %02X\n",tmp)

		// PC + the number of bytes to jump on carry clear
		PC += uint16(offset)

		// Increment PC
		PC += 2

		// Add 1 to cycles if branch occurs on same page
		Beam_index += 1

		// // Add one extra cycle if branch occurs in a differente memory page
		if MemPageBoundary(uint16(tmp), PC) {
			Beam_index += 1
		}

	// If not negative
	} else {
		if Debug {
			fmt.Printf("\n\tOpcode %02X%02X [2 bytes]\tBMI  Branch on Result Minus (relative).\t\tNEGATIVE Flag DISABLED, PC+=2\n", Opcode, Memory[PC+1])
		}
		PC += 2
	}

	Beam_index += 2
}
