package VGS

import (
	"fmt"
	"os"
	"time"
)


// Initialization
func Initialize() {

	// Clean Memory Array
	Memory					= [65536]byte{}
	MemTIAWrite				= [14]byte{}
	// Clean CPU Variables
	PC				= 0
	opcode			= 0
	X				= 0
	Y				= 0
	A				= 0
	P				= [8]byte{}

	// Cycles
	counter_F_Cycle	= 0
	opc_cycle_count	= 1		// Opcode cycle counter
	opc_cycle_extra	= 0		// Opcode extra cycle

	// Counters
	counter_Frame	= 0
	counter_F_Cycle	= 0
	counter_IPS		= 0
	counter_FPS		= 0
	counter_DPS		= 0

	// Beamer
	beamIndex		= 0
	old_beamIndex	= 0

	// TIA
	line			=   1
	line_max		= 262
	TIA_Update		=  -1
	VSYNC_passed	= false	// Workaround for WSYNC before VSYNC

	// Debug Timing
	debugTiming_Limit = 0.00001

	// Player Vertical Positioning
	XPositionP0			= 0
	XFinePositionP0		= 0
	XPositionP1			= 0
	XFinePositionP1		= 0

	// Reset Controllers Buttons to 1 (not pressed)
	Memory[SWCHA] = 0xFF //1111 11111

	// Debug screen opcode message Slice
	dbg_opc_messages = dbg_opc_messages[:0]
	debug_opc_text = ""
	dbg_running_opc = true
}

func InitializeTimers() {
	// Start Timers
	clock_timer		= time.NewTicker(time.Nanosecond)	// CPU Clock
	screenRefresh_timer	= time.NewTicker(time.Second / 60)	// 60Hz Clock for screen refresh rate
	messagesClock_timer		= time.NewTicker(time.Second * 5)			// Clock used to display messages on screen
}


// Reset Vector // 0xFFFC | 0xFFFD (Little Endian)
func Reset() {
	// Read the Opcode from PC+1 and PC bytes (Little Endian)
	PC = uint16(Memory[0xFFFD])<<8 | uint16(Memory[0xFFFC])
}

func Show() {
	// fmt.Printf("\n\nCycle: %d\tOpcode: %02X\tPC: 0x%02X(%d)\tA: 0x%02X\tX: 0x%02X\tY: 0x%02X\tP: %d\tSP: %02X\tStack: [%d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d]\tRESPO0: %d\tGRP0: %08b\tCOLUP0: %02X\tCTRLPF: %08b", Cycle, Opcode, PC, PC, A, X, Y, P, SP, Memory[0xFF], Memory[0xFE], Memory[0xFD], Memory[0xFC], Memory[0xFB], Memory[0xFA], Memory[0xF9], Memory[0xF8], Memory[0xF7], Memory[0xF6], Memory[0xF5], Memory[0xF4], Memory[0xF3], Memory[0xF2], Memory[0xF1], Memory[0xF0], Memory[RESP0], Memory[GRP0], Memory[COLUP0], Memory[CTRLPF] )
	fmt.Printf("\nCycle: %d\tOpcode: %02X\tPC: 0x%02X(%d)\tA: 0x%02X\tX: 0x%02X\tY: 0x%02X\tP: %d\tSP: %02X\tGRP0: %08b\tHMP0: %02X\tBeam_index: %d\n", counter_F_Cycle, opcode, PC, PC, A, X, Y, P, SP, Memory[GRP0], Memory[HMP0], beamIndex )
}


