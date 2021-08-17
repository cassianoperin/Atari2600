package VGS

import "fmt"

// RTI  Return from Interrupt
//
//      pull SR, pull PC
//
//                                      N Z C I D V
//                                      from stack
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      implied       RTI           40    1    6

// Order:
// Restore P
// Restore PC(lo)
// Restore PC(hi)

func opc_RTI(bytes uint16, opc_cycles byte) {

	// Update Global Opc_cycles value
	Opc_cycles = opc_cycles

	// Print internal opcode cycle
	debugInternalOpcCycle(opc_cycles)

	// Just increment the Opcode cycle Counter
	if Opc_cycle_count < opc_cycles {
		Opc_cycle_count++

		// After spending the cycles needed, execute the opcode
	} else {

		// ---------- Restore P ---------- //

		// 6502 handle Stack at the end of first memory page
		SP_Address := uint16(SP+1) + 256

		// Read data from Memory (adress in Memory Bus) into Data Bus
		memData := dataBUS_Read(SP_Address)

		// Turn the stack value into the processor status
		for i := 0; i < len(P); i++ {

			// The B Flag, PLP and RTI pull a byte from the stack and set all the flags. They ignore bits 5 and 4.
			if i == 4 || i == 5 {
				// Just ignore both
			} else {
				P[i] = (memData >> i) & 0x01
			}
		}

		SP++

		// ---------- Restore PC ---------- //

		// Read the Opcode from PC+1 and PC bytes (Little Endian)
		memData_LSB := dataBUS_Read(SP_Address + 2) // Read data from Memory (adress in Memory Bus) into Data Bus
		memData_MSB := dataBUS_Read(SP_Address + 1)

		PC = uint16(memData_LSB)<<8 | uint16(memData_MSB)
		SP += 2

		// Print Opcode Debug Message
		opc_RTI_DebugMsg(bytes, SP_Address)

		// Reset Internal Opcode Cycle counters
		resetIntOpcCycleCounters()
	}
}

func opc_RTI_DebugMsg(bytes uint16, SP_Address uint16) {
	if Debug {
		opc_string := debug_decode_opc(bytes)
		dbg_show_message = fmt.Sprintf("\n\tOpcode %s [Mode: Implied]\tRTI  Return from Interrupt (P and PC from Stack).\tP = Memory[0x%02X] %d | PC = 0x%04X | SP: 0x%02X\n", opc_string, SP_Address, P, PC, SP)
		fmt.Println(dbg_show_message)
	}
}
