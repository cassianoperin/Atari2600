package VGS

import "fmt"

// Pull  Processor Status from Stack
//
//      pull SR                          N Z C I D V
//                                       from stack
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      implied       PLP           28    1     4

func opc_PLP(bytes uint16, opc_cycles byte) {

	// Update Global Opc_cycles value
	Opc_cycles = opc_cycles

	// Print internal opcode cycle
	debugInternalOpcCycle(opc_cycles)

	// Just increment the Opcode cycle Counter
	if Opc_cycle_count < opc_cycles {
		Opc_cycle_count++

		// After spending the cycles needed, execute the opcode
	} else {

		// 6502 handle Stack at the end of first memory page
		SP_Address := uint16(SP+1) + 256

		// Read data from Memory (adress in Memory Bus) into Data Bus
		memData := dataBUS_Read(SP_Address)

		// Turn the stack value into the processor status
		for i := 0; i < len(P); i++ {

			// The B Flag, PLP and RTI pull a byte from the stack and set all the flags. They ignore bits 5 and 4.
			if i == 4 || i == 5 {
				// P[i] = 1
				// Just ignore both
			} else {
				P[i] = (memData >> i) & 0x01
			}
		}

		// Print Opcode Debug Message
		opc_PLP_DebugMsg(bytes, SP_Address)

		SP++

		// Increment PC
		PC += bytes

		// Reset Internal Opcode Cycle counters
		resetIntOpcCycleCounters()
	}
}

func opc_PLP_DebugMsg(bytes uint16, SP_Address uint16) {
	if Debug {
		opc_string := debug_decode_opc(bytes)
		dbg_show_message = fmt.Sprintf("\n\tOpcode %s [Mode: Implied]\tPLP  Processor Status from Stack.\tP = Memory[0x%02X] %d | SP++\n", opc_string, SP_Address, P)
		fmt.Println(dbg_show_message)
	}
}
