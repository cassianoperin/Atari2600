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
		fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\n", counter_F_Cycle, opc_cycle_count, opc_cycles)
	}

	// Just increment the Opcode cycle Counter
	if opc_cycle_count < opc_cycles {
		opc_cycle_count ++

		// Reset to default value
		TIA_Update = -1

	// After spending the cycles needed, execute the opcode
	} else {

		// Update Memory[memAddr] with value of A and notify TIA about the update
		memUpdate(memAddr, A)

		if Debug {
			if bytes == 2 {
				dbg_show_message = fmt.Sprintf("\n\tOpcode %02X%02X [2 bytes] [Mode: %s]\tSTA  Store Accumulator in Memory.\tMemory[%02X] = A (%d)\n", opcode, Memory[PC+1], mode, memAddr, Memory[memAddr] )
				fmt.Println(dbg_show_message)
			} else if bytes == 3 {
				dbg_show_message = fmt.Sprintf("\n\tOpcode %02X %02X%02X [3 bytes] [Mode: %s]\tSTA  Store Accumulator in Memory.\tMemory[%02X] = A (%d)\n", opcode, Memory[PC+2], Memory[PC+1], mode, memAddr, Memory[memAddr] )
				fmt.Println(dbg_show_message)
			}

			// Collect data for debug interface after finished running the opcode
			dbg_opcode_message("STA", bytes, opc_cycle_count + opc_cycle_extra)
		}

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1
	}

}
