package VGS

import "fmt"

// STA  Store Accumulator in Memory
//
//      A -> M                           N Z C I D V
//                                       - - - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      zeropage      STA oper      85    2     3
//      zeropage,X    STA oper,X    95    2     4
//      absolute      STA oper      8D    3     4
//      absolute,X    STA oper,X    9D    3     5
//      absolute,Y    STA oper,Y    99    3     5
//      (indirect,X)  STA (oper,X)  81    2     6
//      (indirect),Y  STA (oper),Y  91    2     6

func opc_STA(memAddr uint16, mode string, bytes uint16, opc_cycles byte) {

	// Update Global Opc_cycles value
	Opc_cycles = opc_cycles

	// Print internal opcode cycle
	debugInternalOpcCycle(opc_cycles)

	// Just increment the Opcode cycle Counter
	if Opc_cycle_count < opc_cycles {
		Opc_cycle_count++

		// TIA_Update = -1

		// After spending the cycles needed, execute the opcode
	} else {

		// Write data to Memory (adress in Memory Bus) and update the value in Data BUS
		memData := dataBUS_Write(memAddr, A)

		// Print Opcode Debug Message
		opc_STA_DebugMsg(bytes, mode, memAddr, memData)

		// Increment PC
		PC += bytes

		// Reset Internal Opcode Cycle counters
		resetIntOpcCycleCounters()
	}
}

func opc_STA_DebugMsg(bytes uint16, mode string, memAddr uint16, memData byte) {
	if Debug {
		opc_string := debug_decode_opc(bytes)
		dbg_show_message = fmt.Sprintf("\n\tOpcode %s [Mode: %s]\tSTA  Store Accumulator in Memory.\tMemory[0x%02X] = A (0x%02X)\n", opc_string, mode, memAddr, memData)
		fmt.Println(dbg_show_message)
	}
}
