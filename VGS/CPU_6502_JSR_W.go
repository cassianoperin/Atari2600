package VGS

import "fmt"

// JSR  Jump to New Location Saving Return Address
//
//      push (PC+2) to Stack,            N Z C I D V
//      (PC+1) -> PCL                    - - - - - -
//      (PC+2) -> PCH
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      absolute      JSR oper      20    3     6

func opc_JSR(memAddr uint16, mode string, bytes uint16, opc_cycles byte) {

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

		// Store the first byte into the Stack
		_ = dataBUS_Write(SP_Address, byte((PC+2)>>8)) // Write data to Memory (adress in Memory Bus) and update the value in Data BUS
		SP--
		SP_Address--

		// Store the second byte into the Stack
		_ = dataBUS_Write(SP_Address, byte((PC+2)&0xFF)) // Write data to Memory (adress in Memory Bus) and update the value in Data BUS
		SP_Address--
		SP--

		// Print Opcode Debug Message
		opc_JSR_DebugMsg(bytes, mode, memAddr, SP_Address)

		// Update PC
		PC = memAddr

		// Reset Internal Opcode Cycle counters
		resetIntOpcCycleCounters()
	}
}

func opc_JSR_DebugMsg(bytes uint16, mode string, memAddr uint16, SP_Address uint16) {
	if Debug {
		opc_string := debug_decode_opc(bytes)
		dbg_show_message = fmt.Sprintf("\n\tOpcode %s [Mode: %s]\tJSR  Jump to New Location Saving Return Address.\tPC = Memory[0x%02X]\t|\t Stack[0x%02X] = %02X\t Stack[0x%02X] = 0x%02X\n", opc_string, mode, memAddr, SP_Address+2, Memory[SP_Address+2], SP_Address+1, Memory[SP_Address+1])
		fmt.Println(dbg_show_message)
	}
}