// CPU Interpreter
func CPU_Interpreter() {

	// Read the Next Instruction to be executed
	opcode = Memory[PC]

	// Print Cycle and Debug Information
	if Debug {
		// Just show in the first opcode cycle
		if opc_cycle_count == 1 {
			Show()
		}

		// Clean opcode message
		debug_opc_text = ""

	}

	// Map Opcode
	switch opcode {

		//-------------------------------------------------- Implied --------------------------------------------------//

		case 0x78:	// Instruction SEI
			opc_SEI( 1, 2 )

		case 0x38:	// Instruction SEC
			opc_SEC( 1, 2 )

		case 0x18:	// Instruction CLC
			opc_CLC( 1, 2 )

		case 0xD8:	// Instruction CLD
			opc_CLD( 1, 2 )

		case 0x8A:	// Instruction TXA
			opc_TXA( 1, 2 )

		case 0xAA:	// Instruction TAX
			opc_TAX( 1, 2 )

		case 0xA8:	// Instruction TAY
			opc_TAY( 1, 2 )

		case 0xCA:	// Instruction DEX
			opc_DEX( 1, 2 )

		case 0x88:	// Instruction DEY
			opc_DEY( 1, 2 )

		case 0x9A:	// Instruction TXS
			opc_TXS( 1, 2 )

		case 0x48:	// Instruction PHA
			opc_PHA( 1, 3 )

		case 0x68:	// Instruction PLA
			opc_PLA( 1, 4 )

		case 0x00:	// Instruction BRK
			opc_BRK( 1, 7 )

		case 0xC8:	// Instruction INY
			opc_INY( 1, 2 )

		case 0xE8:	// Instruction INX
			opc_INX( 1, 2 )

		case 0x60:	// Instruction RTS
			opc_RTS( 1, 6 )

		case 0xEA:	// Instruction NOP
			opc_NOP( 1, 2 )

		//-------------------------------------------------- Just zeropage --------------------------------------------------//

		case 0xE6:	// Instruction INC (zeropage)
			if opc_cycle_count == 1 {
				memAddr, memMode = addr_mode_Zeropage(PC+1)
			}
			opc_INC( memAddr, memMode, 2, 5 )

		//-------------------------------------------- Branches - just relative ---------------------------------------------//

		case 0xD0:	// Instruction BNE (relative)
			if opc_cycle_count == 1 {
				memValue = addr_mode_Relative(PC+1)
			}
			opc_BNE( memValue, 2, 2 )

		case 0x90:	// Instruction BCC (relative)
			if opc_cycle_count == 1 {
				memValue = addr_mode_Relative(PC+1)
			}
			opc_BCC( memValue, 2, 2 )

		case 0xB0:	// Instruction BCS (relative)
			if opc_cycle_count == 1 {
				memValue = addr_mode_Relative(PC+1)
			}
			opc_BCS( memValue, 2, 2 )

		case 0x30:	// Instruction BMI (relative)
			if opc_cycle_count == 1 {
				memValue = addr_mode_Relative(PC+1)
			}
			opc_BMI( memValue, 2, 2 )

		case 0x10:	// Instruction BPL (relative)
			if opc_cycle_count == 1 {
				memValue = addr_mode_Relative(PC+1)
			}
			opc_BPL( memValue, 2, 2 )

		//-------------------------------------------------- LDX --------------------------------------------------//

		case 0xA2:	// Instruction LDX (immediate)
			if opc_cycle_count == 1 {
				memAddr, memMode = addr_mode_Immediate(PC+1)
			}
			opc_LDX( memAddr, memMode, 2, 2 )

		//-------------------------------------------------- STX --------------------------------------------------//

		case 0x86: // Instruction STX (zeropage)
			if opc_cycle_count == 1 {
				memAddr, memMode = addr_mode_Zeropage(PC+1)
			}
			opc_STX( memAddr, memMode, 2, 3 )

		//-------------------------------------------------- JMP --------------------------------------------------//

		case 0x4C:	// Instruction JMP (absolute)
			if opc_cycle_count == 1 {
				memAddr, memMode = addr_mode_Absolute(PC+1)
			}
			opc_JMP( memAddr, memMode, 3, 3 )

		case 0x20:	// Instruction JSR (absolute)
			if opc_cycle_count == 1 {
				memAddr, memMode = addr_mode_Absolute(PC+1)
			}
			opc_JSR( memAddr, memMode, 3, 6 )

		//-------------------------------------------------- BIT --------------------------------------------------//

		case 0x2C:	// Instruction BIT (absolute)
			if opc_cycle_count == 1 {
				memAddr, memMode = addr_mode_Absolute(PC+1)
			}
			opc_BIT( memAddr, memMode, 3, 4 )

		case 0x24:	// Instruction BIT (zeropage)
			if opc_cycle_count == 1 {
				memAddr, memMode = addr_mode_Zeropage(PC+1)
			}
			opc_BIT( memAddr, memMode, 2, 3 )

		//-------------------------------------------------- LDA --------------------------------------------------//

		case 0xA9:	// Instruction LDA (immediate)
			if opc_cycle_count == 1 {
				memAddr, memMode = addr_mode_Immediate(PC+1)
			}
			opc_LDA( memAddr, memMode, 2, 2 )

		case 0xA5:	// Instruction LDA (zeropage)
			if opc_cycle_count == 1 {
				memAddr, memMode = addr_mode_Zeropage(PC+1)
			}
			opc_LDA( memAddr, memMode, 2, 3 )

		case 0xB9:	// Instruction LDA (absolute,Y)
			if opc_cycle_count == 1 {
				memAddr, memMode = addr_mode_AbsoluteY(PC+1)
			}
			opc_LDA( memAddr, memMode, 3, 4 )

		case 0xBD:	// Instruction LDA (absolute,X)
			if opc_cycle_count == 1 {
				memAddr, memMode = addr_mode_AbsoluteX(PC+1)
			}
			opc_LDA( memAddr, memMode, 3, 4 )

		case 0xB1:	// Instruction LDA (indirect,Y)
			if opc_cycle_count == 1 {
				memAddr, memMode = addr_mode_IndirectY(PC+1)
			}
			opc_LDA( memAddr, memMode, 2, 5 )

		case 0xB5:	// Instruction LDA (zeropage,X)
			if opc_cycle_count == 1 {
				memAddr, memMode = addr_mode_ZeropageX(PC+1)
			}
			opc_LDA( memAddr, memMode, 2, 4 )

		case 0xAD:	// Instruction LDA (absolute)
			if opc_cycle_count == 1 {
				memAddr, memMode = addr_mode_Absolute(PC+1)
			}
			opc_LDA( memAddr, memMode, 3, 4 )

		//-------------------------------------------------- LDY --------------------------------------------------//

		case 0xA0:	// Instruction LDY (immediate)
			if opc_cycle_count == 1 {
				memAddr, memMode = addr_mode_Immediate(PC+1)
			}
			opc_LDY( memAddr, memMode, 2, 2 )

		case 0xA4:	// Instruction LDY (zeropage)
			if opc_cycle_count == 1 {
				memAddr, memMode = addr_mode_Zeropage(PC+1)
			}
			opc_LDY( memAddr, memMode, 2, 3 )

		//-------------------------------------------------- STY --------------------------------------------------//

		case 0x84:	// Instruction STY (zeropage)
			if opc_cycle_count == 1 {
				memAddr, memMode = addr_mode_Zeropage(PC+1)
			}
			opc_STY( memAddr, memMode,2 , 3 )

		//-------------------------------------------------- CPY --------------------------------------------------//

		case 0xC0:	// Instruction CPY (immediate)
			if opc_cycle_count == 1 {
				memAddr, memMode = addr_mode_Immediate(PC+1)
			}
			opc_CPY( memAddr, memMode, 2, 2 )

		case 0xC4:	// Instruction STY (zeropage)
			if opc_cycle_count == 1 {
				memAddr, memMode = addr_mode_Zeropage(PC+1)
			}
			opc_CPY( memAddr, memMode, 2, 3 )

		//-------------------------------------------------- CPX --------------------------------------------------//

		case 0xE0:	// Instruction CPX (immediate)
			if opc_cycle_count == 1 {
				memAddr, memMode = addr_mode_Immediate(PC+1)
			}
			opc_CPX( memAddr, memMode, 2, 2 )

		//-------------------------------------------------- SBC --------------------------------------------------//

		case 0xE5:	// Instruction STY (zeropage)
			if opc_cycle_count == 1 {
				memAddr, memMode = addr_mode_Zeropage(PC+1)
			}
			opc_SBC( memAddr, memMode, 2, 3 )

		case 0xE9:	// Instruction STY (immediate)
			if opc_cycle_count == 1 {
				memAddr, memMode = addr_mode_Immediate(PC+1)
			}
			opc_SBC( memAddr, memMode, 2, 2 )

		//-------------------------------------------------- DEC --------------------------------------------------//

		case 0xC6:	// Instruction DEC (zeropage)
			if opc_cycle_count == 1 {
				memAddr, memMode = addr_mode_Zeropage(PC+1)
			}
			opc_DEC( memAddr, memMode, 2, 5 )

		//-------------------------------------------------- AND --------------------------------------------------//

		case 0x29:	// Instruction AND (immediate)
			if opc_cycle_count == 1 {
				memAddr, memMode = addr_mode_Immediate(PC+1)
			}
			opc_AND( memAddr, memMode, 2 , 2 )

		//-------------------------------------------------- ORA --------------------------------------------------//

		case 0x05:	// Instruction ORA (zeropage)
			if opc_cycle_count == 1 {
				memAddr, memMode = addr_mode_Zeropage(PC+1)
			}
			opc_ORA( memAddr, memMode, 2, 3 )

		//-------------------------------------------------- EOR --------------------------------------------------//

		case 0x49:	// Instruction EOR (immediate)
			if opc_cycle_count == 1 {
				memAddr, memMode = addr_mode_Immediate(PC+1)
			}
			opc_EOR( memAddr, memMode, 2, 2 )

		case 0x45:	// Instruction EOR (zeropage)
			if opc_cycle_count == 1 {
				memAddr, memMode = addr_mode_Zeropage(PC+1)
			}
			opc_EOR( memAddr, memMode, 2, 3 )

		//-------------------------------------------------- SHIFT --------------------------------------------------//

		case 0x0A:	// Instruction ASL (accumulator)
			opc_ASL( 1, 2 )

		case 0x4A:	// Instruction LSR (accumulator)
			opc_LSR( 1, 2 )

		//-------------------------------------------------- CMP --------------------------------------------------//

		case 0xC5:	// Instruction CMP (zeropage)
			if opc_cycle_count == 1 {
				memAddr, memMode = addr_mode_Zeropage(PC+1)
			}
			opc_CMP( memAddr, memMode, 2, 3 )

		case 0xC9:	// Instruction CMP (immediate)
			if opc_cycle_count == 1 {
				memAddr, memMode = addr_mode_Immediate(PC+1)
			}
			opc_CMP( memAddr, memMode, 2, 2 )

		//-------------------------------------------------- STA --------------------------------------------------//

		case 0x95:	// Instruction STA (zeropage,X)
			if opc_cycle_count == 1 {
				memAddr, memMode = addr_mode_ZeropageX(PC+1)
			}
			opc_STA( memAddr, memMode, 2, 4 )

		case 0x85:	// Instruction STA (zeropage)
			if opc_cycle_count == 1 {
				memAddr, memMode = addr_mode_Zeropage(PC+1)
			}
			opc_STA( memAddr, memMode, 2, 3 )

		case 0x99:	// Instruction STA (absolute,Y)
			if opc_cycle_count == 1 {
				memAddr, memMode = addr_mode_AbsoluteY(PC+1)
			}
			opc_STA( memAddr, memMode, 3, 5 )

		case 0x8D:	// Instruction STA (absolute)
			if opc_cycle_count == 1 {
				memAddr, memMode = addr_mode_Absolute(PC+1)
			}
			opc_STA( memAddr, memMode, 3, 4 )

		//-------------------------------------------------- ADC --------------------------------------------------//

		case 0x65:	// Instruction ADC (zeropage)
			if opc_cycle_count == 1 {
				memAddr, memMode = addr_mode_Zeropage(PC+1)
			}
			opc_ADC( memAddr, memMode, 2, 3 )

		//-------------------------------------------------- ROL --------------------------------------------------//

		case 0x26:	// Instruction ROL (zeropage)
			if opc_cycle_count == 1 {
				memAddr, memMode = addr_mode_Zeropage(PC+1)
			}
			opc_ROL( memAddr, memMode, 2, 5 )

		//-------------------------------------------------- ISB? FF --------------------------------------------------//

		// ISB (INC FOLLOWED BY SBC - IMPLEMENT IT!!!!!!)
		// FF (Filled ROM)
		// case 0xFF:
		// 	if Debug {
		// 		fmt.Printf("\tOpcode %02X [1 byte]\tFilled ROM.\tPC incremented.\n", opcode)
		//
		// 		// Collect data for debug interface just on first cycle
		// 		if opc_cycle_count == 1 {
		// 			debug_opc_text		= fmt.Sprintf("%04x     ISB*     ;%d", PC, opc_cycles)
		// 			dbg_opc_bytes		= bytes
		// 			dbg_opc_opcode		= opcode
		// 		}
		// 	}
		// 	PC +=1

		//-------------------------------------------- No Opcode Found --------------------------------------------//

		default:
			fmt.Printf("\tOPCODE %02X NOT IMPLEMENTED!\n\n", opcode)
			os.Exit(0)

	}

	// Increment Cycle
	counter_F_Cycle ++

	// Pause = true
	// Increment Instructions per second counter
	counter_IPS ++

}

// Collect data for debug interface after finished running the opcode
func dbg_opcode_message(mnm string, bytes uint16, opc_cycles_sum byte) {

	if bytes == 1 {
		debug_opc_text = fmt.Sprintf("%04x\t\t%s\t\t;%d\t\t\t%02x", PC, mnm, opc_cycles_sum, opcode)
		dbg_opc_bytes = bytes
	} else if bytes == 2 {
		debug_opc_text = fmt.Sprintf("%04x\t\t%s\t\t;%d\t\t\t%02x %02x", PC, mnm, opc_cycles_sum, opcode, Memory[PC+1])
		dbg_opc_bytes = bytes
	} else if bytes == 3 {
		debug_opc_text = fmt.Sprintf("%04x\t\t%s\t\t;%d\t\t\t%02x %02x %02x", PC, mnm, opc_cycles_sum, opcode, Memory[PC+2], Memory[PC+1])
		dbg_opc_bytes = bytes
	}

	// Reset running opcode flag
	dbg_running_opc = false

}
