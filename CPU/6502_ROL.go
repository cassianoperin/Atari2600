package CPU

import	"fmt"

// ROL  Rotate One Bit Left (Memory or Accumulator)
//
//      C <- [76543210] <- C             N Z C I D V
//                                       + + + - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      zeropage      ROL oper      26    2     5
func opc_ROL(memAddr uint16, mode string) {
	// Original Carry Value
	carry_orig := P[0]

	if Debug {
		fmt.Printf("\tOpcode %02X%02X [2 bytes] [Mode: %s]\tROL  Rotate One Bit Left.\tMemory[%d] Roll Left 1 bit\t(%d)\n", Opcode, Memory[PC+1], mode, memAddr, ( Memory[memAddr] << 1 ) + carry_orig )
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

	PC += 2
	Beam_index += 5
}
