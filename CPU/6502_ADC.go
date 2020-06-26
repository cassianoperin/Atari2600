package CPU

import	"fmt"

// ADC  Add Memory to Accumulator with Carry (zeropage)
//
// 	A + M + C -> A, C                N Z C I D V
// 	                          	   + + + - - +
//
// 	addressing    assembler    opc  bytes  cyles
// 	--------------------------------------------
// 	zeropage      ADC oper      65    2     3
func opc_ADC(memAddr uint16, mode string) {

	// Original value of A
	tmp := A

	if Debug {
		fmt.Printf("\tOpcode %02X%02X [2 bytes] [Mode: %s]\tADC  Add Memory to Accumulator with Carry (zeropage).\tA = A(%d) + Memory[Memory[%02X]](%d) + Carry (%d)) = %d\n", Opcode, Memory[PC+1], mode, A, PC+1, Memory[Memory[PC+1]], P[0] , A + Memory[Memory[PC+1]] + P[0] )
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

	PC += 2
	Beam_index += 3
}
