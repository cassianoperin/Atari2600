package VGS

import "fmt"

// ROL  Rotate One Bit Left (Memory or Accumulator)
//
//      C <- [76543210] <- C             N Z C I D V
//                                       + + + - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      accumulator   ROL A         2A    1     2
//      zeropage      ROL oper      26    2     5
//      zeropage,X    ROL oper,X    36    2     6
//      absolute      ROL oper      2E    3     6
//      absolute,X    ROL oper,X    3E    3     7

//Move each of the bits in either A or M one place to the left.
//Bit 0 is filled with the current value of the carry flag whilst the old bit 7 becomes the new carry flag value.

// ------------------------------------ Accumulator ------------------------------------ //

func opc_ROL_A(bytes uint16, opc_cycles byte) {

	// Update Global Opc_cycles value
	Opc_cycles = opc_cycles

	// Print internal opcode cycle
	debugInternalOpcCycle(opc_cycles)

	// Just increment the Opcode cycle Counter
	if Opc_cycle_count < opc_cycles {
		Opc_cycle_count++

		// After spending the cycles needed, execute the opcode
	} else {

		// Original Carry Value
		carry_orig := P[0]

		// Print Opcode Debug Message
		opc_ROL_A_DebugMsg(bytes, carry_orig)

		flags_C(A & 0x80 >> 7) // Calculate the original bit7 and save it as the new Carry

		// Shift left the byte and put the original bit7 value in bit 1 to make the complete ROL
		A = (A << 1) + carry_orig

		flags_N(A)
		flags_Z(A)

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		Opc_cycle_count = 1
	}
}

func opc_ROL_A_DebugMsg(bytes uint16, carry_orig byte) {
	if Debug {
		opc_string := debug_decode_opc(bytes)
		dbg_show_message = fmt.Sprintf("\n\tOpcode %s [Mode: Accumulator]\tROL  Rotate One Bit Left.\tA(%d) Roll Left 1 bit + carry(%d)\t: %d\n", opc_string, A, P[0], (A<<1)+carry_orig)
		fmt.Println(dbg_show_message)
	}
}

// --------------------------------------- Memory -------------------------------------- //

func opc_ROL(memAddr uint16, mode string, bytes uint16, opc_cycles byte) {

	// Update Global Opc_cycles value
	Opc_cycles = opc_cycles

	// Print internal opcode cycle
	debugInternalOpcCycle(opc_cycles)

	// Just increment the Opcode cycle Counter
	if Opc_cycle_count < opc_cycles {
		Opc_cycle_count++

		// After spending the cycles needed, execute the opcode
	} else {

		// Original Carry Value
		carry_orig := P[0]

		// Read data from Memory (adress in Memory Bus) into Data Bus
		memData := dataBUS_Read(memAddr)

		// Print Opcode Debug Message
		opc_ROL_DebugMsg(bytes, mode, memAddr, carry_orig, memData)

		flags_C(memData & 0x80 >> 7) // Calculate the original bit7 and save it as the new Carry

		// Write data to Memory (adress in Memory Bus) and update the value in Data BUS
		memData = dataBUS_Write(memAddr, (memData<<1)+carry_orig)

		flags_N(memData)
		flags_Z(memData)

		// Increment PC
		PC += bytes

		// Reset Internal Opcode Cycle counters
		resetIntOpcCycleCounters()
	}
}

func opc_ROL_DebugMsg(bytes uint16, mode string, memAddr uint16, carry_orig byte, memData byte) {
	if Debug {
		opc_string := debug_decode_opc(bytes)
		dbg_show_message = fmt.Sprintf("\n\tOpcode %s [Mode: %s]\tROL  Rotate One Bit Left.\tMemory[0x%02X](%d) Roll Left 1 bit + Carry(%d)\t(%d)\n", opc_string, mode, memAddr, memData, carry_orig, (memData<<1)+carry_orig)
		fmt.Println(dbg_show_message)
	}
}
