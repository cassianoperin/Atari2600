package VGS

import	"fmt"

// ASL  Shift Left One Bit (Memory or Accumulator)
//
//      C <- [76543210] <- 0             N Z C I D V
//                                       + + + - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      accumulator   ASL A         0A    1     2
func opc_ASL(bytes uint16, opc_cycles byte) {

	// Increment the beam
	beamIndex ++

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
			fmt.Printf("\n\tOpcode %02X [1 byte] [Mode: Accumulator]\tASL  Shift Left One Bit.\tA = A(%d) Shift Left 1 bit\t(%d)\n", opcode, A, A << 1 )
		}

		flags_C(A, A << 1)

		A = A << 1

		flags_N(A)
		flags_Z(A)

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		opc_cycle_count = 1
	}

}
