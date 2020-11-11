package VGS

import	"fmt"

// SEI  Set Interrupt Disable Status
//
//      1 -> I                           N Z C I D V
//                                       - - - 1 - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      implied       SEI           78    1     2
func opc_SEI(bytes uint16, opc_cycles byte) {

	// Increment the beam
	beamIndex ++

	// Show current opcode cycle
	if Debug {
		fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\n", counter_F_Cycle, opc_cycle_count, opc_cycles)
		// debug_opc_text = fmt.Sprintf("%04x\tSEI\t%d\t%02x", PC, bytes, opcode)
		dbg_opc_PC			= PC
		dbg_opc_mnm			= "SEI"
		dbg_opc_bytes		= bytes
		dbg_opc_opcode		= opcode
		dbg_opc_payload1	= Memory[PC+1]
	}

	// Just increment the Opcode cycle Counter
	if opc_cycle_count < opc_cycles {
		opc_cycle_count ++

	// After spending the cycles needed, execute the opcode
	} else {

		P[2]	=  1

		if Debug {
			fmt.Printf("\n\tOpcode %02X [1 byte] [Mode: Implied]\tSEI  Set Interrupt Disable Status.\tP[2]=%d\n", opcode, P[2])
		}

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1
	}

}
