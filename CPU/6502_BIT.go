package CPU

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
//      absolute      BIT oper      2C    3     4
func opc_BIT(memAddr uint16, mode string) {
	if Debug {
		fmt.Printf("\n\tOpcode %02X %02X%02X [3 bytes] [Mode: %s]\tBIT  Test Bits in Memory with Accumulator.\tA (%08b) AND Memory[%04X] (%08b) = %08b \tM7 -> N, M6 -> V\n", Opcode, Memory[PC+2], Memory[PC+1], mode, A, memAddr, Memory[memAddr], A & Memory[memAddr] )
	}
	// fmt.Printf("\n\n%08b\n\n",A & Memory[memAddr])

	// Memory Address bit 7 (A) -> N (Negative)
	if Debug {
		fmt.Printf("\n\tFlag N: %d -> ",P[7])
	}
	P[7] = A >> 7 & 0x1
	if Debug {
		fmt.Printf("%d\n",P[7])
	}

	// Memory Address bit 6 (A) -> V (oVerflow)
	if Debug {
		fmt.Printf("\n\tFlag V: %d -> ",P[6])
	}
	P[6] = A >> 6 & 0x1
	if Debug {
		fmt.Printf("%d\n",P[6])
	}

	flags_Z(A & Memory[memAddr])

	PC += 3
	Beam_index += 4
}
