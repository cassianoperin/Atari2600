package VGS

import "fmt"

// BNE  Branch on Result not Zero
//
//      branch on Z = 0                  N Z C I D V
//                                       - - - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      relative      BNE oper      D0    2     2**

func opc_BNE(memAddr uint16, bytes uint16, opc_cycles byte) {

	// Update Global Opc_cycles value
	Opc_cycles = opc_cycles

	// Update Global Opc_cycles value
	Opc_cycles = opc_cycles

	// Read data from Memory (adress in Memory Bus) into Data Bus
	memData := dataBUS_Read(memAddr)

	// Get the Two's complement value of value in Memory
	value := DecodeTwoComplement(memData) // value is SIGNED

	if P[1] == 1 { // If P[1] = 1 (Zero Flag)

		// Print internal opcode cycle
		debugInternalOpcCycle(opc_cycles)

		// Just increment the Opcode cycle Counter
		if Opc_cycle_count < opc_cycles {
			Opc_cycle_count++

			// After spending the cycles needed, execute the opcode
		} else {
			// Print Opcode Debug Message
			opc_BNE_DebugMsg(bytes, value)

			// Increment PC
			PC += bytes

			// Reset Internal Opcode Cycle counters
			resetIntOpcCycleCounters()
		}

	} else { // If P[1] = 0 (Not Zero) Jump to address

		// Print internal opcode cycle
		debugInternalOpcCycleBranch(opc_cycles)

		// Just increment the Opcode cycle Counter
		if Opc_cycle_count < opc_cycles+1+Opc_cycle_extra {
			Opc_cycle_count++

			// After spending the cycles needed, execute the opcode
		} else {
			// Print Opcode Debug Message
			opc_BNE_DebugMsg(bytes, value)

			// PC + the number of bytes to jump on carry clear
			PC += uint16(value)

			// Increment PC
			PC += bytes

			// Reset Internal Opcode Cycle counters
			resetIntOpcCycleCounters()
		}
	}
}

func opc_BNE_DebugMsg(bytes uint16, value int8) {
	if Debug {
		opc_string := debug_decode_opc(bytes)
		if P[1] == 1 { // If P[1] = 1 (Zero Flag)
			dbg_show_message = fmt.Sprintf("\n\tOpcode %s [Mode: Relative]\tBNE  Branch on Result not Zero.\t| Zero Flag(P1) = %d | PC += 2\n", opc_string, P[1])
		} else { // If P[1] = 0 (Not Zero) Jump to address
			dbg_show_message = fmt.Sprintf("\n\tOpcode %s\tBNE  Branch on Result not Zero.\tZero Flag(P1) = %d, JUMP TO 0x%04X\n", opc_string, P[1], PC+2+uint16(value))
		}
		fmt.Println(dbg_show_message)
	}
}
