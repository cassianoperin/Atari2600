package VGS

import "fmt"

// INC  Increment Memory by One
//
//      M + 1 -> M                       N Z C I D V
//                                       + + - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      zeropage      INC oper      E6    2     5
//      zeropage,X    INC oper,X    F6    2     6
//      absolute      INC oper      EE    3     6
//      absolute,X    INC oper,X    FE    3     7

func opc_INC(memAddr uint16, mode string, bytes uint16, opc_cycles byte) {

	// Update Global Opc_cycles value
	Opc_cycles = opc_cycles

	// Print internal opcode cycle
	debugInternalOpcCycle(opc_cycles)

	// Just increment the Opcode cycle Counter
	if Opc_cycle_count < opc_cycles {
		Opc_cycle_count++

		// After spending the cycles needed, execute the opcode
	} else {

		// Read data from Memory (adress in Memory Bus) into Data Bus
		memData := dataBUS_Read(memAddr)

		// Print Opcode Debug Message
		opc_INC_DebugMsg(bytes, mode, memAddr, memData)

		// Write data to Memory (adress in Memory Bus) and update the value in Data BUS
		memData = dataBUS_Write(memAddr, memData+1)

		flags_Z(memData)
		flags_N(memData)

		// Increment PC
		PC += bytes

		// Reset Internal Opcode Cycle counters
		resetIntOpcCycleCounters()
	}
}

func opc_INC_DebugMsg(bytes uint16, mode string, memAddr uint16, memData byte) {
	if Debug {
		opc_string := debug_decode_opc(bytes)
		dbg_show_message = fmt.Sprintf("\n\tOpcode %s [Mode: %s]\tINC  Increment Memory[0x%02X](%d) by One (%d)\n", opc_string, mode, memAddr, memData, memData+1)
		fmt.Println(dbg_show_message)
	}
}
