package VGS

import	"fmt"

// BEQ  Branch on Result Zero
//
//      branch on Z = 1                  N Z C I D V
//                                       - - - - - -
//
//      addressing	assembler	   opc	bytes	 cyles
//      --------------------------------------------
//      relative	  BEQ oper	    F0  	2	     2**
func opc_BEQ(value int8, bytes uint16, opc_cycles byte) {	// value is SIGNED

	// Increment the beam
	beamIndex ++

	// If zero flag is set
	if P[1] == 1 {

		// Check for extra cycles (*) in the first opcode cycle
		if opc_cycle_count == 1 {
			// Add 1 to cycles if page boundery is crossed
			if MemPageBoundary(PC, PC + uint16(value) + 2 ) {
				opc_cycle_extra = 1
				fmt.Println("TEST BEQ COM PAGE BOUNDARY")
				Pause = true
			}
		}

		// Show current opcode cycle
		if Debug {
			fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\t(%d cycles + 1 cycle for branch + %d extra cycles for branch in different page)\n", counter_F_Cycle, opc_cycle_count, opc_cycles + opc_cycle_extra + 1, opc_cycles, opc_cycle_extra)
		}

		// Just increment the Opcode cycle Counter
		if opc_cycle_count < opc_cycles + 1 + opc_cycle_extra {
			opc_cycle_count ++

			// Reset to default value
			TIA_Update = -1

		// After spending the cycles needed, execute the opcode
		} else {
			if Debug {
				dbg_show_message = fmt.Sprintf("\n\tOpcode %02X%02X [2 bytes] [Mode: Relative]\tBEQ  Branch on Result Zero.\tZero flag EQUAL 1, JUMP TO %04X\n", opcode, Memory[PC+1], PC+2+uint16(value) )
				fmt.Println(dbg_show_message)

				// Collect data for debug interface after finished running the opcode
				dbg_opcode_message("BEQ", bytes, opc_cycle_count + opc_cycle_extra)

			}

			// PC + the number of bytes to jump on carry clear
			PC += uint16(value)

			// Increment PC
			PC += bytes

			// Reset Opcode Cycle counter
			opc_cycle_count = 1

			// Reset Opcode Extra Cycle counter
			opc_cycle_extra = 0
		}


	// If zero flag is clear
	} else {

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
				dbg_show_message = fmt.Sprintf("\n\tOpcode %02X%02X [2 bytes]\tBEQ  Branch on Result Zero.\tZero flag NOT EQUAL 1, PC+2 \n", opcode, Memory[PC+1])
				fmt.Println(dbg_show_message)

				// Collect data for debug interface after finished running the opcode
				dbg_opcode_message("BEQ", bytes, opc_cycle_count + opc_cycle_extra)
			}

			// Increment PC
			PC += bytes

			// Reset Opcode Cycle counter
			opc_cycle_count = 1
		}

	}

}
