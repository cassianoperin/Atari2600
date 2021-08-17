package VGS

import "fmt"

// LDA  Load Accumulator with Memory
//
//      M -> A                           N Z C I D V
//                                       + + - - - -
//
//      addressing    assembler    opc  bytes  cycles
//      --------------------------------------------
//      immediate     LDA #oper     A9    2     2
//      zeropage      LDA oper      A5    2     3
//      zeropage,X    LDA oper,X    B5    2     4
//      absolute      LDA oper      AD    3     4
//      absolute,X    LDA oper,X    BD    3     4*
//      absolute,Y    LDA oper,Y    B9    3     4*
//      (indirect,X)  LDA (oper,X)  A1    2     6
//      (indirect),Y  LDA (oper),Y  B1    2     5*

func opc_LDA(memAddr uint16, mode string, bytes uint16, opc_cycles byte) {

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

		A = memData

		// Print Opcode Debug Message
		opc_LDA_DebugMsg(bytes, mode, memAddr)

		flags_Z(A)
		flags_N(A)

		// Increment PC
		PC += bytes

		// Reset Internal Opcode Cycle counters
		resetIntOpcCycleCounters()
	}
}

func opc_LDA_DebugMsg(bytes uint16, mode string, memAddr uint16) {
	if Debug {
		opc_string := debug_decode_opc(bytes)
		dbg_show_message = fmt.Sprintf("\n\tOpcode %s [Mode: %s]\tLDA  Load Accumulator with Memory.\tA = Memory[0x%02X] (%d)\n", opc_string, mode, memAddr, A)
		fmt.Println(dbg_show_message)
	}
}
