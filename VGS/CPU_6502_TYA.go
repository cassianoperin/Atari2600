package VGS

import "fmt"

// TYA Transfer Index Y to Accumulator
//
//     Y -> A                            N Z C I D V
//                                       + + - - - -
//
//     addressing	   assembler    opc  bytes  cyles
//     --------------------------------------------
//     implied	     TYA	         98    1     2

func opc_TYA(bytes uint16, opc_cycles byte) {

	// Update Global Opc_cycles value
	Opc_cycles = opc_cycles

	// Print internal opcode cycle
	debugInternalOpcCycle(opc_cycles)

	// Just increment the Opcode cycle Counter
	if Opc_cycle_count < opc_cycles {
		Opc_cycle_count++

		// After spending the cycles needed, execute the opcode
	} else {

		A = Y

		// Print Opcode Debug Message
		opc_TYA_DebugMsg(bytes)

		flags_Z(A)
		flags_N(A)

		// Increment PC
		PC += bytes

		// Reset Internal Opcode Cycle counters
		resetIntOpcCycleCounters()
	}
}

func opc_TYA_DebugMsg(bytes uint16) {
	if Debug {
		opc_string := debug_decode_opc(bytes)
		dbg_show_message = fmt.Sprintf("\n\tOpcode %s [Mode: Implied]\tTYA  Transfer Index Y to Accumulator.\tA = Y (%d)\n", opc_string, Y)
		fmt.Println(dbg_show_message)
	}
}
