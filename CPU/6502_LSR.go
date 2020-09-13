package CPU

import	"fmt"

// LSR  Shift One Bit Right (Memory or Accumulator)
//
//      0 -> [76543210] -> C             N Z C I D V
//                                       0 + + - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      accumulator   LSR A         4A    1     2
func opc_LSR(bytes uint16, opc_cycles byte) {

	// Increment the beam
	Beam_index ++

	// Show current opcode cycle
	if Debug {
		fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\n", Cycle, opc_cycle_count, opc_cycles)
	}

	// Just increment the Opcode cycle Counter
	if opc_cycle_count < opc_cycles {
		opc_cycle_count ++

	// After spending the cycles needed, execute the opcode
	} else {

		// Save the original Carry value
		carry_orig := P[0]

		// Least significant bit turns into the new Carry
		P[0] = A & 0x01

		if Debug {
			fmt.Printf("\tOpcode %02X [1 byte] [Mode: Accumulator]\tLSR  Shift One Bit Right.\tA = A(%d) Shift Right 1 bit\t(%d)\n", Opcode, A, A >> 1 )
		}

		A = A >> 1

		flags_N(A)
		flags_Z(A)
		if Debug {
			fmt.Printf("\tFlag C: %d -> %d", carry_orig, P[0])
		}

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1
	}

}
