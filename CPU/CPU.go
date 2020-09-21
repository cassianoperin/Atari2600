package CPU

import (
	"fmt"
	"os"
	"time"
)

// Components
var (

	Memory		[65536]byte	// Memory
	MemTIAWrite	[14]byte	// TIA Read-Only additional Registers
	PC			uint16		// Program Counter
	A			byte			// Accumulator
	X			byte			// Index Register X
	Y			byte			// Index Register Y
	SP			byte			// Stack Pointer
	// The stack pointer is addressing 256 bytes in page 1 of memory, ie. values 00h-FFh will address memory at 0100h-01FFh.
	// As for most other CPUs, the stack pointer is decrementing when storing data.
	// However, in the 65XX world, it points to the first FREE byte on stack, so, when initializing stack to top set S=(1)FFh (rather than S=(2)00h).
	P			[8]byte
	//P			byte			// Processor Status Register
	// Processor Status Register (Flags)
	// Bit  Name  Expl.
	// 7    N     Negative/Sign (0=Positive, 1=Negative)
	// 6    V     Overflow      (0=No Overflow, 1=Overflow)
	// 5    -     Not used      (Always 1)
	// 4    B     Break Flag    (0=IRQ/NMI, 1=RESET or BRK/PHP opcode)
	// 3    D     Decimal Mode  (0=Normal, 1=BCD Mode for ADC/SBC opcodes)
	// 2    I     IRQ Disable   (0=IRQ Enable, 1=IRQ Disable)
	// 1    Z     Zero          (0=Nonzero, 1=Zero)
	// 0    C     Carry         (0=No Carry, 1=Carry)

	// CPU Variables
	Opcode			byte		// CPU Operation Code
	Cycle			uint16		// CPU Cycle Counter
	opc_cycle_count	byte		// Opcode cycle counter
	opc_cycle_extra	byte		// Opcode extra cycle
	memAddr			uint16		// Receive the memory address needed by the opcode
	mode			string		// Receive the addressing mode used in the debug
	value			int8		// Receive the value needed by the opcode (branches)
	IPS				uint16		// Instructions per second Counter

	// Timers
	Clock			*time.Ticker	// CPU Clock // CPU: MOS Technology 6507 @ 1.19 MHz;
	ScreenRefresh		*time.Ticker	// Screen Refresh
	Second			= time.Tick(time.Second)			// 1 second to track FPS and draws

	// Players Vertical Positioning
	XPositionP0		byte
	XFinePositionP0	int8
	XPositionP1		byte
	XFinePositionP1	int8

	// ------------------ Personal Control Flags ------------------ //
	Beam_index	byte = 0		// Beam index to control where to draw objects using cpu cycles

	TIA_Update	int8 = -1		// Tells Graphics that a TIA register was changed (values >= 0 (addresses) will be detected)

	// Debug Timing Measurement
	DebugTiming 		bool	= false
	DebugTimingLimit	float64 = 0.00001
	StartCycle			time.Time
	StartTIA			time.Time
	StartTIA_BG			time.Time

	// Pause
	Pause		bool = false

	// Debug
	Debug 		bool = false
)


