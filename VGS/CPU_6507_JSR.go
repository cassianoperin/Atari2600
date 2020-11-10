package VGS

import	"fmt"

// JSR  Jump to New Location Saving Return Address
//
//      push (PC+2) to Stack,            N Z C I D V
//      (PC+1) -> PCL                    - - - - - -
//      (PC+2) -> PCH
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      absolute      JSR oper      20    3     6
func opc_JSR(memAddr uint16, mode string, bytes uint16, opc_cycles byte) {

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

		// Push PC+2 (will be increased in 1 in RTS to match the next address (3 bytes operation))
		// Store the first byte into the Stack
		Memory[SP] = byte( (PC+2) >> 8 )
		SP--
		// Store the second byte into the Stack
		Memory[SP] = byte( (PC+2) & 0xFF )
		SP--
		// fmt.Printf("\nPC+3: %02X",PC+3)
		// fmt.Printf("\nF0: %02X",(PC+3) >> 8)
		// fmt.Printf("\n42: %02X",(PC+3) & 0xFF)

		if Debug {
			fmt.Printf("\tOpcode %02X %02X%02X [3 bytes] [Mode: %s]\tJSR  Jump to New Location Saving Return Address.\tPC = Memory[%02X]\t|\t Stack[%02X] = %02X\t Stack[%02X] = %02X\n", opcode, Memory[PC+2], Memory[PC+1], mode, memAddr, SP+2, Memory[SP+2], SP+1, Memory[SP+1])
		}

		// Update PC
		PC = memAddr

		// Reset Opcode Cycle counter
		opc_cycle_count = 1
	}

}
