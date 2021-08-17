package VGS

import "fmt"

// BCC  Branch on Carry Clear
//
//      branch on C = 0                  N Z C I D V
//                                       - - - - - -
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      relative      BCC oper      90    2     2**

func opc_BCC(memAddr uint16, bytes uint16, opc_cycles byte) {

	// Update Global Opc_cycles value
	Opc_cycles = opc_cycles

	// Read data from Memory (adress in Memory Bus) into Data Bus
	memData := dataBUS_Read(memAddr)

	// Get the Two's complement value of value in Memory
	value := DecodeTwoComplement(memData) // value is SIGNED

	if P[0] == 0 { // If carry is clear

		// Print internal opcode cycle
		debugInternalOpcCycleBranch(opc_cycles)

		// Just increment the Opcode cycle Counter
		if Opc_cycle_count < opc_cycles+1+Opc_cycle_extra {
			Opc_cycle_count++

			// After spending the cycles needed, execute the opcode
		} else {

			// Print Opcode Debug Message
			opc_BCC_DebugMsg(bytes, value)

			// PC + the number of bytes to jump on carry clear
			PC += uint16(value)

			// Increment PC
			PC += bytes

			// Reset Internal Opcode Cycle counters
			resetIntOpcCycleCounters()
		}
	} else { // If carry is set

		// Print internal opcode cycle
		debugInternalOpcCycle(opc_cycles)

		// Just increment the Opcode cycle Counter
		if Opc_cycle_count < opc_cycles {
			Opc_cycle_count++

			// After spending the cycles needed, execute the opcode
		} else {

			// Print Opcode Debug Message
			opc_BCC_DebugMsg(bytes, value)

			// Increment PC
			PC += bytes

			// Reset Internal Opcode Cycle counters
			resetIntOpcCycleCounters()
		}
	}
}

func opc_BCC_DebugMsg(bytes uint16, value int8) {
	if Debug {
		opc_string := debug_decode_opc(bytes)
		if P[0] == 0 { // If carry is clear
			dbg_show_message = fmt.Sprintf("\n\tOpcode %s [Mode: Relative]\tBCC  Branch on Carry Clear.\tCarry EQUAL 0, JUMP TO 0x%04X\n", opc_string, PC+2+uint16(value))
		} else { // If carry is set
			dbg_show_message = fmt.Sprintf("\n\tOpcode %s\tBCC  Branch on Carry Clear.\tCarry NOT EQUAL 0, PC+2\n", opc_string)
		}
		fmt.Println(dbg_show_message)
	}
}
