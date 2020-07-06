package CPU

import	"fmt"

// NOP  No Operation
//
//      ---                              N Z C I D V
//                                       - - - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      implied       NOP           EA    1     2
func opc_NOP() {

	if Debug {
		fmt.Printf("\tOpcode %02X [1 byte] [Mode: Implied]\tNOP  No Operation. PC++\n", Opcode)
	}

	PC++
	Beam_index += 2
}
