package CPU

import	"fmt"
// import	"os"

// STA  Store Accumulator in Memory (zeropage,X)
//
//      A -> M                           N Z C I D V
//                                       - - - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      zeropage,X    STA oper,X    95    2     4
//      zeropage      STA oper      85    2     3
//      absolute,Y    STA oper,Y    99    3     5
func opc_STA(memAddr uint16, mode string) {

	Memory[ memAddr ] = A

	if Debug {
		// If mode = zeropage,X OR zeropage
		if Opcode == 0x95 || Opcode == 0x85{
			fmt.Printf("\n\tOpcode %02X%02X [2 bytes] [Mode: %s]\tSTA  Store Accumulator in Memory.\tMemory[%02X] = A (%d)\n", Opcode, Memory[PC+1], mode, memAddr, Memory[memAddr] )

		// If mode = absolute,Y
		} else if Opcode == 0x99 {
			fmt.Printf("\n\tOpcode %02X %02X%02X [3 bytes] [Mode: %s]\tSTA  Store Accumulator in Memory.\tMemory[%02X] = A (%d)\n", Opcode, Memory[PC+2], Memory[PC+1], mode, memAddr, Memory[memAddr] )
		}
	}

	// If mode = zeropage,X
	if Opcode == 0x95 {
		PC += 2
		Beam_index += 4

	// If mode = zeropage
	} else if Opcode == 0x85 {
		PC += 2
		Beam_index += 3

	// If mode = absolute,Y
	} else if Opcode == 0x99 {
		PC += 3
		Beam_index += 5
	}

	// Test STA (Store A in Memory), which address is being set to draw the element
	testAction(memAddr)

}