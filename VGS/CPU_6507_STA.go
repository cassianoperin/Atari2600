package VGS

import	"fmt"

// STA  Store Accumulator in Memory (zeropage,X)
//
//      A -> M                           N Z C I D V
//                                       - - - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      zeropage,X    STA oper,X    95    2     4
//      zeropage      STA oper      85    2     3
//      absolute,Y    STA oper,Y    99    3     5
//      absolute      STA oper      8D    3     4
func opc_STA(memAddr uint16, mode string, bytes uint16, opc_cycles byte) {

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

		// Change variable to a positive number to TIA detect the change
		if memAddr < 128 {
			TIA_Update = int8(memAddr)
		}

		Memory[ memAddr ] = A

		if Debug {
			if bytes == 2 {
				fmt.Printf("\n\tOpcode %02X%02X [2 bytes] [Mode: %s]\tSTA  Store Accumulator in Memory.\tMemory[%02X] = A (%d)\n", opcode, Memory[PC+1], mode, memAddr, Memory[memAddr] )

			} else if bytes == 3 {
				fmt.Printf("\n\tOpcode %02X %02X%02X [3 bytes] [Mode: %s]\tSTA  Store Accumulator in Memory.\tMemory[%02X] = A (%d)\n", opcode, Memory[PC+2], Memory[PC+1], mode, memAddr, Memory[memAddr] )
			}
		}

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1
	}

}
