package VGS

import "fmt"

// LDX  Load Index X with Memory
//
//      M -> X                           N Z C I D V
//                                       + + - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      immediate     LDX #oper     A2     2     2
//      zeropage	  LDX oper	    A6     2     3
//      zeropage,Y    LDX oper,Y    B6     2     4
//      absolute      LDX oper      AE     3     4
//      absolute,Y    LDX oper,Y    BE     3     4*

func opc_LDX(memAddr uint16, mode string, bytes uint16, opc_cycles byte) {

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

		X = memData

		// Print Opcode Debug Message
		opc_LDX_DebugMsg(bytes, mode, memAddr)

		flags_Z(X)
		flags_N(X)

		// Increment PC
		PC += bytes

		// Reset Internal Opcode Cycle counters
		resetIntOpcCycleCounters()
	}
}

func opc_LDX_DebugMsg(bytes uint16, mode string, memAddr uint16) {
	if Debug {
		opc_string := debug_decode_opc(bytes)
		dbg_show_message = fmt.Sprintf("\n\tOpcode %s [Mode: %s]\tLDX  Load Index X with Memory.\tX = Memory[0x%02X] (%d)\n", opc_string, mode, memAddr, X)
		fmt.Println(dbg_show_message)
	}
}
