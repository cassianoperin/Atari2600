package CPU

import	"fmt"

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

	// If mode = zeropage,X
	if Opcode == 0x95 {
		Beam_index += 4
	// If mode = zeropage
	} else if Opcode == 0x85 {
	// If mode = absolute,Y
	} else if Opcode == 0x99 {
	}

	Memory[ memAddr ] = A

	if Debug {
		fmt.Printf("\n\tOpcode %02X%02X [2 bytes] [Mode: %s]\tSTA  Store Accumulator in Memory.\tMemory[%02X] = A (%d)\n", Opcode, Memory[PC+1], mode, memAddr, Memory[memAddr] )
	}

	// Every time we set something in memory must check which to performe the graphical action
	// Map the range to verify that to optimize?
	if Opcode == 0x95 || Opcode == 0x85 {
		testAction()

	// WORKAROUND TO FIX X MOVEMENT in 103 BOMBER
	} else {

		// Check Draw related memory addresses to send instructions to TV
		if Memory[PC+1] == WSYNC {
			if Debug {
				fmt.Printf("\nWSYNC SET\n")
			}
			DrawLine = true
			Beam_index = 0

			if Memory[GRP0] != 0 {
				if Debug {
					fmt.Printf("\nGRP0 SET\n")
				}
				DrawP0 = true
			}

			if Memory[GRP1] != 0 {
				if Debug {
					fmt.Printf("\nGRP1 SET\n")
				}
				DrawP1 = true
			}

		}

		if Memory[PC+1+uint16(Y)] == RESP0 {
			if Memory[RESP0] != 0 {
				XPositionP0 = Beam_index
				// if Debug {
					fmt.Printf("\nRESP0 SET\tXPositionP0: %d\tScreen Pos: %d\n", XPositionP0, (XPositionP0*3)-68)
					fmt.Printf("\nJetXPos: %d", Memory[0x80])
				// }
			}
		}

		if Memory[PC+1] == RESP1 {
			if Memory[RESP1] != 0 {
				XPositionP1 = Beam_index
				if Debug {
					fmt.Printf("\nRESP1 SET\tXPositionP1: %d\n", XPositionP1)
				}

			}
		}


		if Memory[PC+1+uint16(Y)] == HMP0 {
			XFinePositionP0 = Fine(Memory[HMP0])

			// if Debug {
				fmt.Printf("\nHMP0 SET: %d\n", XFinePositionP0)
			// }

		}

		if Memory[PC+1+uint16(Y)] == HMP1 {
			XFinePositionP1 = Fine(Memory[HMP1])
			if Debug {
				fmt.Printf("\nHMP1 SET: %d\n", XFinePositionP1)
			}
		}

	}

	// If mode = zeropage,X
	if Opcode == 0x95 {
		PC += 2
	// If mode = zeropage
	} else if Opcode == 0x85 {
		PC += 2
		Beam_index += 3

	// If mode = absolute,Y
	} else if Opcode == 0x99 {
		PC += 3
		Beam_index += 5

	}
}
