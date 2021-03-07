package VGS

// import	"os"
import	"fmt"

// ROR  Rotate One Bit Right (Memory or Accumulator)
//
//      C -> [76543210] -> C             N Z C I D V
//                                       + + + - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      accumulator      ROR A      6A    1     2

// Move each of the bits in either A or M one place to the right.
// Bit 7 is filled with the current value of the carry flag whilst the old bit 0 becomes the new carry flag value.

// ------------------------------------ Accumulator ------------------------------------ //

func opc_ROR_A(bytes uint16, opc_cycles byte) {

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

		// Keep original Accumulator value for debug
		original_A := A
		original_carry := P[0]
		// fmt.Printf("A original: %08b\n", A)

		// Keep the original bit 0 from Accumulator to be used as new Carry
		new_Carry := A & 0x01
		// fmt.Printf("New Carry: %b\n",new_Carry) // NEW CARRY

		// Shift Right Accumulator
		A = A >> 1
		// fmt.Printf("A shifted: %08b\nP[0]: %d\n", A, P[0])

		// Bit 7 is filled with the current value of the carry flag
		A += P[0] << 7
		// fmt.Printf("A final: %08b\n", A)

		// The old bit 0 becomes the new carry flag value
		P[0] = new_Carry
		// fmt.Printf("New Carry: %b\n", P[0])


		if Debug {
			dbg_show_message = fmt.Sprintf("\n\tOpcode %02X%02X [2 bytes] [Mode: Accumulator]\tROR  Rotate One Bit Right.\tA(%d) Roll Right 1 bit\t(%d) + Current Carry(%d) as new bit 7.\tA = %d\n", opcode, Memory[PC+1], original_A, A >> 1, P[0], A )
			// dbg_show_message = fmt.Sprintf("\n\tOpcode %02X%02X [2 bytes] [Mode: %s]\tROR  Rotate One Bit Right.\tMemory[%d] Roll Left 1 bit\t(%d)\n", opcode, Memory[PC+1], mode, memAddr, ( Memory[memAddr] << 1 ) + carry_orig )
			fmt.Println(dbg_show_message)

			// Collect data for debug interface after finished running the opcode
			dbg_opcode_message("ROR", bytes, opc_cycle_count + opc_cycle_extra)
		}


		flags_N(Memory[memAddr])
		flags_Z(Memory[memAddr])
		if Debug {
			fmt.Printf("\tFlag C: %d -> %d", original_carry, P[0])
		}

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1

	}


}


// -------------------------------------- Memory --------------------------------------- //
// func opc_ROR_Meeee(memAddr uint16, mode string, bytes uint16, opc_cycles byte) {
//
// 	// Some tests of instructions that tryes to read from TIA addresses (00 - 127)
// 	if memAddr < 0x80 {
// 		fmt.Printf("ROR - Tryed to read from TIA ADDRESS! Memory[%X]\tEXIT\n", memAddr)
// 		os.Exit(2)
// 	}
//
// 	fmt.Printf("A: %08b", A)
//
// 	// Increment the beam
// 	beamIndex ++
//
// 	// Show current opcode cycle
// 	if Debug {
// 		fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\n", counter_F_Cycle, opc_cycle_count, opc_cycles)
// 	}
//
// 	// Just increment the Opcode cycle Counter
// 	if opc_cycle_count < opc_cycles {
// 		opc_cycle_count ++
//
// 		// Reset to default value
// 		TIA_Update = -1
//
// 	// After spending the cycles needed, execute the opcode
// 	} else {
//
// 		// Original Carry Value
// 		carry_orig := P[0]
//
// 		if Debug {
// 			dbg_show_message = fmt.Sprintf("\n\tOpcode %02X%02X [2 bytes] [Mode: %s]\tROR  Rotate One Bit Right.\tA[%d] Roll Left 1 bit\t(%d)\n", opcode, Memory[PC+1], mode, A, ( Memory[memAddr] << 1 ) + carry_orig )
// 			// dbg_show_message = fmt.Sprintf("\n\tOpcode %02X%02X [2 bytes] [Mode: %s]\tROR  Rotate One Bit Right.\tMemory[%d] Roll Left 1 bit\t(%d)\n", opcode, Memory[PC+1], mode, memAddr, ( Memory[memAddr] << 1 ) + carry_orig )
// 			fmt.Println(dbg_show_message)
//
// 			// Collect data for debug interface after finished running the opcode
// 			dbg_opcode_message("ROL", bytes, opc_cycle_count + opc_cycle_extra)
// 		}
//
// 		// Calculate the original bit7 and save it as the new Carry
// 		P[0] = Memory[memAddr] & 0x80 >> 7
//
// 		// Shift left the byte and put the original bit7 value in bit 1 to make the complete ROL
// 		A = ( A >> 1 )
//
// 		flags_N(Memory[memAddr])
// 		flags_Z(Memory[memAddr])
// 		if Debug {
// 			fmt.Printf("\tFlag C: %d -> %d", carry_orig, P[0])
// 		}
//
// 		// Increment PC
// 		PC += bytes
//
// 		// Reset Opcode Cycle counter
// 		opc_cycle_count = 1
// 	}
//
// 	os.Exit(2)
// }
