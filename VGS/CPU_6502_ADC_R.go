package VGS

import (
	"fmt"
	"strconv"
)

// ADC  Add Memory to Accumulator with Carry (zeropage)
//
//      A + M + C -> A, C                N Z C I D V
//     	                                 + + + - - +
//
//      addressing    assembler    opc  bytes  cyles
//      --------------------------------------------
//      immediate	  ADC #oper	    69    2     2
//      zeropage      ADC oper      65    2     3
//      zeropage,X    ADC oper,X    75    2     4
//      absolute      ADC oper      6D    3     4
//      absolute,X    ADC oper,X    7D    3     4*
//      absolute,Y    ADC oper,Y    79    3     4*
//      (indirect,X)  ADC (oper,X)  61    2     6
//      (indirect),Y  ADC (oper),Y  71    2     5*

func opc_ADC(memAddr uint16, mode string, bytes uint16, opc_cycles byte) {

	// Update Global Opc_cycles value
	Opc_cycles = opc_cycles

	// Print internal opcode cycle
	debugInternalOpcCycleExtras(opc_cycles)

	if Opc_cycle_count < opc_cycles+Opc_cycle_extra { // Just increment the Opcode cycle Counter
		Opc_cycle_count++

	} else { // After spending the cycles needed, execute the opcode

		// Original value of A and P0
		var (
			original_A  byte = A
			original_P0 byte = P[0]
			memData     byte = dataBUS_Read(memAddr) // Read data from Memory (adress in Memory Bus) into Data Bus
		)

		// --------------------------------- Binary / Hex Mode -------------------------------- //

		if P[3] == 0 {

			A = A + memData + P[0]

			flags_V(original_A, memData, original_P0)
			flags_C_ADC_SBC(original_A, memData, original_P0)
			flags_Z(A)
			flags_N(A)

			// ----------------------------------- Decimal Mode ----------------------------------- //

		} else {

			var bcd_Mem int64

			// Store the decimal value of the original A (hex)
			bcd_A, _ := strconv.ParseInt(fmt.Sprintf("%X", A), 0, 32)

			// Store the decimal value of the original Memory Address (hex)
			bcd_Mem, _ = strconv.ParseInt(fmt.Sprintf("%X", memData), 0, 32)

			// Store the decimal result of A (must be trasformed in hex to be stored)
			tmp_A := byte(bcd_A) + byte(bcd_Mem) + P[0]

			// Convert the Decimal Result in to Hex to be returned to Accumulator
			bcd_Result, _ := strconv.ParseInt(fmt.Sprintf("%d", tmp_A), 16, 32)

			// Tranform the uint64 into a byte (if > 255 will be rotated)
			A = byte(bcd_Result)

			flags_V(original_A, memData, original_P0)
			flags_C_ADC_DECIMAL(bcd_Result)
			flags_Z(A)
			flags_N(A)

		}

		// Print Opcode Debug Message
		opc_ADC_DebugMsg(bytes, mode, original_A, memAddr, original_P0, memData)

		// Increment PC
		PC += bytes

		// Reset Internal Opcode Cycle counters
		resetIntOpcCycleCounters()
	}
}

func opc_ADC_DebugMsg(bytes uint16, mode string, original_A byte, memAddr uint16, original_P0 byte, memData byte) {
	if Debug {
		opc_string := debug_decode_opc(bytes)
		if P[3] == 0 { // Decimal flag OFF (Binary or Hex Mode)
			dbg_show_message = fmt.Sprintf("\n\tOpcode %s [Mode: %s]\tADC  Add Memory to Accumulator with Carry [Binary/Hex Mode]\tA = A(%d) + Memory[0x%02X](%d) + Carry (%d)) = %d\n", opc_string, mode, original_A, memAddr, memData, original_P0, A)

		} else { // Decimal flag ON (Decimal Mode)
			dbg_show_message = fmt.Sprintf("\n\tOpcode %s [Mode: %s]\tADC  Add Memory to Accumulator with Carry [Decimal Mode]\tA = A(0x%02x) + Memory[0x%02X](0x%02x) + Carry (0x%02x)) = 0x%02X\n", opc_string, mode, original_A, memAddr, memData, original_P0, A)
		}
		fmt.Println(dbg_show_message)
	}
}