const (
	//-------------------------------------------------- Memory locations -------------------------------------------------//

	//0000-002C - TIA (write)
	//0030-003D - TIA (read) - (sometimes mirrored at 0030-003D)
	//0080-00FF - RIOT (RAM) (128 bytes) -- Stack uses the last addresses
	//0280-0297 - RIOT (I/O, Timer)
	//F000-FFFF - Cartridge (ROM)

	//------------------- 0000-002C - TIA (write)
	VSYNC 			byte = 0x00		//0000 00x0   Vertical Sync Set-Clear
	VBLANK			byte = 0x01		//xx00 00x0   Vertical Blank Set-Clear
	WSYNC			byte = 0x02		//---- ----   Wait for Horizontal Blank
	RSYNC			byte = 0x03		//---- ----   Reset Horizontal Sync Counter
	NUSIZ0			byte = 0x04		//00xx 0xxx   Number-Size player/missle 0
	NUSIZ1			byte = 0x05		//00xx 0xxx   Number-Size player/missle 1
	COLUP0			byte = 0x06		//xxxx xxx0   Color-Luminance Player 0
	COLUP1			byte = 0x07		//xxxx xxx0   Color-Luminance Player 1
	COLUPF			byte	= 0x08		//xxxx xxx0   Color-Luminance Playfield
	COLUBK			byte	= 0x09		//xxxx xxx0   Color-Luminance Background
	CXCLR			byte =	0x2C	//---- ----   Clear Collision Latches


	// CTRLPLF (8 bits register)
	// D0 = 0 Repeat the PF, D0 = 1 = Reflect the PF
	// D1 = Score == Color of the score will be the same as player
	// D2 = Priority == Player behind the playfield
	// D4-5 = Ball Size (1, 2, 4, 8)
	CTRLPF			byte = 0x0A		//00xx 0xxx   Control Playfield, Ball, Collisions
	REFP0			byte = 0x0B		//0000 x000   Reflection Player 0
	REFP1			byte = 0x0C		//0000 x000   Reflection Player 1
	PF0 				byte	= 0x0D		//xxxx 0000   Playfield Register Byte 0
	PF1 				byte	= 0x0E		//xxxx 0000   Playfield Register Byte 1
	PF2 				byte	= 0x0F		//xxxx 0000   Playfield Register Byte 2
	GRP0				byte = 0x1B		//xxxx xxxx   Graphics Register Player 0
	GRP1				byte = 0x1C		//xxxx xxxx   Graphics Register Player 1
	RESP0 			byte	= 0x10		//---- ----   Reset Player 0
	RESP1 			byte	= 0x11		//---- ----   Reset Player 1
	HMP0				byte = 0x20		// xxxx 0000   Horizontal Motion Player 0
	HMP1				byte = 0x21		// xxxx 0000   Horizontal Motion Player 1


	// ;-------------------------------------------------------------------------------
	//
	// 			SEG.U TIA_REGISTERS_READ
	// 			ORG TIA_BASE_READ_ADDRESS
	//
	// ;															bit 7   bit 6
	CXM0P			byte = 0x00		//xx00 0000     Read Collision  M0-P1   M0-P0
	CXM1P			byte = 0x01		//xx00 0000                     M1-P0   M1-P1
	CXP0FB			byte = 0x02		//xx00 0000                     P0-PF   P0-BL
	CXP1FB			byte = 0x03		//xx00 0000                     P1-PF   P1-BL
	CXM0FB			byte = 0x04		//xx00 0000                     M0-PF   M0-BL
	CXM1FB			byte = 0x05		//xx00 0000                     M1-PF   M1-BL
	CXBLPF			byte = 0x06		//x000 0000                     BL-PF   -----
	CXPPMM			byte = 0x07		//xx00 0000                     P0-P1   M0-M1



	//------------------- 0280-0297 - RIOT (I/O, Timer)
	SWCHA			uint16 = 0x280		// Port A data register for joysticks: Bits 4-7 for player 1.  Bits 0-3 for player 2.



)



func Fine(HMPX byte) int8 {

	var value int8

	switch HMPX {
		case 0x70:
			value = -7
		case 0x60:
			value = -6
		case 0x50:
			value = -5
		case 0x40:
			value = -4
		case 0x30:
			value = -3
		case 0x20:
			value = -2
		case 0x10:
			value = -1
		case 0x00:
			value =  0
		case 0xF0:
			value =  1
		case 0xE0:
			value =  2
		case 0xD0:
			value =  3
		case 0xC0:
			value =  4
		case 0xB0:
			value =  5
		case 0xA0:
			value =  6
		case 0x90:
			value =  7
		case 0x80:
			value =  8
		default:
			fmt.Printf("\n\tInvalid HMP0 %02X!\n\n", HMP0)
			os.Exit(0)
		}

	return value

}


func MemPageBoundary(Address1, Address2 uint16) bool {

	var cross bool = false

	// Get the High byte only to compare
	// Page Boundary Cross detected
	if Address1 >> 8 != Address2 >> 8 {
		cross = true
		if Debug {
			fmt.Printf("\tMemory Page Boundary Cross detected! Add 1 cycle.\tPC High byte: %02X\tBranch High byte: %02X\n",Address1 >>8, Address2 >>8)
		}
	// NO Page Boundary Cross detected
	} else {
		if Debug {
			fmt.Printf("\tNo Memory Page Boundary Cross detected.\tPC High byte: %02X\tBranch High byte: %02X\n",Address1 >>8, Address2 >>8)
		}
	}

	return cross

}


// Initialization
func Initialize() {

	// Clean Memory Array
	Memory					= [65536]byte{}
	MemTIAWrite				= [14]byte{}
	// Clean CPU Variables
	PC				= 0
	Opcode			= 0
	Cycle			= 0
	opc_cycle_count	= 1		// Opcode cycle counter
	opc_cycle_extra	= 0		// Opcode extra cycle
	IPS				= 0

	// Start Timers
	Clock		= time.NewTicker(time.Nanosecond)	// CPU Clock
	ScreenRefresh	= time.NewTicker(time.Second / 60)	// 60Hz Clock for screen refresh rate

	// Reset Controllers Buttons to 1 (not pressed)
	Memory[SWCHA] = 0xFF //1111 11111

}


