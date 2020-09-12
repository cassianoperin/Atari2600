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
//      absolute      STA oper      8D    3     4
func opc_STA(memAddr uint16, mode string) {

	// Just increment the Beam (To let TIA Draw with the cycles used first)
	if !TIA_CPU_HOLD {

		// If mode = zeropage,X
		if Opcode == 0x95 {
			Beam_index += 4

		// If mode = zeropage
		} else if Opcode == 0x85 {
			Beam_index += 3

		// If mode = absolute,Y
		} else if Opcode == 0x99 {
			Beam_index += 5

		// If mode = absolute
		} else if Opcode == 0x8D {
			Beam_index += 4
		}

		// Execute the operation and increment PC
		} else {
		if memAddr < 128 {
			TIA_Update = int8(memAddr)	// Change variable to a positive number to TIA detect the change
		}

		Memory[ memAddr ] = A

		if Debug {
			// If mode = zeropage,X OR zeropage
			if Opcode == 0x95 || Opcode == 0x85 {
				fmt.Printf("\tOpcode %02X%02X [2 bytes] [Mode: %s]\tSTA  Store Accumulator in Memory.\tMemory[%02X] = A (%d)\n", Opcode, Memory[PC+1], mode, memAddr, Memory[memAddr] )

			// If mode = absolute,Y
			} else if Opcode == 0x99 || Opcode == 0x8D {
				fmt.Printf("\tOpcode %02X %02X%02X [3 bytes] [Mode: %s]\tSTA  Store Accumulator in Memory.\tMemory[%02X] = A (%d)\n", Opcode, Memory[PC+2], Memory[PC+1], mode, memAddr, Memory[memAddr] )
			}
		}

		// If mode = zeropage,X
		if Opcode == 0x95 {
			PC += 2

		// If mode = zeropage
		} else if Opcode == 0x85 {
			PC += 2

		// If mode = absolute,Y
		} else if Opcode == 0x99 {
			PC += 3

		// If mode = absolute
		} else if Opcode == 0x8D {
			PC += 3
		}
	}

	TIA_CPU_HOLD = !TIA_CPU_HOLD
}
