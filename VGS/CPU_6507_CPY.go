package VGS

import	"fmt"

// CPY  Compare Memory and Index Y (immidiate)
//
//      Y - M                            N Z C I D V
//                                       + + + - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      immidiate     CPY #oper     C0    2     2
//      zeropage      CPY oper      C4    2     3
func opc_CPY(memAddr uint16, mode string, bytes uint16, opc_cycles byte) {

	// Increment the beam
	beamIndex ++

	// Show current opcode cycle
	if Debug {
		fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\n", Cycle, opc_cycle_count, opc_cycles)
	}

	// Just increment the Opcode cycle Counter
	if opc_cycle_count < opc_cycles {
		opc_cycle_count ++

	// After spending the cycles needed, execute the opcode
	} else {

		tmp := Y - Memory[memAddr]

		if Debug {
			if tmp == 0 {
				fmt.Printf("\n\tOpcode %02X%02X [2 bytes] [Mode: %s]\tCPY  Compare Memory and Index Y.\tY(%d) - Memory[%02X](%d) = (%d) EQUAL\n", opcode, Memory[PC+1], mode, Y, PC+1, Memory[memAddr], tmp)
			} else {
				fmt.Printf("\n\tOpcode %02X%02X [2 bytes] [Mode: %s]\tCPY  Compare Memory and Index Y.\tY(%d) - Memory[%02X](%d) = (%d) NOT EQUAL\n", opcode, Memory[PC+1], mode, Y, PC+1, Memory[memAddr], tmp)
			}
		}

		flags_Z(tmp)
		flags_N(tmp)
		flags_C(Y,Memory[memAddr])

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1
	}

}
