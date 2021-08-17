package VGS

import "fmt"

// TAX  Transfer Accumulator to Index X
//
//      A -> X                           N Z C I D V
//                                       + + - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      implied       TAX           AA    1     2

func opc_TAX(bytes uint16, opc_cycles byte) {

	// Update Global Opc_cycles value
	Opc_cycles = opc_cycles

	// Print internal opcode cycle
	debugInternalOpcCycle(opc_cycles)

	// Just increment the Opcode cycle Counter
	if Opc_cycle_count < opc_cycles {
		Opc_cycle_count++

		// After spending the cycles needed, execute the opcode
	} else {

		X = A

		// Print Opcode Debug Message
		opc_TAX_DebugMsg(bytes)

		flags_Z(X)
		flags_N(X)

		// Increment PC
		PC += bytes

		// Reset Internal Opcode Cycle counters
		resetIntOpcCycleCounters()
	}
}

func opc_TAX_DebugMsg(bytes uint16) {
	if Debug {
		opc_string := debug_decode_opc(bytes)
		dbg_show_message = fmt.Sprintf("\n\tOpcode %s [Mode: Implied]\tTAX  Transfer Accumulator to Index X.\tX = A (%d)\n", opc_string, A)
		fmt.Println(dbg_show_message)
	}
}
