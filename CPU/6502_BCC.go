package CPU

import	"fmt"

// BCC  Branch on Carry Clear
//
//      branch on C = 0                  N Z C I D V
//                                       - - - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      relative      BCC oper      90    2     2**
func opc_BCC(value int8, bytes uint16, opc_cycles byte) {	// value is SIGNED

	// Increment the beam
	Beam_index ++

	// If carry is clear
	if P[0] == 0 {

		// Check for extra cycles (*) in the first opcode cycle
		if opc_cycle_count == 1 {
			// Add 1 to cycles if page boundery is crossed
			if MemPageBoundary(PC, PC + uint16(value) + 2 ) { // REVISAR SE A LOGICA ESTA CORRETA
				opc_cycle_extra = 1
				fmt.Println("PAUSA PARA VALIDAR BCC COM PAGE BOUNDARY")
				Pause = true
			}
		}

		// Show current opcode cycle
		if Debug {
			fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\t(%d cycles + 1 cycle for branch + %d extra cycles for branch in different page)\n", Cycle, opc_cycle_count, opc_cycles + opc_cycle_extra + 1, opc_cycles, opc_cycle_extra)
		}

		// Just increment the Opcode cycle Counter
		if opc_cycle_count < opc_cycles + 1 + opc_cycle_extra {
			opc_cycle_count ++

		// After spending the cycles needed, execute the opcode
		} else {
			if Debug {
				fmt.Printf("\tOpcode %02X%02X [2 bytes] [Mode: Relative]\tBCC  Branch on Carry Clear.\tCarry EQUAL 0, JUMP TO %04X\n", Opcode, Memory[PC+1], PC+2+uint16(value) )
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


	// If carry is set
	} else {

		// Show current opcode cycle
		if Debug {
			fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\n", Cycle, opc_cycle_count, opc_cycles)
		}

		// Just increment the Opcode cycle Counter
		if opc_cycle_count < opc_cycles {
			opc_cycle_count ++

		// After spending the cycles needed, execute the opcode
		} else {
			if Debug {
				fmt.Printf("\n\tOpcode %02X%02X [2 bytes]\tBCC  Branch on Carry Clear.\tCarry NOT EQUAL 0, PC+2\n", Opcode, Memory[PC+1])
			}

			// Increment PC
			PC += bytes

			// Reset Opcode Cycle counter
			opc_cycle_count = 1
		}

	}

}
