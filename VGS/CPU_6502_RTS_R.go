package VGS

import "fmt"

// RTS  Return from Subroutine
//
//      pull PC, PC+1 -> PC              N Z C I D V
//                                       - - - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      implied       RTS           60    1     6

func opc_RTS(bytes uint16, opc_cycles byte) {

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

		// Read data from Memory (adress in Memory Bus) into Data Bus
		memData_LSB := dataBUS_Read(SP_Address + 2)
		memData_MSB := dataBUS_Read(SP_Address + 1)

		PC = uint16(memData_LSB)<<8 | uint16(memData_MSB)

		// Update the Stack Pointer (Increase the two values retrieved)
		SP += 2

		// Print Opcode Debug Message
		opc_RTS_DebugMsg(bytes)

		// Increment PC
		PC += bytes

		// Reset Internal Opcode Cycle counters
		resetIntOpcCycleCounters()
	}
}

func opc_RTS_DebugMsg(bytes uint16) {
	if Debug {
		opc_string := debug_decode_opc(bytes)
		dbg_show_message = fmt.Sprintf("\n\tOpcode %s [Mode: Implied]\tRTS  Return from Subroutine.\tPC = 0x%04X (+ 1 RTS instruction byte) = 0x%04X\n", opc_string, PC, PC+0x01)
		fmt.Println(dbg_show_message)
	}
}
