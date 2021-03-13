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
		// Bigger than 63 (READ ONLY TIA) is allowed
		if memAddr > 0x3F && memAddr < 0x80 {
			fmt.Printf("LDX - Tryed to read from TIA ADDRESS! Memory[%X]\tEXIT\n", memAddr)
			os.Exit(2)
		}
		// Read from RIOT RO addresses are allowed (0x280(640) - 0x29F(671))
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

		// Atari 2600 interpreter mode
		if CPU_MODE == 0 {
			// Read from TIA RO Registers (0x00-0x0D)
			if memAddr < 64 {
				X = Memory_TIA_RO[memAddr]
			// Read from other reserved TIA registers
			} else if memAddr < 128 {
				fmt.Printf("LDX - Controlled Exit to map access to TIA Write Addresses. COULD BE MIRRORS!!!!!.\t EXITING\n")
				os.Exit(2)
			// Read from RIOT Memory Map (> 0x280)
			} else {
				X = Memory[memAddr]
			}
		// 6507 interpreter mode
		} else {
			X = Memory[memAddr]
		}


		if Debug {
			if bytes == 2 {
				dbg_show_message = fmt.Sprintf("\n\tOpcode %02X%02X [2 bytes] [Mode: %s]\tLDX  Load Index X with Memory.\tX = Memory[%02X] (%d)\n", opcode, Memory[PC+1], mode, memAddr, X)
				fmt.Println(dbg_show_message)
			} else if bytes == 3 {
				dbg_show_message = fmt.Sprintf("\n\tOpcode %02X %02X%02X [3 bytes] [Mode: %s]\tLDX  Load Index X with Memory.\tX = Memory[%02X] (%d)\n", opcode, Memory[PC+2], Memory[PC+1], mode, memAddr, X)
				fmt.Println(dbg_show_message)
			}

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
