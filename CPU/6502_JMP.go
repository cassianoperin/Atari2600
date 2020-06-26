package CPU

import	"fmt"

// JMP  Jump to New Location (absolute)
//
//      (PC+1) -> PCL                    N Z C I D V
//      (PC+2) -> PCH                    - - - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      absolute      JMP oper      4C    3     3
func opc_JMP(memAddr uint16, mode string) {
	if Debug {
		fmt.Printf("\tOpcode %02X %02X%02X [3 bytes] [Mode: %s]\tJMP  Jump to New Location.\t\tPC = 0x%04X\n", Opcode, Memory[PC+2], Memory[PC+1], mode, memAddr)
	}
	PC = memAddr
	Beam_index += 3
}