// Reset Vector // 0xFFFC | 0xFFFD (Little Endian)
func Reset() {
	// Read the Opcode from PC+1 and PC bytes (Little Endian)
	PC = uint16(Memory[0xFFFD])<<8 | uint16(Memory[0xFFFC])
	//fmt.Printf("\nRESET: %04X\n",PC)
}


// Decode Two's Complement
func DecodeTwoComplement(num byte) int8 {

	var sum int8 = 0

	for i := 0 ; i < 8 ; i++ {
		// Sum each bit and sum the value of the bit power of i (<<i)
		sum += (int8(num) >> i & 0x01) << i
	}

	return sum
}


func Show() {
	// fmt.Printf("\n\nCycle: %d\tOpcode: %02X\tPC: 0x%02X(%d)\tA: 0x%02X\tX: 0x%02X\tY: 0x%02X\tP: %d\tSP: %02X\tStack: [%d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d]\tRESPO0: %d\tGRP0: %08b\tCOLUP0: %02X\tCTRLPF: %08b", Cycle, Opcode, PC, PC, A, X, Y, P, SP, Memory[0xFF], Memory[0xFE], Memory[0xFD], Memory[0xFC], Memory[0xFB], Memory[0xFA], Memory[0xF9], Memory[0xF8], Memory[0xF7], Memory[0xF6], Memory[0xF5], Memory[0xF4], Memory[0xF3], Memory[0xF2], Memory[0xF1], Memory[0xF0], Memory[RESP0], Memory[GRP0], Memory[COLUP0], Memory[CTRLPF] )
	fmt.Printf("\nCycle: %d\tOpcode: %02X\tPC: 0x%02X(%d)\tA: 0x%02X\tX: 0x%02X\tY: 0x%02X\tP: %d\tSP: %02X\tGRP0: %08b\tHMP0: %02X\tBeam_index: %d\n", Cycle, Opcode, PC, PC, A, X, Y, P, SP, Memory[GRP0], Memory[HMP0], Beam_index )
}


