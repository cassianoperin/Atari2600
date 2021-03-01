package VGS

import	"os"
import	"fmt"

// BIT  Test Bits in Memory with Accumulator
//
//      bits 7 and 6 of operand are transfered to bit 7 and 6 of SR (N,V);
//      the zeroflag is set to the result of operand AND accumulator.
//
//      A AND M, M7 -> N, M6 -> V        N Z C I D V
//                                      M7 + - - - M6
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      zeropage      BIT oper      24    2     3
//      absolute      BIT oper      2C    3     4
func opc_BIT(memAddr uint16, mode string, bytes uint16, opc_cycles byte) {

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

		if Debug {
			if bytes == 2 {
				dbg_show_message = fmt.Sprintf("\n\tOpcode %02X %02X [2 bytes] [Mode: %s]\tBIT  Test Bits in Memory with Accumulator.\tA (%08b) AND Memory[%04X] (%08b) = %08b \tM7 -> N, M6 -> V\n", opcode, Memory[PC+1], mode, A, memAddr, Memory[memAddr], A & Memory[memAddr] )
				fmt.Println(dbg_show_message)
			} else if bytes == 3 {
				dbg_show_message = fmt.Sprintf("\n\tOpcode %02X %02X%02X [3 bytes] [Mode: %s]\tBIT  Test Bits in Memory with Accumulator.\tA (%08b) AND Memory[%04X] (%08b) = %08b \tM7 -> N, M6 -> V\n", opcode, Memory[PC+2], Memory[PC+1], mode, A, memAddr, Memory[memAddr], A & Memory[memAddr] )
				fmt.Println(dbg_show_message)
			}

			// Collect data for debug interface after finished running the opcode
			dbg_opcode_message("BIT", bytes, opc_cycle_count + opc_cycle_extra)
		}

		// Memory Address bit 7 (A) -> N (Negative)
		if Debug {
			fmt.Printf("\tFlag N: %d -> ",P[7])
		}
		P[7] = A >> 7 & 0x1
		if Debug {
			fmt.Printf("%d\n",P[7])
		}

		// Memory Address bit 6 (A) -> V (oVerflow)
		if Debug {
			fmt.Printf("\tFlag V: %d -> ",P[6])
		}
		P[6] = A >> 6 & 0x1
		if Debug {
			fmt.Printf("%d\n",P[6])
		}

		// FIRST ATTEMPT TO DETECT ACCESS TO A TIA READ ONLY REGISTER (0x00-0x0D)
		if memAddr < 14 {
			flags_Z(A & Memory_TIA_RO[memAddr])
		// Read from other reserved TIA registers
		} else if memAddr < 128 {
			fmt.Printf("BIT - Controlled Exit to map access to TIA Write Addresses. COULD BE MIRRORS!!!!!.\t EXITING\n")
			os.Exit(2)
		// Read from RIOT Memory Map (> 0x280)
		} else {
			flags_Z(A & Memory[memAddr])
		}

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1
	}

}
