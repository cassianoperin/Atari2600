package VGS

import	"os"
import	"fmt"

// LDA  Load Accumulator with Memory (immidiate)
//
//      M -> A                           N Z C I D V
//                                       + + - - - -
//
//      addressing    assembler    opc  bytes  cycles
//      --------------------------------------------
//      immediate     LDA #oper     A9    2     2
//      zeropage      LDA oper      A5    2     3
//      absolute,Y    LDA oper,Y    B9    3     4*
//      (indirect),Y  LDA (oper),Y  B1    2     5*
//      zeropage,X    LDA oper,X    B5    2     4
//      absolute      LDA oper      AD    3     4
//      absolute,X    LDA oper,X    BD    3     4*
func opc_LDA(memAddr uint16, mode string, bytes uint16, opc_cycles byte) {

	// Some tests of instructions that tryes to read from TIA addresses (00 - 127)
	if memAddr < 0x80 {
		fmt.Printf("LDA - Tryed to read from TIA ADDRESS! Memory[%X]\tEXIT\n", memAddr)
		os.Exit(2)
	}

	// Increment the beam
	beamIndex ++

	// Check for extra cycles (*) in the first opcode cycle
	if opc_cycle_count == 1 {
		if opcode == 0xB9 || opcode == 0xBD || opcode == 0xB1 {
			// Add 1 to cycles if page boundery is crossed
			if MemPageBoundary(memAddr, PC) {
				opc_cycle_extra = 1
			}
		}
	}

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

		A = Memory[memAddr]

		if Debug {
			if bytes == 2 {
				dbg_show_message = fmt.Sprintf("\n\tOpcode %02X%02X [2 bytes] [Mode: %s]\tLDA  Load Accumulator with Memory.\tA = Memory[%02X] (%d)\n", opcode, Memory[PC+1], mode, memAddr, A)
				fmt.Println(dbg_show_message)
			} else if bytes == 3 {
				dbg_show_message = fmt.Sprintf("\n\tOpcode %02X %02X%02X [3 bytes] [Mode: %s]\tLDA  Load Accumulator with Memory.\tA = Memory[%02X] (%d)\n", opcode, Memory[PC+2], Memory[PC+1], mode, memAddr, A)
				fmt.Println(dbg_show_message)
			}

			// Collect data for debug interface after finished running the opcode
			dbg_opcode_message("LDA", bytes, opc_cycle_count + opc_cycle_extra)
		}

		flags_Z(A)
		flags_N(A)

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1

		// Reset Opcode Extra Cycle counter
		opc_cycle_extra = 0
	}

}