// CPU Interpreter
func Interpreter() {

	// Read the Next Instruction to be executed
	Opcode = Memory[PC]

	// Print Cycle and Debug Information
	if Debug {
		// Just show in the first opcode cycle
		if opc_cycle_count == 1 {
			Show()
		}
	}

	// Map Opcode
	switch Opcode {

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
				memAddr, mode = addr_mode_Zeropage(PC+1)
			}
			opc_INC( memAddr, mode, 2, 5 )

		//-------------------------------------------- Branches - just relative ---------------------------------------------//

		case 0xD0:	// Instruction BNE (relative)
			if opc_cycle_count == 1 {
				value = addr_mode_Relative(PC+1)
			}
			opc_BNE( value, 2, 2 )

		case 0x90:	// Instruction BCC (relative)
			if opc_cycle_count == 1 {
				value = addr_mode_Relative(PC+1)
			}
			opc_BCC( value, 2, 2 )

		case 0xB0:	// Instruction BCS (relative)
			if opc_cycle_count == 1 {
				value = addr_mode_Relative(PC+1)
			}
			opc_BCS( value, 2, 2 )

		case 0x30:	// Instruction BMI (relative)
			if opc_cycle_count == 1 {
				value = addr_mode_Relative(PC+1)
			}
			opc_BMI( value, 2, 2 )

		case 0x10:	// Instruction BPL (relative)
			if opc_cycle_count == 1 {
				value = addr_mode_Relative(PC+1)
			}
			opc_BPL( value, 2, 2 )

		//-------------------------------------------------- LDX --------------------------------------------------//

		case 0xA2:	// Instruction LDX (immediate)
			if opc_cycle_count == 1 {
				memAddr, mode = addr_mode_Immediate(PC+1)
			}
			opc_LDX( memAddr, mode, 2, 2 )

		//-------------------------------------------------- STX --------------------------------------------------//

		case 0x86: // Instruction STX (zeropage)
			if opc_cycle_count == 1 {
				memAddr, mode = addr_mode_Zeropage(PC+1)
			}
			opc_STX( memAddr, mode, 2, 3 )

		//-------------------------------------------------- JMP --------------------------------------------------//

		case 0x4C:	// Instruction JMP (absolute)
			if opc_cycle_count == 1 {
				memAddr, mode = addr_mode_Absolute(PC+1)
			}
			opc_JMP( memAddr, mode, 3, 3 )

		case 0x20:	// Instruction JSR (absolute)
			if opc_cycle_count == 1 {
				memAddr, mode = addr_mode_Absolute(PC+1)
			}
			opc_JSR( memAddr, mode, 3, 6 )

		//-------------------------------------------------- BIT --------------------------------------------------//

		case 0x2C:	// Instruction BIT (absolute)
			if opc_cycle_count == 1 {
				memAddr, mode = addr_mode_Absolute(PC+1)
			}
			opc_BIT( memAddr, mode, 3, 4 )

		case 0x24:	// Instruction BIT (zeropage)
			if opc_cycle_count == 1 {
				memAddr, mode = addr_mode_Zeropage(PC+1)
			}
			opc_BIT( memAddr, mode, 2, 3 )

		//-------------------------------------------------- LDA --------------------------------------------------//

		case 0xA9:	// Instruction LDA (immediate)
			if opc_cycle_count == 1 {
				memAddr, mode = addr_mode_Immediate(PC+1)
			}
			opc_LDA( memAddr, mode, 2, 2 )

		case 0xA5:	// Instruction LDA (zeropage)
			if opc_cycle_count == 1 {
				memAddr, mode = addr_mode_Zeropage(PC+1)
			}
			opc_LDA( memAddr, mode, 2, 3 )

		case 0xB9:	// Instruction LDA (absolute,Y)
			if opc_cycle_count == 1 {
				memAddr, mode = addr_mode_AbsoluteY(PC+1)
			}
			opc_LDA( memAddr, mode, 3, 4 )

		case 0xBD:	// Instruction LDA (absolute,X)
			if opc_cycle_count == 1 {
				memAddr, mode = addr_mode_AbsoluteX(PC+1)
			}
			opc_LDA( memAddr, mode, 3, 4 )

		case 0xB1:	// Instruction LDA (indirect,Y)
			if opc_cycle_count == 1 {
				memAddr, mode = addr_mode_IndirectY(PC+1)
			}
			opc_LDA( memAddr, mode, 2, 5 )

		case 0xB5:	// Instruction LDA (zeropage,X)
			if opc_cycle_count == 1 {
				memAddr, mode = addr_mode_ZeropageX(PC+1)
			}
			opc_LDA( memAddr, mode, 2, 4 )

		case 0xAD:	// Instruction LDA (absolute)
			if opc_cycle_count == 1 {
				memAddr, mode = addr_mode_Absolute(PC+1)
			}
			opc_LDA( memAddr, mode, 3, 4 )

		//-------------------------------------------------- LDY --------------------------------------------------//

		case 0xA0:	// Instruction LDY (immediate)
			if opc_cycle_count == 1 {
				memAddr, mode = addr_mode_Immediate(PC+1)
			}
			opc_LDY( memAddr, mode, 2, 2 )

		case 0xA4:	// Instruction LDY (zeropage)
			if opc_cycle_count == 1 {
				memAddr, mode = addr_mode_Zeropage(PC+1)
			}
			opc_LDY( memAddr, mode, 2, 3 )

		//-------------------------------------------------- STY --------------------------------------------------//

		case 0x84:	// Instruction STY (zeropage)
			if opc_cycle_count == 1 {
				memAddr, mode = addr_mode_Zeropage(PC+1)
			}
			opc_STY( memAddr, mode,2 , 3 )

		//-------------------------------------------------- CPY --------------------------------------------------//

		case 0xC0:	// Instruction CPY (immediate)
			if opc_cycle_count == 1 {
				memAddr, mode = addr_mode_Immediate(PC+1)
			}
			opc_CPY( memAddr, mode, 2, 2 )

		case 0xC4:	// Instruction STY (zeropage)
			if opc_cycle_count == 1 {
				memAddr, mode = addr_mode_Zeropage(PC+1)
			}
			opc_CPY( memAddr, mode, 2, 3 )

		//-------------------------------------------------- CPX --------------------------------------------------//

		case 0xE0:	// Instruction CPX (immediate)
			if opc_cycle_count == 1 {
				memAddr, mode = addr_mode_Immediate(PC+1)
			}
			opc_CPX( memAddr, mode, 2, 2 )

		//-------------------------------------------------- SBC --------------------------------------------------//

		case 0xE5:	// Instruction STY (zeropage)
			if opc_cycle_count == 1 {
				memAddr, mode = addr_mode_Zeropage(PC+1)
			}
			opc_SBC( memAddr, mode, 2, 3 )

		case 0xE9:	// Instruction STY (immediate)
			if opc_cycle_count == 1 {
				memAddr, mode = addr_mode_Immediate(PC+1)
			}
			opc_SBC( memAddr, mode, 2, 2 )

		//-------------------------------------------------- DEC --------------------------------------------------//

		case 0xC6:	// Instruction DEC (zeropage)
			if opc_cycle_count == 1 {
				memAddr, mode = addr_mode_Zeropage(PC+1)
			}
			opc_DEC( memAddr, mode, 2, 5 )

		//-------------------------------------------------- AND --------------------------------------------------//

		case 0x29:	// Instruction AND (immediate)
			if opc_cycle_count == 1 {
				memAddr, mode = addr_mode_Immediate(PC+1)
			}
			opc_AND( memAddr, mode, 2 , 2 )

		//-------------------------------------------------- AND --------------------------------------------------//

		case 0x05:	// Instruction ORA (zeropage)
			if opc_cycle_count == 1 {
				memAddr, mode = addr_mode_Zeropage(PC+1)
			}
			opc_ORA( memAddr, mode, 2, 3 )

		//-------------------------------------------------- EOR --------------------------------------------------//

		case 0x49:	// Instruction EOR (immediate)
			if opc_cycle_count == 1 {
				memAddr, mode = addr_mode_Immediate(PC+1)
			}
			opc_EOR( memAddr, mode, 2, 2 )

		case 0x45:	// Instruction EOR (zeropage)
			if opc_cycle_count == 1 {
				memAddr, mode = addr_mode_Zeropage(PC+1)
			}
			opc_EOR( memAddr, mode, 2, 3 )

		//-------------------------------------------------- SHIFT --------------------------------------------------//

		case 0x0A:	// Instruction ASL (accumulator)
			opc_ASL( 1, 2 )

		case 0x4A:	// Instruction LSR (accumulator)
			opc_LSR( 1, 2 )

		//-------------------------------------------------- CMP --------------------------------------------------//

		case 0xC5:	// Instruction CMP (zeropage)
			if opc_cycle_count == 1 {
				memAddr, mode = addr_mode_Zeropage(PC+1)
			}
			opc_CMP( memAddr, mode, 2, 3 )

		case 0xC9:	// Instruction CMP (immediate)
			if opc_cycle_count == 1 {
				memAddr, mode = addr_mode_Immediate(PC+1)
			}
			opc_CMP( memAddr, mode, 2, 2 )

		//-------------------------------------------------- STA --------------------------------------------------//

		case 0x95:	// Instruction STA (zeropage,X)
			if opc_cycle_count == 1 {
				memAddr, mode = addr_mode_ZeropageX(PC+1)
			}
			opc_STA( memAddr, mode, 2, 4 )

		case 0x85:	// Instruction STA (zeropage)
			if opc_cycle_count == 1 {
				memAddr, mode = addr_mode_Zeropage(PC+1)
			}
			opc_STA( memAddr, mode, 2, 3 )

		case 0x99:	// Instruction STA (absolute,Y)
			if opc_cycle_count == 1 {
				memAddr, mode = addr_mode_AbsoluteY(PC+1)
			}
			opc_STA( memAddr, mode, 3, 5 )

		case 0x8D:	// Instruction STA (absolute)
			if opc_cycle_count == 1 {
				memAddr, mode = addr_mode_Absolute(PC+1)
			}
			opc_STA( memAddr, mode, 3, 4 )

		//-------------------------------------------------- ADC --------------------------------------------------//

		case 0x65:	// Instruction ADC (zeropage)
			if opc_cycle_count == 1 {
				memAddr, mode = addr_mode_Zeropage(PC+1)
			}
			opc_ADC( memAddr, mode, 2, 3 )

		//-------------------------------------------------- ROL --------------------------------------------------//

		case 0x26:	// Instruction ROL (zeropage)
			if opc_cycle_count == 1 {
				memAddr, mode = addr_mode_Zeropage(PC+1)
			}
			opc_ROL( memAddr, mode, 2, 5 )

		//-------------------------------------------------- ISB? FF --------------------------------------------------//

		// ISB (INC FOLLOWED BY SBC - IMPLEMENT IT!!!!!!)
		// FF (Filled ROM)
		case 0xFF:
			if Debug {
				fmt.Printf("\tOpcode %02X [1 byte]\tFilled ROM.\tPC incremented.\n", Opcode)
			}
			PC +=1

		//-------------------------------------------- No Opcode Found --------------------------------------------//

		default:
			fmt.Printf("\tOPCODE %02X NOT IMPLEMENTED!\n\n", Opcode)
			os.Exit(0)

	}

	// Increment Cycle
	Cycle ++

	// Pause = true
	// Increment Instructions per second counter
	IPS ++

}
