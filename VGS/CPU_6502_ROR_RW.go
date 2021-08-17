package VGS

import "fmt"

// ROR  Rotate One Bit Right (Memory or Accumulator)
//
//      C -> [76543210] -> C             N Z C I D V
//                                       + + + - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      accumulator      ROR A      6A    1     2
//      zeropage         ROR oper   66    2     5
//      zeropage,X       ROR oper,X 76    2     6
//      absolute         ROR oper   6E    3     6
//      absolute,X       ROR oper,X 7E    3     7

// Move each of the bits in either A or M one place to the right.
// Bit 7 is filled with the current value of the carry flag whilst the old bit 0 becomes the new carry flag value.

// ------------------------------------ Accumulator ------------------------------------ //

func opc_ROR_A(bytes uint16, opc_cycles byte) {

	// Update Global Opc_cycles value
	Opc_cycles = opc_cycles

	// Print internal opcode cycle
	debugInternalOpcCycle(opc_cycles)

	// Just increment the Opcode cycle Counter
	if Opc_cycle_count < opc_cycles {
		Opc_cycle_count++

		// After spending the cycles needed, execute the opcode
	} else {

		// Keep original Accumulator value for debug
		original_A := A
		original_carry := P[0]

		// Keep the original bit 0 from Accumulator to be used as new Carry
		new_Carry := A & 0x01

		// Shift Right Accumulator
		A = A >> 1

		// Bit 7 is filled with the current value of the carry flag
		A += (P[0] << 7)

		// Print Opcode Debug Message
		opc_ROR_A_DebugMsg(bytes, original_A, original_carry)

		flags_C(new_Carry) // The old bit 0 becomes the new carry flag value
		flags_N(A)
		flags_Z(A)

		// Increment PC
		PC += bytes

		// Reset Opcode Cycle counter
		Opc_cycle_count = 1
	}

}

func opc_ROR_A_DebugMsg(bytes uint16, original_A byte, original_carry byte) {
	if Debug {
		opc_string := debug_decode_opc(bytes)
		dbg_show_message = fmt.Sprintf("\n\tOpcode %s [Mode: Accumulator]\tROR  Rotate One Bit Right.\tA(%d) Roll Right 1 bit\t(%d) + Current Carry(%d) as new bit 7.\tA = %d\n", opc_string, original_A, original_A>>1, original_carry, A)
		fmt.Println(dbg_show_message)
	}
}

// --------------------------------------- Memory -------------------------------------- //

func opc_ROR(memAddr uint16, mode string, bytes uint16, opc_cycles byte) {

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

		// Keep original Accumulator value for debug
		original_MemValue := memData
		original_carry := P[0]

		// Keep the original bit 0 from Accumulator to be used as new Carry
		new_Carry := memData & 0x01

		// Write data to Memory (adress in Memory Bus) and update the value in Data BUS
		// Shift Right Memory Value
		memData = dataBUS_Write(memAddr, memData>>1)
		// Bit 7 is filled with the current value of the carry flag
		memData = dataBUS_Write(memAddr, memData+(P[0]<<7))

		// Print Opcode Debug Message
		opc_ROR_DebugMsg(bytes, mode, memAddr, original_MemValue, original_carry, memData)

		flags_C(new_Carry) // The old bit 0 becomes the new carry flag value
		flags_N(memData)
		flags_Z(memData)

		// Increment PC
		PC += bytes

		// Reset Internal Opcode Cycle counters
		resetIntOpcCycleCounters()
	}
}

func opc_ROR_DebugMsg(bytes uint16, mode string, memAddr uint16, original_MemValue byte, original_carry byte, memData byte) {
	if Debug {
		opc_string := debug_decode_opc(bytes)
		dbg_show_message = fmt.Sprintf("\n\tOpcode %s [Mode: %s]\tROR  Rotate One Bit Right.\tMemory[0x%02d](%d) Roll Right 1 bit\t(%d) + Current Carry(%d) as new bit 7.\tA = %d\n", opc_string, mode, memAddr, original_MemValue, original_MemValue>>1, original_carry, memData)
		fmt.Println(dbg_show_message)
	}
}
