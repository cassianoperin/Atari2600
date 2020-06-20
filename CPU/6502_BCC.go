package CPU

import	"fmt"

// BCC  Branch on Carry Clear
//
//      branch on C = 0                  N Z C I D V
//                                       - - - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      relative      BCC oper      90    2     2**
func opc_BCC(offset uint16) {
	// If carry is clear
	if P[0] == 0 {

		if Debug {
			fmt.Printf("\n\tOpcode %02X%02X [2 bytes] [Mode: Relative]\tBCC  Branch on Carry Clear (relative).\tCarry EQUAL 0, JUMP TO %04X\n", Opcode, Memory[PC+1], PC+2+offset )
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

	// If carry is set
	} else {
		if Debug {
			fmt.Printf("\n\tOpcode %02X%02X [2 bytes]\tBCC  Branch on Carry Clear (relative).\tCarry NOT EQUAL 0, PC+2 \n", Opcode, Memory[PC+1])
		}
		PC += 2
	}

	Beam_index += 2
}
