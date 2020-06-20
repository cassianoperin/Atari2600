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
func opc_BNE(offset uint16) {
	// If P[1] = 1 (Zero Flag)
	if P[1] == 1 {

		if Debug {
			fmt.Printf("\n\tOpcode %02X%02X [2 bytes] [Mode: Relative]\tBNE  Branch on Result not Zero.\t| Zero Flag(P1) = %d | PC += 2\n", Opcode, Memory[PC+1], P[1])
		}
		PC += 2

	} else {

		if Debug {
			fmt.Printf("\n\tOpcode %02X%02X [2 bytes]\tBNE  Branch on Result not Zero.\tZero Flag(P1) = %d, JUMP TO %04X\n", Opcode, Memory[PC+1], P[1], PC+2+offset )
		}

		// Current PC (To detect page bounday cross)
		tmp := PC
		// fmt.Printf("\ntmp: %02X\n",tmp)

		// PC + the number of bytes to jump on carry clear
		PC += offset

		// Increment PC
		PC += 2

		// Add 1 to cycles if branch occurs on same page
		Beam_index += 1

		// Add one extra cycle if branch occurs in a differente memory page
		if MemPageBoundary(uint16(tmp), PC) {
			Beam_index += 1
		}
		// os.Exit(2)
	}
	Beam_index += 2
}
