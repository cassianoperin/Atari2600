package VGS

import "fmt"

// EOR  Exclusive-OR Memory with Accumulator
//
//      A EOR M -> A                     N Z C I D V
//                                       + + - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      immediate     EOR #oper     49    2     2
//      zeropage      EOR oper      45    2     3
//      zeropage,X    EOR oper,X    55    2     4
//      absolute      EOR oper      4D    3     4
//      absolute,X    EOR oper,X    5D    3     4*
//      absolute,Y    EOR oper,Y    59    3     4*
//      (indirect,X)  EOR (oper,X)  41    2     6
//      (indirect),Y  EOR (oper),Y  51    2     5*

func opc_EOR(memAddr uint16, mode string, bytes uint16, opc_cycles byte) {

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

		// Print Opcode Debug Message
		opc_EOR_DebugMsg(bytes, mode, memAddr, memData)

		A = A ^ memData

		flags_Z(A)
		flags_N(A)

		// Increment PC
		PC += bytes

		// Reset Internal Opcode Cycle counters
		resetIntOpcCycleCounters()
	}
}

func opc_EOR_DebugMsg(bytes uint16, mode string, memAddr uint16, memData byte) {
	if Debug {
		opc_string := debug_decode_opc(bytes)
		dbg_show_message = fmt.Sprintf("\n\tOpcode %s [Mode: %s]\tEOR  Exclusive-OR Memory with Accumulator.\tA = A(%d) XOR Memory[0x%02X](%d)\t(%d)\n", opc_string, mode, A, memAddr, memData, A^memData)
		fmt.Println(dbg_show_message)
	}
}
