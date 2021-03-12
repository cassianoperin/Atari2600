package VGS

import	"os"
import	"fmt"

// LDX  Load Index X with Memory
//
//      M -> X                           N Z C I D V
//                                       + + - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      immediate     LDX #oper     A2    2     2
//      zeropage	    LDX oper	    A6  	2	    3
func opc_LDX(memAddr uint16, mode string, bytes uint16, opc_cycles byte) {

	// Atari 2600 interpreter mode
	if CPU_MODE == 0 {
		// Some tests of instructions that tryes to read from TIA addresses (00 - 127)
		if memAddr < 0x80 {
			fmt.Printf("LDX - Tryed to read from TIA ADDRESS! Memory[%X]\tEXIT\n", memAddr)
			os.Exit(2)
		}

		// Some tests of instructions that tryes to read from RIOT addresses (640 - 671)
		if memAddr > 0x280 &&  memAddr <= 0x29F {
			fmt.Printf("LDX - Tryed to read from RIOT ADDRESS! Memory[%X]\tEXIT\n", memAddr)
			os.Exit(2)
		}
	}

	// Increment the beam
	beamIndex ++

	// Check for extra cycles (*) in the first opcode cycle
	// if opc_cycle_count == 1 {
	// 	if Opcode == 0xBE {
	// 		// Add 1 to cycles if page boundery is crossed
	// 		if MemPageBoundary(memAddr, PC) {
	// 			opc_cycle_extra = 1
	// 		}
	// 	}
	// }

	// Show current opcode cycle
	if Debug {
		fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\t(%d cycles + %d extra cycles)\n", counter_F_Cycle, opc_cycle_count, opc_cycles + opc_cycle_extra, opc_cycles, opc_cycle_extra)
	}

	// Just increment the Opcode cycle Counter
	if opc_cycle_count < opc_cycles +  opc_cycle_extra {
		opc_cycle_count ++

		// Reset to default value
		TIA_Update = -1

	// After spending the cycles needed, execute the opcode
	} else {

		X = Memory[memAddr]

		if Debug {
			dbg_show_message = fmt.Sprintf("\n\tOpcode %02X%02X [2 bytes] [Mode: %s]\tLDX  Load Index X with Memory.\tX = Memory[%02X] (%d)\n", opcode, Memory[PC+1], mode, PC+1, X)
			fmt.Println(dbg_show_message)

			// Collect data for debug interface after finished running the opcode
			dbg_opcode_message("LDX", bytes, opc_cycle_count + opc_cycle_extra)
		}

		flags_Z(X)
		flags_N(X)

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1

		// Reset Opcode Extra Cycle counter
		opc_cycle_extra = 0
	}

}
