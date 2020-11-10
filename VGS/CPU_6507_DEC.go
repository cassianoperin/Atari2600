package VGS

import	"fmt"

// DEC  Decrement Memory by One
//
//      M - 1 -> M                       N Z C I D V
//                                       + + - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      zeropage      DEC oper      C6    2     5
func opc_DEC(memAddr uint16, mode string, bytes uint16, opc_cycles byte) {

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

		if Debug {
			fmt.Printf("\n\tOpcode %02X%02X [2 bytes] [Mode: %s]\tDEC  Decrement Memory by One.\tMemory[%02X] -= 1 (%d)\n", opcode, Memory[PC+1], mode, memAddr, Memory[memAddr] - 1 )
		}
		Memory[memAddr] -= 1

		flags_Z(Memory[memAddr])
		flags_N(Memory[memAddr])

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1
	}

}
