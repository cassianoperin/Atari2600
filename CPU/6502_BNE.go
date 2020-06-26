package CPU

import	"fmt"

// BNE  Branch on Result not Zero
//
//      branch on Z = 0                  N Z C I D V
//                                       - - - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      relative      BNE oper      D0    2     2**
func opc_BNE(value int8) {	// Receive a SIGNED value
	// If P[1] = 1 (Zero Flag)
	if P[1] == 1 {

		if Debug {
			fmt.Printf("\tOpcode %02X%02X [2 bytes] [Mode: Relative]\tBNE  Branch on Result not Zero.\t| Zero Flag(P1) = %d | PC += 2\n", Opcode, Memory[PC+1], P[1])
		}
		PC += 2
	// If P[1] = 0 (Not Zero) Jump to address
	} else {

		if Debug {
			fmt.Printf("\tOpcode %02X%02X [2 bytes]\tBNE  Branch on Result not Zero.\tZero Flag(P1) = %d, JUMP TO %04X\n", Opcode, Memory[PC+1], P[1], PC+2+uint16(value) )
		}

		// Current PC (To detect page bounday cross)
		tmp := PC
		// fmt.Printf("\ntmp: %02X\n",tmp)

		// PC + the number of bytes to jump on carry clear
		PC += uint16(value)

		// Increment PC
		PC += 2

		// Add 1 to cycles if branch occurs on same page
		Beam_index += 1

		// Add one extra cycle if branch occurs in a differente memory page
		if MemPageBoundary(tmp, PC) {
			Beam_index += 1
		}
	}

	Beam_index += 2
}
