package CPU

import	"fmt"

// BCS  Branch on Carry Set
//
//      branch on C = 1                  N Z C I D V
//                                       - - - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      relative      BCS oper      B0    2     2**
func opc_BCS(offset uint16) {
	// If carry is clear
	if P[0] == 1 {

		if Debug {
			fmt.Printf("\n\tOpcode %02X%02X [2 bytes] [Mode: Relative]\tBCS  Branch on Carry Set (relative).\tCarry EQUAL 1, JUMP TO %04X\n", Opcode, Memory[PC+1], PC+2+offset )
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

		// // Add one extra cycle if branch occurs in a differente memory page
		if MemPageBoundary(uint16(tmp), PC) {
			Beam_index += 1
		}

	// If carry is set
	} else {
		if Debug {
			fmt.Printf("\n\tOpcode %02X%02X [2 bytes]\tBCS  Branch on Carry Set (relative).\tCarry NOT EQUAL 1, PC+2 \n", Opcode, Memory[PC+1])
		}
		PC += 2
	}

	Beam_index += 2
}
