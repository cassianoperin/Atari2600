package VGS

import	"os"
import	"fmt"

// ROL  Rotate One Bit Left (Memory or Accumulator)
//
//      C <- [76543210] <- C             N Z C I D V
//                                       + + + - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      zeropage      ROL oper      26    2     5

//Move each of the bits in either A or M one place to the left.
//Bit 0 is filled with the current value of the carry flag whilst the old bit 7 becomes the new carry flag value.
func opc_ROL(memAddr uint16, mode string, bytes uint16, opc_cycles byte) {

	// Some tests of instructions that tryes to read from TIA addresses (00 - 127)
	if memAddr < 0x80 {
		fmt.Printf("ROL - Tryed to read from TIA ADDRESS! Memory[%X]\tEXIT\n", memAddr)
		os.Exit(2)
	}

	// Some tests of instructions that tryes to read from RIOT addresses (640 - 671)
	if memAddr > 0x280 &&  memAddr <= 0x29F {
		fmt.Printf("ROL - Tryed to read from RIOT ADDRESS! Memory[%X]\tEXIT\n", memAddr)
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

		// Original Carry Value
		carry_orig := P[0]

		if Debug {
			dbg_show_message = fmt.Sprintf("\n\tOpcode %02X%02X [2 bytes] [Mode: %s]\tROL  Rotate One Bit Left.\tMemory[%d] Roll Left 1 bit\t(%d)\n", opcode, Memory[PC+1], mode, memAddr, ( Memory[memAddr] << 1 ) + carry_orig )
			fmt.Println(dbg_show_message)

			// Collect data for debug interface after finished running the opcode
			dbg_opcode_message("ROL", bytes, opc_cycle_count + opc_cycle_extra)
		}

		// Calculate the original bit7 and save it as the new Carry
		P[0] = Memory[memAddr] & 0x80 >> 7

		// Shift left the byte and put the original bit7 value in bit 1 to make the complete ROL
		Memory[memAddr] = ( Memory[memAddr] << 1 ) + carry_orig

		flags_N(Memory[memAddr])
		flags_Z(Memory[memAddr])
		if Debug {
			fmt.Printf("\tFlag C: %d -> %d", carry_orig, P[0])
		}

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1
	}

}
