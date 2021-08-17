package VGS

import "fmt"

// PHA  Push Accumulator on Stack
//
//      push A                           N Z C I D V
//                                       - - - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      implied       PHA           48    1     3

func opc_PHA(bytes uint16, opc_cycles byte) {

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
		SP_Address := uint16(SP) + 256

		// Write data to Memory (adress in Memory Bus) and update the value in Data BUS
		memData := dataBUS_Write(SP_Address, A)

		// Print Opcode Debug Message
		opc_PHA_DebugMsg(bytes, SP_Address, memData)

		SP--

		// Increment PC
		PC += bytes

		// Reset Internal Opcode Cycle counters
		resetIntOpcCycleCounters()
	}
}

func opc_PHA_DebugMsg(bytes uint16, SP_Address uint16, memData byte) {
	if Debug {
		opc_string := debug_decode_opc(bytes)
		dbg_show_message = fmt.Sprintf("\n\tOpcode %s [Mode: Implied]\tPHA  Push Accumulator on Stack.\tMemory[0x%02X] = A (%d) | SP--\n", opc_string, SP_Address, memData)
		fmt.Println(dbg_show_message)
	}
}
