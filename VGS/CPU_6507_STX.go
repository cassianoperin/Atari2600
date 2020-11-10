package VGS

import	"fmt"

// STX  Store Index X in Memory (zeropage)
//
//      X -> M                           N Z C I D V
//                                       - - - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      zeropage      STX oper      86    2     3
func opc_STX(memAddr uint16, mode string, bytes uint16, opc_cycles byte) {

	// Increment the beam
	beamIndex ++

	// Show current opcode cycle
	if Debug {
		fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\n", counter_F_Cycle, opc_cycle_count, opc_cycles)
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

		Memory[memAddr] = X

		if Debug {
			fmt.Printf("\n\tOpcode %02X%02X [2 bytes] [Mode: %s]\tSTX  Store Index X in Memory.\tMemory[%02X] = X (%d)\n", opcode, Memory[PC+1], mode, memAddr, X)
		}

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1
	}

}
