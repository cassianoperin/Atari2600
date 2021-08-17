package VGS

import "fmt"

// CLV  Clear Overflow Flag
//
//      0 -> V                           N Z C I D V
//                                       - - - - - 0
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      implied       CLV           B8     1    2

func opc_CLV(bytes uint16, opc_cycles byte) {

	// Update Global Opc_cycles value
	Opc_cycles = opc_cycles

	// Print internal opcode cycle
	debugInternalOpcCycle(opc_cycles)

	// Just increment the Opcode cycle Counter
	if Opc_cycle_count < opc_cycles {
		Opc_cycle_count++

		// After spending the cycles needed, execute the opcode
	} else {

		P[6] = 0

		// Print Opcode Debug Message
		opc_CLV_DebugMsg(bytes)

		// Increment PC
		PC += bytes

		// Reset Internal Opcode Cycle counters
		resetIntOpcCycleCounters()
	}
}

func opc_CLV_DebugMsg(bytes uint16) {
	if Debug {
		opc_string := debug_decode_opc(bytes)
		dbg_show_message = fmt.Sprintf("\n\tOpcode %s [Mode: Implied]\tCLV  Clear Overflow Flag.\tP[6]=%d\n", opc_string, P[6])
		fmt.Println(dbg_show_message)
	}
}
