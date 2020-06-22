package CPU

import	"fmt"

// JSR  Jump to New Location Saving Return Address
//
//      push (PC+2) to Stack,            N Z C I D V
//      (PC+1) -> PCL                    - - - - - -
//      (PC+2) -> PCH
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      absolute      JSR oper      20    3     6
func opc_JSR(memAddr uint16, mode string) {
	
	// Push PC+2 (will be increased in 1 in RTS to match the next address (3 bytes operation))
	// Store the first byte into the Stack
	Memory[SP] = byte( (PC+2) >> 8 )
	SP--
	// Store the second byte into the Stack
	Memory[SP] = byte( (PC+2) & 0xFF )
	SP--
	// fmt.Printf("\nPC+3: %02X",PC+3)
	// fmt.Printf("\nF0: %02X",(PC+3) >> 8)
	// fmt.Printf("\n42: %02X",(PC+3) & 0xFF)

	if Debug {
		fmt.Printf("\n\tOpcode %02X %02X%02X [3 bytes] [Mode: %s]\tJSR  Jump to New Location Saving Return Address.\tPC = Memory[%02X]\t|\t Stack[%02X] = %02X\t Stack[%02X] = %02X\n", Opcode, Memory[PC+2], Memory[PC+1], mode, memAddr, SP+2, Memory[SP+2], SP+1, Memory[SP+1])
	}

	PC = memAddr
	Beam_index += 6
}
