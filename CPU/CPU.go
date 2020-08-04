package CPU

import (
	"fmt"
	"os"
	"time"
)

// Components
var (

	Memory		[65536]byte	// Memory
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
	Opcode		byte			// CPU Operation Code
	Cycle		uint16		// CPU Cycle

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
	// Draw instuctions
	DrawLine		bool = false	// Instruct Graphics to draw a new line
	DrawP0		bool = false	// Instruct Graphics to draw Player 0 sprite
	DrawP1		bool = false	// Instruct Graphics to draw Player 1 sprite

	// Pause
	Pause		bool = false

	//Debug
	Debug 		bool = true
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

	//------------------- 0280-0297 - RIOT (I/O, Timer)
	SWCHA			uint16 = 0x280		// Port A data register for joysticks: Bits 4-7 for player 1.  Bits 0-3 for player 2.

)

func testAction(memAddr uint16) {
	// Wait for a new line and authorize graphics to draw the line
	// Wait for Horizontal Blank to draw the new line

	// TODO: when tested, maybe tranform this to byte???
	// address := byte(memAddr)

	if memAddr == uint16(WSYNC) {
		if Debug {
			fmt.Printf("\nWSYNC SET\n")
		}
		DrawLine = true
		Beam_index = 0

		if Memory[GRP0] != 0 {
			if Debug {
				fmt.Printf("\nGRP0 SET\n")
			}
			DrawP0 = true
		}

		if Memory[GRP1] != 0 {
			if Debug {
				fmt.Printf("\nGRP1 SET\n")
			}
			DrawP1 = true
		}

	}

	if memAddr == uint16(RESP0) {
		if Memory[RESP0] != 0 {
			XPositionP0 = Beam_index
			if Debug {
				fmt.Printf("\nRESP0 SET\tXPositionP0: %d\n", XPositionP0)
			}
		}
	}

	if memAddr == uint16(RESP1) {
		if Memory[RESP1] != 0 {
			XPositionP1 = Beam_index
			if Debug {
				fmt.Printf("\nRESP1 SET\tXPositionP1: %d\n", XPositionP1)
			}

		}
	}


	if memAddr == uint16(HMP0) {
		XFinePositionP0 = Fine(Memory[HMP0])

		if Debug {
			fmt.Printf("\nHMP0 SET: %d\n", XFinePositionP0)
		}

	}

	if memAddr == uint16(HMP1) {
		XFinePositionP1 = Fine(Memory[HMP1])
		if Debug {
			fmt.Printf("\nHMP1 SET: %d\n", XFinePositionP1)
		}
	}
}



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
	Memory		= [65536]byte{}
	// Clean CPU Variables
	PC			= 0
	Opcode		= 0
	Cycle		= 0

	// Start Timers
	Clock		= time.NewTicker(time.Nanosecond)	// CPU Clock
	// Clock		= time.NewTicker(time.Second/60)	// CPU Clock
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
		Show()
	}

	// Map Opcode
	switch Opcode {

		//-------------------------------------------------- Implied --------------------------------------------------//

		case 0x78:	// Instruction SEI
			opc_SEI()

		case 0x38:	// Instruction SEC
			opc_SEC()

		case 0x18:	// Instruction CLC
			opc_CLC()

		case 0xD8:	// Instruction CLD
			opc_CLD()

		case 0x8A:	// Instruction TXA
			opc_TXA()

		case 0xA8:	// Instruction TAY
			opc_TAY()

		case 0xCA:	// Instruction DEX
			opc_DEX()

		case 0x88:	// Instruction DEY
			opc_DEY()

		case 0x9A:	// Instruction TXS
			opc_TXS()

		case 0x48:	// Instruction PHA
			opc_PHA()

		case 0x68:	// Instruction PLA
			opc_PLA()

		case 0x00:	// Instruction BRK
			opc_BRK()

		case 0xC8:	// Instruction INY
			opc_INY()

		case 0xE8:	// Instruction INX
			opc_INX()

		case 0x60:	// Instruction RTS
			opc_RTS()

		case 0xEA:	// Instruction NOP
			opc_NOP()

		//-------------------------------------------------- Just zeropage --------------------------------------------------//

		case 0xE6:	// Instruction INC (zeropage)
			opc_INC( addr_mode_Zeropage(PC+1) )

		//-------------------------------------------- Branches - just relative ---------------------------------------------//

		case 0xD0:	// Instruction BNE (relative)
			opc_BNE( addr_mode_Relative(PC+1) )

		case 0x90:	// Instruction BCC (relative)
			opc_BCC( addr_mode_Relative(PC+1) )

		case 0xB0:	// Instruction BCS (relative)
			opc_BCS( addr_mode_Relative(PC+1) )

		case 0x30:	// Instruction BMI (relative)
			opc_BMI( addr_mode_Relative(PC+1) )

		case 0x10:	// Instruction BPL (relative)
			opc_BPL( addr_mode_Relative(PC+1) )

		//-------------------------------------------------- LDX --------------------------------------------------//

		case 0xA2:	// Instruction LDX (immediate)
			opc_LDX( addr_mode_Immediate(PC+1) )

		//-------------------------------------------------- STX --------------------------------------------------//

		case 0x86: // Instruction STX (zeropage)
			opc_STX( addr_mode_Zeropage(PC+1) )

		//-------------------------------------------------- JMP --------------------------------------------------//

		case 0x4C:	// Instruction JMP (absolute)
			opc_JMP( addr_mode_Absolute(PC+1) )

		case 0x20:	// Instruction JSR (absolute)
			opc_JSR( addr_mode_Absolute(PC+1) )

		//-------------------------------------------------- BIT --------------------------------------------------//

		case 0x2C:	// Instruction BIT (absolute)
			opc_BIT( addr_mode_Absolute(PC+1) )

		case 0x24:	// Instruction BIT (zeropage)
			opc_BIT( addr_mode_Zeropage(PC+1) )

		//-------------------------------------------------- LDA --------------------------------------------------//

		case 0xA9:	// Instruction LDA (immediate)
			opc_LDA( addr_mode_Immediate(PC+1) )

		case 0xA5:	// Instruction LDA (zeropage)
			opc_LDA( addr_mode_Zeropage(PC+1) )

		case 0xB9:	// Instruction LDA (absolute,Y)
			opc_LDA( addr_mode_AbsoluteY(PC+1) )

		case 0xB1:	// Instruction LDA (indirect,Y)
			opc_LDA( addr_mode_IndirectY(PC+1) )

		//-------------------------------------------------- LDY --------------------------------------------------//

		case 0xA0:	// Instruction LDY (immediate)
			opc_LDY( addr_mode_Immediate(PC+1) )

		// Used by the wrong horizontal demo
		// case 0xA4:	// Instruction LDY (zeropage)
		// 	opc_LDY( addr_mode_Zeropage(PC+1) )
		// 	// os.Exit(2)

		//-------------------------------------------------- STY --------------------------------------------------//

		case 0x84:	// Instruction STY (zeropage)
			opc_STY( addr_mode_Zeropage(PC+1) )

		//-------------------------------------------------- CPY --------------------------------------------------//

		case 0xC0:	// Instruction CPY (immediate)
			opc_CPY( addr_mode_Immediate(PC+1) )

		case 0xC4:	// Instruction STY (zeropage)
			opc_CPY( addr_mode_Zeropage(PC+1) )

		//-------------------------------------------------- CPX --------------------------------------------------//

		case 0xE0:	// Instruction CPX (immediate)
			opc_CPX( addr_mode_Immediate(PC+1) )

		//-------------------------------------------------- SBC --------------------------------------------------//

		case 0xE5:	// Instruction STY (zeropage)
			opc_SBC( addr_mode_Zeropage(PC+1) )

		case 0xE9:	// Instruction STY (immediate)
			opc_SBC( addr_mode_Immediate(PC+1) )

		//-------------------------------------------------- DEC --------------------------------------------------//

		case 0xC6:	// Instruction DEC (zeropage)
			opc_DEC( addr_mode_Zeropage(PC+1) )

		//-------------------------------------------------- AND --------------------------------------------------//

		case 0x29:	// Instruction AND (immediate)
			opc_AND( addr_mode_Immediate(PC+1) )

		//-------------------------------------------------- EOR --------------------------------------------------//

		case 0x49:	// Instruction EOR (immediate)
			opc_EOR( addr_mode_Immediate(PC+1) )

		case 0x45:	// Instruction EOR (zeropage)
			opc_EOR( addr_mode_Zeropage(PC+1) )

		//-------------------------------------------------- SHIFT --------------------------------------------------//

		case 0x0A:	// Instruction ASL (accumulator)
			opc_ASL()

		case 0x4A:	// Instruction LSR (accumulator)
			opc_LSR()

		//-------------------------------------------------- CMP --------------------------------------------------//

		case 0xC5:	// Instruction CMP (zeropage)
			opc_CMP( addr_mode_Zeropage(PC+1) )

		case 0xC9:	// Instruction CMP (immediate)
			opc_CMP( addr_mode_Immediate(PC+1) )

		//-------------------------------------------------- STA --------------------------------------------------//

		// Used in 1-cleanmem
		case 0x95:	// Instruction STA (zeropage,X)
			opc_STA( addr_mode_ZeropageX(PC+1) )

		// Used in 2-colorbg, to 103-bomber
		case 0x85:	// Instruction STA (zeropage)
			opc_STA( addr_mode_Zeropage(PC+1) )

		// Used in 103-bomber
		case 0x99:	// Instruction STA (zeropage)
			opc_STA( addr_mode_AbsoluteY(PC+1) )

		//-------------------------------------------------- ADC --------------------------------------------------//

		case 0x65:	// Instruction ADC (zeropage)
			opc_ADC( addr_mode_Zeropage(PC+1) )

		//-------------------------------------------------- ROL --------------------------------------------------//

		case 0x26:	// Instruction ROL (zeropage)
			opc_ROL( addr_mode_Zeropage(PC+1) )

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

	Cycle ++

}
