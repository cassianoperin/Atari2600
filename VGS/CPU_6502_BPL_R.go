package VGS

import "fmt"

// BPL  Branch on Result Plus
//
//      branch on N = 0                  N Z C I D V
//                                       - - - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      relative      BPL oper      10    2     2**

func opc_BPL(memAddr uint16, bytes uint16, opc_cycles byte) {

	// Update Global Opc_cycles value
	Opc_cycles = opc_cycles

	// Read data from Memory (adress in Memory Bus) into Data Bus
	memData := dataBUS_Read(memAddr)

	// Get the Two's complement value of value in Memory
	value := DecodeTwoComplement(memData) // value is SIGNED

	if P[7] == 0 { // If Positive

		// Print internal opcode cycle
		debugInternalOpcCycleBranch(opc_cycles)

		// Just increment the Opcode cycle Counter
		if Opc_cycle_count < opc_cycles+1+Opc_cycle_extra {
			Opc_cycle_count++

			// After spending the cycles needed, execute the opcode
		} else {
			// Print Opcode Debug Message
			opc_BPL_DebugMsg(bytes, value)

			// PC + the number of bytes to jump on carry clear
			PC += uint16(value)

			// Increment PC
			PC += bytes

			// Reset Internal Opcode Cycle counters
			resetIntOpcCycleCounters()
		}

	} else { // If not positive

		// Print internal opcode cycle
		debugInternalOpcCycle(opc_cycles)

		// Just increment the Opcode cycle Counter
		if Opc_cycle_count < opc_cycles {
			Opc_cycle_count++

			// After spending the cycles needed, execute the opcode
		} else {
			// Print Opcode Debug Message
			opc_BPL_DebugMsg(bytes, value)

			// Increment PC
			PC += bytes

			// Reset Internal Opcode Cycle counters
			resetIntOpcCycleCounters()
		}
	}
}

func opc_BPL_DebugMsg(bytes uint16, value int8) {
	if Debug {
		opc_string := debug_decode_opc(bytes)
		if P[7] == 0 { // If Positive
			dbg_show_message = fmt.Sprintf("\n\tOpcode %0s [Mode: Relative]\tBPL  Branch on Result POSITIVE.\tNEGATIVE flag DISABLED, JUMP TO 0x%04X\n", opc_string, PC+2+uint16(value))
		} else { // If not positive
			dbg_show_message = fmt.Sprintf("\n\tOpcode %s\tBPL  Branch on Result POSITIVE.\t\tNEGATIVE flag enabled, PC+=2\n", opc_string)
		}
		fmt.Println(dbg_show_message)
	}
}
