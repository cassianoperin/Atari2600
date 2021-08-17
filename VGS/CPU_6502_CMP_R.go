package VGS

import "fmt"

// CMP  Compare Memory with Accumulator
//
//      A - M                          N Z C I D V
//                                     + + + - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      immediate     CMP #oper     C9    2     2
//      zeropage      CMP oper      C5    2     3
//      zeropage,X    CMP oper,X    D5    2     4
//      absolute      CMP oper      CD    3     4
//      absolute,X    CMP oper,X    DD    3     4*
//      absolute,Y    CMP oper,Y    D9    3     4*
//      (indirect,X)  CMP (oper,X)  C1    2     6
//      (indirect),Y  CMP (oper),Y  D1    2     5*

func opc_CMP(memAddr uint16, mode string, bytes uint16, opc_cycles byte) {

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

		tmp := A - memData

		// Print Opcode Debug Message
		opc_CMP_DebugMsg(bytes, tmp, mode, memAddr, memData)

		flags_Z(tmp)
		flags_N(tmp)
		flags_C_CPX_CPY_CMP(A, memData) // Set if A >= M

		// Increment PC
		PC += bytes

		// Reset Internal Opcode Cycle counters
		resetIntOpcCycleCounters()
	}

}

func opc_CMP_DebugMsg(bytes uint16, tmp byte, mode string, memAddr uint16, memData byte) {
	if Debug {
		opc_string := debug_decode_opc(bytes)
		if tmp == 0 {
			dbg_show_message = fmt.Sprintf("\n\tOpcode %s [Mode: %s]\tCMP  Compare Memory with Accumulator.\tA(%d) - Memory[0x%02X](%d) = (%d) EQUAL\n", opc_string, mode, A, memAddr, memData, tmp)
		} else {
			dbg_show_message = fmt.Sprintf("\n\tOpcode %s [Mode: %s]\tCMP  Compare Memory with Accumulator.\tA(%d) - Memory[0x%02X](%d) = (%d) NOT EQUAL\n", opc_string, mode, A, memAddr, memData, tmp)
		}
		fmt.Println(dbg_show_message)
	}
}
