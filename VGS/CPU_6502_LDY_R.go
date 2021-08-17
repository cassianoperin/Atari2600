package VGS

import "fmt"

// LDY  Load Index Y with Memory (immediate)
//
//      M -> Y                           N Z C I D V
//                                       + + - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      immediate     LDY #oper     A0    2     2
//      zeropage      LDY oper      A4    2     3
//      zeropage,X    LDY oper,X    B4    2     4
//      absolute      LDY oper      AC    3     4
//      absolute,X    LDY oper,X    BC    3     4*

func opc_LDY(memAddr uint16, mode string, bytes uint16, opc_cycles byte) {

	// Update Global Opc_cycles value
	Opc_cycles = opc_cycles

	// Print internal opcode cycle
	debugInternalOpcCycleExtras(opc_cycles)

	// Just increment the Opcode cycle Counter
	if Opc_cycle_count < opc_cycles+Opc_cycle_extra {
		Opc_cycle_count++

		// After spending the cycles needed, execute the opcode
	} else {

		// Read data from Memory (adress in Memory Bus) into Data Bus
		memData := dataBUS_Read(memAddr)

		Y = memData

		// Print Opcode Debug Message
		opc_LDY_DebugMsg(bytes, mode, memAddr)

		flags_Z(Y)
		flags_N(Y)

		// Increment PC
		PC += bytes

		// Reset Internal Opcode Cycle counters
		resetIntOpcCycleCounters()
	}
}

func opc_LDY_DebugMsg(bytes uint16, mode string, memAddr uint16) {
	if Debug {
		opc_string := debug_decode_opc(bytes)
		dbg_show_message = fmt.Sprintf("\n\tOpcode %s [Mode: %s]\tLDY  Index Y with Memory.\tY = Memory[0x%02X] (%d)\n", opc_string, mode, memAddr, Y)
		fmt.Println(dbg_show_message)
	}
}
