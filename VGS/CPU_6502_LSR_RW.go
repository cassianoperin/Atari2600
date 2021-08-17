package VGS

import "fmt"

// LSR  Shift One Bit Right (Memory or Accumulator)
//
//      0 -> [76543210] -> C             N Z C I D V
//                                       0 + + - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      accumulator   LSR A         4A    1     2
//      zeropage      LSR oper      46    2     5
//      zeropage,X    LSR oper,X    56    2     6
//      absolute      LSR oper      4E    3     6
//      absolute,X    LSR oper,X    5E    3     7

// ------------------------------------ Accumulator ------------------------------------ //

func opc_LSR_A(bytes uint16, opc_cycles byte) {

	// Update Global Opc_cycles value
	Opc_cycles = opc_cycles

	// Print internal opcode cycle
	debugInternalOpcCycle(opc_cycles)

	// Just increment the Opcode cycle Counter
	if Opc_cycle_count < opc_cycles {
		Opc_cycle_count++

		// After spending the cycles needed, execute the opcode
	} else {

		// Print Opcode Debug Message
		opc_LSR_A_DebugMsg(bytes)

		flags_C(A & 0x01) // Least significant bit turns into the new Carry

		A = A >> 1

		flags_N(A)
		flags_Z(A)

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		Opc_cycle_count = 1
	}

}

func opc_LSR_A_DebugMsg(bytes uint16) {
	if Debug {
		opc_string := debug_decode_opc(bytes)
		dbg_show_message = fmt.Sprintf("\n\tOpcode %s [Mode: Accumulator]\tLSR  Shift One Bit Right.\tA = A(%d) Shift Right 1 bit\t(%d)\n", opc_string, A, A>>1)
		fmt.Println(dbg_show_message)
	}
}

// --------------------------------------- Memory -------------------------------------- //

func opc_LSR(memAddr uint16, mode string, bytes uint16, opc_cycles byte) {

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
		opc_LSR_DebugMsg(bytes, mode, memAddr, memData)

		flags_C(memData & 0x01) // Least significant bit turns into the new Carry

		// Write data to Memory (adress in Memory Bus) and update the value in Data BUS
		memData = dataBUS_Write(memAddr, memData>>1)

		flags_N(memData)
		flags_Z(memData)

		// Increment PC
		PC += bytes

		// Reset Internal Opcode Cycle counters
		resetIntOpcCycleCounters()
	}
}

func opc_LSR_DebugMsg(bytes uint16, mode string, memAddr uint16, memData byte) {
	if Debug {
		opc_string := debug_decode_opc(bytes)
		dbg_show_message = fmt.Sprintf("\n\tOpcode %s [Mode: %s]\tLSR  Shift One Bit Right.\tMemory[0x%02X]: (%d) Shift Right 1 bit\t(%d)\n", opc_string, mode, memAddr, memData, memData>>1)
		fmt.Println(dbg_show_message)
	}
}
