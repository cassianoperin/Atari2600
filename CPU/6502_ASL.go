package CPU

import	"fmt"

// ASL  Shift Left One Bit (Memory or Accumulator) (accumulator)
//
//      C <- [76543210] <- 0             N Z C I D V
//                                       + + + - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      accumulator   ASL A         0A    1     2
func opc_ASL() {
	if Debug {
		fmt.Printf("\n\tOpcode %02X [1 byte] [Mode: Accumulator]\tASL  Shift Left One Bit (Memory or Accumulator).\tA = A(%d) Shift Left 1 bit\t(%d)\n", Opcode, A, A << 1 )
	}

	flags_C(A, A << 1)

	A = A << 1

	flags_N(A)
	flags_Z(A)

	PC += 1
	Beam_index += 2
}
