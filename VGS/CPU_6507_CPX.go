package VGS

import	"os"
import	"fmt"

// CPX  Compare Memory and Index X
//
//      X - M                            N Z C I D V
//                                       + + + - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      immediate     CPX #oper     E0    2     2
//      zeropage      CPX oper    	E4  	2	    3
func opc_CPX(memAddr uint16, mode string, bytes uint16, opc_cycles byte) {

	// Some tests of instructions that tryes to read from TIA addresses (00 - 127)
	if memAddr < 0x80 {
		fmt.Printf("CPX - Tryed to read from TIA ADDRESS! Memory[%X]\tEXIT\n", memAddr)
		os.Exit(2)
	}

	// Some tests of instructions that tryes to read from RIOT addresses (640 - 671)
	if memAddr > 0x280 &&  memAddr <= 0x29F {
		fmt.Printf("CPX - Tryed to read from RIOT ADDRESS! Memory[%X]\tEXIT\n", memAddr)
		os.Exit(2)
	}

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

		tmp := X - Memory[memAddr]

		if Debug {
			if tmp == 0 {
				dbg_show_message = fmt.Sprintf("\n\tOpcode %02X%02X [2 bytes] [Mode: %s]\tCPX  Compare Memory and Index X.\tX(%d) - Memory[%02X](%d) = (%d) EQUAL\n", opcode, Memory[PC+1], mode, X, PC+1, Memory[memAddr], tmp)
				fmt.Println(dbg_show_message)
			} else {
				dbg_show_message = fmt.Sprintf("\n\tOpcode %02X%02X [2 bytes] [Mode: %s]\tCPX  Compare Memory and Index X.\tX(%d) - Memory[%02X](%d) = (%d) NOT EQUAL\n", opcode, Memory[PC+1], mode, X, PC+1, Memory[memAddr], tmp)
				fmt.Println(dbg_show_message)
			}

			// Collect data for debug interface after finished running the opcode
			dbg_opcode_message("CPX", bytes, opc_cycle_count + opc_cycle_extra)
		}

		flags_Z(tmp)
		flags_N(tmp)
		flags_C(X,Memory[memAddr])

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1
	}

}
