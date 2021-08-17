package VGS

import "fmt"

// PLA  Pull Accumulator from Stack
//
//      pull A                           N Z C I D V
//                                       + + - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      implied       PLA           68    1     4

func opc_PLA(bytes uint16, opc_cycles byte) {

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

		A = memData

		// Print Opcode Debug Message
		opc_PLA_DebugMsg(bytes, SP_Address)

		flags_N(A)
		flags_Z(A)

		SP++

		// Increment PC
		PC += bytes

		// Reset Internal Opcode Cycle counters
		resetIntOpcCycleCounters()
	}
}

func opc_PLA_DebugMsg(bytes uint16, SP_Address uint16) {
	if Debug {
		opc_string := debug_decode_opc(bytes)
		dbg_show_message = fmt.Sprintf("\n\tOpcode %s [Mode: Implied]\tPLA  Pull Accumulator from Stack.\tA = Memory[0x%02X] (%d) | SP++\n", opc_string, SP_Address, A)
		fmt.Println(dbg_show_message)
	}
}
