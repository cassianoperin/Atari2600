package VGS

import	"fmt"

// ROL  Rotate One Bit Left (Memory or Accumulator)
//
//      C <- [76543210] <- C             N Z C I D V
//                                       + + + - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      zeropage      ROL oper      26    2     5
func opc_ROL(memAddr uint16, mode string, bytes uint16, opc_cycles byte) {

	// Increment the beam
	beamIndex ++

	// Show current opcode cycle
	if Debug {
		fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\n", counter_F_Cycle, opc_cycle_count, opc_cycles)

		// Collect data for debug interface just on first cycle
		if opc_cycle_count == 1 {
			debug_opc_text		= fmt.Sprintf("%04x     ROL      ;%d", PC, opc_cycles)
			dbg_opc_bytes		= bytes
			dbg_opc_opcode		= opcode
			dbg_opc_payload1	= Memory[PC+1]
		}
	}

	// Just increment the Opcode cycle Counter
	if opc_cycle_count < opc_cycles {
		opc_cycle_count ++

	// After spending the cycles needed, execute the opcode
	} else {

		// Original Carry Value
		carry_orig := P[0]

		if Debug {
			fmt.Printf("\tOpcode %02X%02X [2 bytes] [Mode: %s]\tROL  Rotate One Bit Left.\tMemory[%d] Roll Left 1 bit\t(%d)\n", opcode, Memory[PC+1], mode, memAddr, ( Memory[memAddr] << 1 ) + carry_orig )
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
