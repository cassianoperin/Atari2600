package VGS

import "fmt"

// CPX  Compare Memory and Index X
//
//      X - M                            N Z C I D V
//                                       + + + - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      immediate     CPX #oper     E0    2     2
//      zeropage      CPX oper    	E4    2	    3
//      absolute      CPX oper      EC    3     4

func opc_CPX(memAddr uint16, mode string, bytes uint16, opc_cycles byte) {

	// Update Global Opc_cycles value
	Opc_cycles = opc_cycles

	// Print internal opcode cycle
	debugInternalOpcCycle(opc_cycles)

	if Opc_cycle_count < opc_cycles { // Just increment the Opcode cycle Counter
		Opc_cycle_count++

	} else { // After spending the cycles needed, execute the opcode

		// Read data from Memory (adress in Memory Bus) into Data Bus
		memData := dataBUS_Read(memAddr)

		tmp := X - memData

		// Print Opcode Debug Message
		opc_CPX_DebugMsg(bytes, tmp, mode, memAddr, memData)

		flags_Z(tmp)                    // Set if X = M
		flags_N(tmp)                    // Set if bit 7 of the result is set
		flags_C_CPX_CPY_CMP(X, memData) // Set if X >= M

		// Increment PC
		PC += bytes

		// Reset Internal Opcode Cycle counters
		resetIntOpcCycleCounters()
	}
}

func opc_CPX_DebugMsg(bytes uint16, tmp byte, mode string, memAddr uint16, memData byte) {
	if Debug {
		// Print Opcode Debug Message
		opc_string := debug_decode_opc(bytes)
		if tmp == 0 {
			dbg_show_message = fmt.Sprintf("\n\tOpcode %s [Mode: %s]\tCPX  Compare Memory and Index X.\tX(%d) - Memory[0x%02X](%d) = (%d) EQUAL\n", opc_string, mode, X, PC+1, memData, tmp)
		} else {
			dbg_show_message = fmt.Sprintf("\n\tOpcode %s [Mode: %s]\tCPX  Compare Memory and Index X.\tX(%d) - Memory[0x%02X](%d) = (%d) NOT EQUAL\n", opc_string, mode, X, PC+1, memData, tmp)
		}
		fmt.Println(dbg_show_message)
	}
}
