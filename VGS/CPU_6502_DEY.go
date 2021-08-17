package VGS

import "fmt"

// DEY  Decrement Index Y by One
//
//      Y - 1 -> Y                       N Z C I D V
//                                       + + - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      implied       DEC           88    1     2

func opc_DEY(bytes uint16, opc_cycles byte) {

	// Update Global Opc_cycles value
	Opc_cycles = opc_cycles

	// Print internal opcode cycle
	debugInternalOpcCycle(opc_cycles)

	// Just increment the Opcode cycle Counter
	if Opc_cycle_count < opc_cycles {
		Opc_cycle_count++

		// After spending the cycles needed, execute the opcode
	} else {

		Y--

		// Print Opcode Debug Message
		opc_DEY_DebugMsg(bytes)

		flags_Z(Y)
		flags_N(Y)

		// Increment PC
		PC += bytes

		// Reset Internal Opcode Cycle counters
		resetIntOpcCycleCounters()
	}
}

func opc_DEY_DebugMsg(bytes uint16) {
	if Debug {
		opc_string := debug_decode_opc(bytes)
		dbg_show_message = fmt.Sprintf("\n\tOpcode %s [Mode: Implied]\tDEY  Decrement Index Y by One.\tY-- (%d)\n", opc_string, Y)
		fmt.Println(dbg_show_message)
	}
}
