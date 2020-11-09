package VGS

import	"fmt"

// ADC  Add Memory to Accumulator with Carry (zeropage)
//
// 	A + M + C -> A, C                N Z C I D V
// 	                          	   + + + - - +
//
// 	addressing    assembler    opc  bytes  cyles
// 	--------------------------------------------
// 	zeropage      ADC oper      65    2     3
func opc_ADC(memAddr uint16, mode string, bytes uint16, opc_cycles byte) {

	// Increment the beam
	beamIndex ++

	// // Check for extra cycles (*) in the first opcode cycle
	// if opc_cycle_count == 1 {
	// 	if Opcode == 0xB9 || Opcode == 0xBD || Opcode == 0xB1 {
	// 		// Add 1 to cycles if page boundery is crossed
	// 		if MemPageBoundary(memAddr, PC) {
	// 			opc_cycle_extra = 1
	// 		}
	// 	}
	// }

	// Show current opcode cycle
	if Debug {
		fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\t(%d cycles + %d extra cycles)\n", Cycle, opc_cycle_count, opc_cycles + opc_cycle_extra, opc_cycles, opc_cycle_extra)
	}

	// Just increment the Opcode cycle Counter
	if opc_cycle_count < opc_cycles +  opc_cycle_extra {
		opc_cycle_count ++

	// After spending the cycles needed, execute the opcode
	} else {

		// Original value of A
		tmp := A

		if Debug {
			fmt.Printf("\tOpcode %02X%02X [2 bytes] [Mode: %s]\tADC  Add Memory to Accumulator with Carry (zeropage).\tA = A(%d) + Memory[Memory[%02X]](%d) + Carry (%d)) = %d\n", opcode, Memory[PC+1], mode, A, PC+1, Memory[Memory[PC+1]], P[0] , A + Memory[Memory[PC+1]] + P[0] )
		}

		// Result
		A = A + Memory[Memory[PC+1]] + P[0]

		// For the flags:
		// The addiction is VALUE1 (A) - VALUE2 (Memory[Memory[PC+1]] + P[0])
		// value2 := Memory[Memory[PC+1]] + P[0]

		// First V because it need the original carry flag value
		Flags_V_ADC(tmp, A)
		// After, update the carry flag value
		flags_C(tmp, A)

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
