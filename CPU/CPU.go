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
	Opcode		uint16		// CPU Operation Code
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
			os.Exit(2)
		}

	return value

}


func MemPageBoundary(Address1, Address2 uint16) bool {

	var cross bool = false
	// if Debug {
	// 	fmt.Printf("\n\tValues: %02X %02X\n",Address1 >>8, Address2 >>8)
	// }

	// Get the High byte only to compare
	if Address1 >> 8 != Address2 >> 8 {
		cross = true
		if Debug {
			fmt.Printf("\tPC High byte: %02X\tBranch High byte: %02X\tMemory Page Boundary Cross detected! Add 1 cycle.\n",Address1 >>8, Address2 >>8)
		}
	// Omit later
	} else {
		if Debug {
			fmt.Printf("\tPC High byte: %02X\tBranch High byte: %02X\tNo Memory Page Boundary Cross detected.\n",Address1 >>8, Address2 >>8)
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


func Break() {
	// Read the Opcode from PC+1 and PC bytes (Little Endian)
	PC = uint16(Memory[0xFFFF])<<8 | uint16(Memory[0xFFFE])
	//fmt.Printf("\nRESET: %04X\n",PC)
}


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
	fmt.Printf("\n\nCycle: %d\tOpcode: %02X\tPC: 0x%02X(%d)\tA: 0x%02X\tX: 0x%02X\tY: 0x%02X\tP: %d\tSP: %02X\tGRP0: %08b\tHMP0: %02X\tBeam_index: %d", Cycle, Opcode, PC, PC, A, X, Y, P, SP, Memory[GRP0], Memory[HMP0], Beam_index )
}


// CPU Interpreter
func Interpreter() {

	// Read the Next Instruction to be executed
	Opcode = uint16(Memory[PC])

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

		case 0x60:	// Instruction RTS
			opc_RTS()

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

		//-------------------------------------------------- LDX --------------------------------------------------//

		case 0xA2:	// Instruction LDX (immediate)
			opc_LDX( addr_mode_Immediate(PC+1) )

		//-------------------------------------------------- STX --------------------------------------------------//

		case 0x86: // Instruction STX (zeropage)
			opc_STX( addr_mode_Zeropage(PC+1) )

		// STX  Store Index X in Memory (zeropage)
		//
		//      X -> M                           N Z C I D V
		//                                       - - - - - -
		//
		//      addressing    assembler    opc  bytes  cyles
		//      --------------------------------------------
		//      zeropage      STX oper      86    2     3
		// case 0x86: // STX
		//
		// 	Memory[Memory[PC+1]] = X
		// 	if Debug {
		// 		fmt.Printf("\n\tOpcode %02X%02X [2 bytes]\tSTX  Store Index X in Memory (zeropage).\tMemory[%02X] = X (%d)\n", Opcode, Memory[PC+1], Memory[PC+1], X)
		// 	}
		//
		// 	PC += 2
		// 	Beam_index += 3

		//-------------------------------------------------- LDA --------------------------------------------------//

		case 0xA9: // InstructionLDA (immediate)
			opc_LDA( addr_mode_Immediate(PC+1) )


		// LDA  Load Accumulator with Memory (zeropage)
		//
		//      M -> A                           N Z C I D V
		//                                       + + - - - -
		//
		//      addressing    assembler    opc  bytes  cyles
		//      --------------------------------------------
		//      zeropage      LDA oper      A5    2     3
		case 0xA5: // LDA (zeropage)
			opc_LDA( addr_mode_Zeropage(PC+1) )

			// A = Memory[Memory[PC+1]]
			// if Debug {
			// 	fmt.Printf("\n\tOpcode %02X%02X [2 bytes]\tLDA  Load Accumulator with Memory (zeropage).\tA = Memory[%02X] (%d)\n", Opcode, Memory[PC+1], Memory[PC+1], A)
			// }
			//
			// flags_Z(A)
			// flags_N(A)
			// PC += 2
			// Beam_index += 3
			// os.Exit(2)


		// LDA  Load Accumulator with Memory (absolute,Y)
		//
		//      M -> A                           N Z C I D V
		//                                       + + - - - -
		//
		//      addressing    assembler    opc  bytes  cyles
		//      --------------------------------------------
		//      absolute,Y    LDA oper,Y    B9    3     4*
		case 0xB9:
			tmp := uint16(Memory[PC+2])<<8 | uint16(Memory[PC+1])
			A = Memory[tmp + uint16(Y)]
			// fmt.Printf("\n\n\n%08b", A)
			// fmt.Printf("\n\n\n%d", A)
			if Debug {
				fmt.Printf("\n\tOpcode %02X %02X%02X [3 bytes]\tLDA  Load Accumulator with Memory (absolute,Y).\tA = Memory[%04X + Y(%d)]  (%d)\n", Opcode, Memory[PC+2], Memory[PC+1], tmp, Y, A)
			}

			PC += 3

			// Add 1 to cycles if page boundery is crossed
			if MemPageBoundary(uint16(tmp) + uint16(Y), PC) {
				Beam_index += 1
			}

			flags_Z(A)
			flags_N(A)
			Beam_index += 4



		// LDA  Load Accumulator with Memory (indirect),Y //
		//
		//      M -> A                           N Z C I D V
		//                                       + + - - - -
		//
		//      addressing    assembler    opc  bytes  cyles
		//      --------------------------------------------
		//      (indirect),Y  LDA (oper),Y  B1    2     5*
		case 0xB1:

			tmp  := uint16(Memory[Memory[PC+1] + 1])<<8 | uint16(Memory[Memory[PC+1]])
			tmp2 := tmp + uint16(Y)

			A = Memory[tmp2]

			if Debug {
				fmt.Printf("\n\tOpcode %02X%02X [2 bytes]\tLDA  Load Accumulator with Memory (Indirect Indexed),Y.\tA = Memory[%02X + Y(%d) = %02X] (%d)\n", Opcode, Memory[PC+1], tmp, Y, tmp2, A)
			}

			PC += 2

			// Add 1 to cycles if page boundery is crossed
			if MemPageBoundary(uint16(tmp2), PC) {
				Beam_index += 1
			}

			flags_Z(A)
			flags_N(A)

			Beam_index += 5


		//-------------------------------------------------- STA --------------------------------------------------//


		// STA  Store Accumulator in Memory (zeropage,X)
		//
		//      A -> M                           N Z C I D V
		//                                       - - - - - -
		//
		//      addressing    assembler    opc  bytes  cyles
		//      --------------------------------------------
		//      zeropage,X    STA oper,X    95    2     4
		case 0x95:

			Beam_index += 4

			Memory[Memory[PC+1]] = A

			if Debug {
				fmt.Printf("\n\tOpcode %02X%02X [2 bytes]\tSTA  Store Accumulator in Memory (zeropage, X).\tMemory[%02X] = A (%d)\n", Opcode,Memory[PC+1], Memory[PC+1], Memory[Memory[PC+1]] )
			}

			// MAPEAR SE NAO TEM QUE SOMAR O X
			// USADO NO 1 CLEANMEM
			os.Exit(2)

			// Wait for a new line and authorize graphics to draw the line
			// Wait for Horizontal Blank to draw the new line
			if Memory[PC+1] == WSYNC {
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

			if Memory[PC+1] == RESP0 {
				if Memory[RESP0] != 0 {
					XPositionP0 = Beam_index
					if Debug {
						fmt.Printf("\nRESP0 SET\tXPositionP0: %d\n", XPositionP0)
					}
				}
			}

			if Memory[PC+1] == RESP1 {
				if Memory[RESP1] != 0 {
					XPositionP1 = Beam_index
					if Debug {
						fmt.Printf("\nRESP1 SET\tXPositionP1: %d\n", XPositionP1)
					}

				}
			}


			if Memory[PC+1] == HMP0 {
				XFinePositionP0 = Fine(Memory[HMP0])

				if Debug {
					fmt.Printf("\nHMP0 SET: %d\n", XFinePositionP0)
				}

			}

			if Memory[PC+1] == HMP1 {
				XFinePositionP1 = Fine(Memory[HMP1])
				if Debug {
					fmt.Printf("\nHMP1 SET: %d\n", XFinePositionP1)
				}
			}



			PC += 2


		// STA  Store Accumulator in Memory (zeropage)
		//
		//      A -> M                           N Z C I D V
		//                                       - - - - - -
		//
		//      addressing    assembler    opc  bytes  cyles
		//      --------------------------------------------
		//      zeropage      STA oper      85    2     3
		case 0x85:
			Memory[Memory[PC+1]] = A

			if Debug {
				fmt.Printf("\n\tOpcode %02X%02X [2 bytes]\tSTA  Store Accumulator in Memory (zeropage).\tMemory[%02X] = A (%d)\n", Opcode, Memory[PC+1], Memory[PC+1], Memory[Memory[PC+1]] )
			}


			// Wait for a new line and authorize graphics to draw the line
			// Wait for Horizontal Blank to draw the new line
			if Memory[PC+1] == WSYNC {
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

			if Memory[PC+1] == RESP0 {
				if Memory[RESP0] != 0 {
					XPositionP0 = Beam_index
					if Debug {
						fmt.Printf("\nRESP0 SET\tXPositionP0: %d\n", XPositionP0)
					}
				}
			}

			if Memory[PC+1] == RESP1 {
				if Memory[RESP1] != 0 {
					XPositionP1 = Beam_index
					if Debug {
						fmt.Printf("\nRESP1 SET\tXPositionP1: %d\n", XPositionP1)
					}

				}
			}


			if Memory[PC+1] == HMP0 {
				XFinePositionP0 = Fine(Memory[HMP0])

				if Debug {
					fmt.Printf("\nHMP0 SET: %d\n", XFinePositionP0)
				}

			}

			if Memory[PC+1] == HMP1 {
				XFinePositionP1 = Fine(Memory[HMP1])
				if Debug {
					fmt.Printf("\nHMP1 SET: %d\n", XFinePositionP1)
				}
			}





			Beam_index += 3
			PC += 2


		// STA  Store Accumulator in Memory (absolute,Y)
		//
		//      A -> M                           N Z C I D V
		//                                       - - - - - -
		//
		//      addressing    assembler    opc  bytes  cyles
		//      --------------------------------------------
		//      absolute,Y    STA oper,Y    99    3     5
		case 0x99:
			tmp := uint16(Memory[PC+2])<<8 | uint16(Memory[PC+1])

			// Memory[tmp + uint16(Y)] = A
			Memory[tmp + uint16(Y)] = A

			if Debug {
				fmt.Printf("\n\tOpcode %02X %02X%02X [3 bytes]\tSTA  Store Accumulator in Memory (absolute,Y).\tA = Memory[%04X + Y(%d)] (%d)\n", Opcode, Memory[PC+2], Memory[PC+1],		tmp, Y, A )
			}

			// Check Draw related memory addresses to send instructions to TV
			if Memory[PC+1] == WSYNC {
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

			if Memory[PC+1+uint16(Y)] == RESP0 {
				if Memory[RESP0] != 0 {
					XPositionP0 = Beam_index
					// if Debug {
						fmt.Printf("\nRESP0 SET\tXPositionP0: %d\tScreen Pos: %d\n", XPositionP0, (XPositionP0*3)-68)
						fmt.Printf("\nJetXPos: %d", Memory[0x80])
					// }
				}
			}

			if Memory[PC+1] == RESP1 {
				if Memory[RESP1] != 0 {
					XPositionP1 = Beam_index
					if Debug {
						fmt.Printf("\nRESP1 SET\tXPositionP1: %d\n", XPositionP1)
					}

				}
			}


			if Memory[PC+1+uint16(Y)] == HMP0 {
				XFinePositionP0 = Fine(Memory[HMP0])

				// if Debug {
					fmt.Printf("\nHMP0 SET: %d\n", XFinePositionP0)
				// }

			}

			if Memory[PC+1+uint16(Y)] == HMP1 {
				XFinePositionP1 = Fine(Memory[HMP1])
				if Debug {
					fmt.Printf("\nHMP1 SET: %d\n", XFinePositionP1)
				}
			}

			PC += 3
			Beam_index += 5


		//-------------------------------------------------- LDY --------------------------------------------------//


		// LDY  Load Index Y with Memory (immidiate)
		//
		//      M -> Y                           N Z C I D V
		//                                       + + - - - -
		//
		//      addressing    assembler    opc  bytes  cyles
		//      --------------------------------------------
		//      immidiate     LDY #oper     A0    2     2
		case 0xA0: // LDY immidiate
			Y = Memory[PC+1]
			if Debug {
				fmt.Printf("\n\tOpcode %02X%02X [2 bytes]\tLDY  Load Index y with Memory (immidiate).\tY = Memory[%02X] (%d)\n", Opcode, Memory[PC+1], PC+1, Y)
			}
			PC += 2

			flags_Z(Y)
			flags_N(Y)
			Beam_index += 2


		// LDY  Load Index Y with Memory (zeropage) //  Same as absolute but in the first page
		//
		//      M -> Y                           N Z C I D V
		//                                       + + - - - -
		//
		//      addressing    assembler    opc  bytes  cyles
		//      --------------------------------------------
		//      zeropage      LDY oper      A4    2     3
		case 0xA4:
			Y = Memory[PC+1]
			if Debug {
				fmt.Printf("\n\tOpcode %02X%02X [2 bytes]\tLDY  Load Index y with Memory (zeropage).\tY = Memory[%02X](%02X)\t(%d)\n", Opcode, Memory[PC+1], PC+1, Memory[PC+1], Y)
			}
			PC += 2


			flags_Z(Y)
			flags_N(Y)
			Beam_index += 3




		//-------------------------------------------------- STY --------------------------------------------------//


		// STY  Store Index Y in Memory (zeropage)
		//
		//      Y -> M                           N Z C I D V
		//                                       - - - - - -
		//
		//      addressing    assembler    opc  bytes  cyles
		//      --------------------------------------------
		//      zeropage      STY oper      84    2     3
		case 0x84: // STY
			Memory[Memory[PC+1]] = Y
			if Debug {
				fmt.Printf("\n\tOpcode %02X%02X [2 bytes]\tSTY  Store Index Y in Memory (zeropage).\tMemory[%02X] = Y (%d)\n", Opcode, Memory[PC+1], Memory[PC+1], Y)
			}

			PC += 2
			Beam_index += 3


		//-------------------------------------------------- JMP --------------------------------------------------//


		// JMP  Jump to New Location (absolute)
		//
		//      (PC+1) -> PCL                    N Z C I D V
		//      (PC+2) -> PCH                    - - - - - -
		//
		//      addressing    assembler    opc  bytes  cyles
		//      --------------------------------------------
		//      absolute      JMP oper      4C    3     3
		case 0x4C:
			tmp := uint16(Memory[PC+2])<<8 | uint16(Memory[PC+1])
			//fmt.Printf("\n%04X\n",tmp)
			if Debug {
				fmt.Printf("\n\tOpcode %02X %02X%02X [3 bytes]\tJMP  Jump to New Location (absolute).\tPC = Memory[%02X]\n", Opcode, Memory[PC+2], Memory[PC+1], tmp)
			}
			PC = tmp
			Beam_index += 3

		// ISB (INC FOLLOWED BY SBC - IMPLEMENT IT!!!!!!)
		// FF (Filled ROM)
		case 0xFF:
			if Debug {
				fmt.Printf("\n\tOpcode %02X [1 byte]\tFilled ROM.\tPC incremented.\n", Opcode)
			}
			PC +=1


		//-------------------------------------------------- JMP --------------------------------------------------//

		// JSR  Jump to New Location Saving Return Address
		//
		//      push (PC+2) to Stack,            N Z C I D V
		//      (PC+1) -> PCL                    - - - - - -
		//      (PC+2) -> PCH
		//
		//      addressing    assembler    opc  bytes  cyles
		//      --------------------------------------------
		//      absolute      JSR oper      20    3     6
		case 0x20:
		    	// Push PC+2 (will be increased in 1 in RTS to match the next address (3 bytes operation))
		    	// Store the first byte into the Stack
		    	Memory[SP] = byte( (PC+2) >> 8 )
		    	SP--
		    	// Store the second byte into the Stack
		    	Memory[SP] = byte( (PC+2) & 0xFF )
		    	SP--
		    	// fmt.Printf("\nPC+3: %02X",PC+3)
		    	// fmt.Printf("\nF0: %02X",(PC+3) >> 8)
		    	// fmt.Printf("\n42: %02X",(PC+3) & 0xFF)

		    	tmp := uint16(Memory[PC+2]) <<8 | uint16(Memory[PC+1])
		    	// fmt.Printf("\n%04X\n",tmp)
		    	if Debug {
		    		fmt.Printf("\n\tOpcode %02X %02X%02X [3 bytes]\tJSR  Jump to New Location Saving Return Address.\tPC = Memory[%02X]\t|\t Stack[%02X] = %02X\t Stack[%02X] = %02X\n", Opcode, Memory[PC+2], Memory[PC+1], tmp, SP+2, Memory[SP+2], SP+1, Memory[SP+1])
		    	}

		    	PC = tmp
		    	Beam_index += 6


		//-------------------------------------------------- CPY --------------------------------------------------//


		// CPY  Compare Memory and Index Y (immidiate)
		//
		//      Y - M                            N Z C I D V
		//                                       + + + - - -
		//
		//      addressing    assembler    opc  bytes  cyles
		//      --------------------------------------------
		//      immidiate     CPY #oper     C0    2     2
		case 0xC0:

			tmp := Y - Memory[PC+1]

			if Debug {
				if tmp == 0 {
					fmt.Printf("\n\tOpcode %02X%02X [2 bytes]\tCPY  Compare Memory and Index Y (immidiate).\tY(%d) - Memory[%02X](%d) = (%d) EQUAL\n", Opcode, Memory[PC+1], Y, PC+1, Memory[PC+1], tmp)
				} else {
					fmt.Printf("\n\tOpcode %02X%02X [2 bytes]\tCPY  Compare Memory and Index Y (immidiate).\tY(%d) - Memory[%02X](%d) = (%d) NOT EQUAL\n", Opcode, Memory[PC+1], Y, PC+1, Memory[PC+1], tmp)
				}
			}

			flags_Z(tmp)
			flags_N(tmp)
			flags_C(Y,Memory[PC+1])

			PC += 2
			Beam_index += 2


		// CPY  Compare Memory and Index Y (zeropage)
		//
		//      Y - M                            N Z C I D V
		//                                       + + + - - -
		//
		//      addressing    assembler    opc  bytes  cyles
		//      --------------------------------------------
		//      zeropage      CPY oper      C4    2     3
		case 0xC4:

			tmp := Y - Memory[Memory[PC+1]]

			if Debug {
				if tmp == 0 {
					fmt.Printf("\n\tOpcode %02X%02X [2 bytes]\tCPY  Compare Memory and Index Y (zeropage).\tY(%d) - Memory[%02X](%d) = (%d) EQUAL\n", Opcode, Memory[PC+1], Y, PC+1, Memory[PC+1], tmp)
				} else {
					fmt.Printf("\n\tOpcode %02X%02X [2 bytes]\tCPY  Compare Memory and Index Y (zeropage).\tY(%d) - Memory[%02X](%d) = (%d) NOT EQUAL\n", Opcode, Memory[PC+1], Y, PC+1, Memory[PC+1], tmp)
				}
			}

			flags_Z(tmp)
			flags_N(tmp)
			flags_C(Y,Memory[PC+1])

			PC += 2
			Beam_index += 3


		//-------------------------------------------------- SBC --------------------------------------------------//


		// SBC  Subtract Memory from Accumulator with Borrow (zeropage)
		//
	     // A - M - C -> A                   N Z C I D V
	     //                                  + + + - - +
		//
	     // addressing    assembler    opc  bytes  cyles
	     // --------------------------------------------
	     // zeropage      SBC oper      E5    2     3
		case 0xE5:

			// Inverted Carry
			var borrow byte = P[0] ^ 1

			tmp := A

			if Debug {
				fmt.Printf("\n\tOpcode %02X%02X [2 bytes]\tSBC  Subtract Memory from Accumulator with Borrow (zeropage).\tA = A(%d) - Memory[Memory[%02X]](%d) - (Borrow(Inverted Carry)(%d)) = %d\n", Opcode, Memory[PC+1],		A, PC+1, Memory[Memory[PC+1]], borrow , A - Memory[Memory[PC+1]] - borrow )
			}

			// Result
			A = A - Memory[Memory[PC+1]] - borrow

			// For the flags:
			// The subtraction is VALUE1 (A) - VALUE2 (Memory[PC+1] - (P[0] ^ 1)
			value2 := Memory[PC+1] - borrow

			// First V because it need the original carry flag value
			Flags_V_SBC(tmp, value2)
			// After, update the carry flag value
			flags_C_Subtraction(tmp, value2)

			// // Clear Carry if overflow in bit 7 // NOT NECESSARY
			// if P[6] == 1 {
			// 	fmt.Printf("\n\tCarry cleared due to an overflow!")
			// 	P[0] = 0
			// }

			flags_Z(A)
			flags_N(A)

			PC += 2
			Beam_index += 3


		// SBC  Subtract Memory from Accumulator with Borrow (immidiate)
		//
		//      A - M - C -> A                   N Z C I D V
		//                                       + + + - - +
		//
		//      addressing    assembler    opc  bytes  cyles
		//      --------------------------------------------
		//      immidiate     SBC #oper     E9    2     2
		case 0xE9:

			// Inverted Carry
			var borrow byte = P[0] ^ 1

			// Original value of A
			tmp := A

			if Debug {
				fmt.Printf("\n\tOpcode %02X%02X [2 bytes]\tSBC  Subtract Memory from Accumulator with Borrow (immidiate).\tA = A(%d) - Memory[%02X](%d) - (Borrow(Inverted Carry)(%d)) = %d\n", Opcode, Memory[PC+1], A, PC+1, Memory[PC+1], borrow , A - Memory[PC+1] - borrow )
			}

			// Result
			A = A - Memory[PC+1] - borrow

			// For the flags:
			// The subtraction is VALUE1 (A) - VALUE2 (Memory[PC+1] - (P[0] ^ 1)
			value2 := Memory[PC+1] - borrow

			// First V because it need the original carry flag value
			Flags_V_SBC(tmp, value2)
			// After, update the carry flag value
			flags_C_Subtraction(tmp, value2)

			// // Clear Carry if overflow in bit 7 // NOT NECESSARY
			// if P[6] == 1 {
			// 	fmt.Printf("\n\tCarry cleared due to an overflow!")
			// 	P[0] = 0
			// }

			flags_Z(A)
			flags_N(A)

			PC += 2
			Beam_index += 2


		//-------------------------------------------------- ADC --------------------------------------------------//
		// ADC  Add Memory to Accumulator with Carry (zeropage)
		//
		// 	A + M + C -> A, C                N Z C I D V
		// 	                          	   + + + - - +
		//
		// 	addressing    assembler    opc  bytes  cyles
		// 	--------------------------------------------
		// 	zeropage      ADC oper      65    2     3
		case 0x65:

			// Original value of A
			tmp := A

			if Debug {
				fmt.Printf("\n\tOpcode %02X%02X [2 bytes]\tADC  Add Memory to Accumulator with Carry (zeropage).\tA = A(%d) + Memory[Memory[%02X]](%d) + Carry (%d)) = %d\n", Opcode, Memory[PC+1],		A, PC+1, Memory[Memory[PC+1]], P[0] , A + Memory[Memory[PC+1]] + P[0] )
			}

			// Result
			A = A + Memory[Memory[PC+1]] + P[0]
			// A = A + Memory[PC+1] + P[0]

			// For the flags:
			// The addiction is VALUE1 (A) - VALUE2 (Memory[Memory[PC+1]] + P[0])
			// value2 := Memory[Memory[PC+1]] + P[0]

			// First V because it need the original carry flag value
			Flags_V_ADC(tmp, A)
			// After, update the carry flag value
			flags_C(tmp, A)

			flags_Z(A)
			flags_N(A)

			PC += 2
			Beam_index += 3
			// os.Exit(2)
			// Pause = true



		//-------------------------------------------------- CMP --------------------------------------------------//


		// CMP  Compare Memory with Accumulator
		//
		//      A - M                          N Z C I D V
		//                                     + + + - - -
		//
		//      addressing    assembler    opc  bytes  cyles
		//      --------------------------------------------
		//      zeropage      CMP oper      C5    2     3
		case 0xC5:

			// WORKAROUND
			// If in a RAM range, read the value from RAM Address [Memory [Memory[PC+1]] ]
			if Memory[PC+1] >= 0x80 && Memory[PC+1] <= 0xFF {
				tmp := A - Memory[Memory[PC+1]]

				if Debug {
					if tmp == 0 {
						fmt.Printf("\n\tOpcode %02X%02X [2 bytes]\tCMP  Compare Memory with Accumulator (zeropage).\tA(%d) - Memory[%02X](%d) = (%d) EQUAL\n", Opcode, Memory[PC+1], A, Memory[PC+1], Memory[Memory[PC+1]], tmp)
					} else {
						fmt.Printf("\n\tOpcode %02X%02X [2 bytes]\tCMP  Compare Memory with Accumulator (zeropage).\tA(%d) - Memory[%02X](%d) = (%d) NOT EQUAL\n", Opcode, Memory[PC+1], A, Memory[PC+1], Memory[Memory[PC+1]], tmp)
					}
				}
				flags_Z(tmp)
				flags_N(tmp)
				flags_C_Subtraction(A,Memory[Memory[PC+1]])


				PC += 2
				Beam_index += 3

			// Else, read directly the value in Memory [PC+1]
			} else {
				// tmp := A - Memory[Memory[PC+1]]
				tmp := A - Memory[PC+1]

				if Debug {
					if tmp == 0 {
						// fmt.Printf("\n\tOpcode %02X%02X [2 bytes]\tCMP  Compare Memory with Accumulator (zeropage).\tA(%d) - Memory[%02X](%d) = (%d) EQUAL\n", Opcode, Memory[PC+1], A, Memory[PC+1], Memory[Memory[PC+1]], tmp)
						fmt.Printf("\n\tOpcode %02X%02X [2 bytes]\tCMP  Compare Memory with Accumulator (zeropage).\tA(%d) - %d = (%d) EQUAL\n", Opcode, Memory[PC+1], A, Memory[PC+1], tmp)
					} else {
						// fmt.Printf("\n\tOpcode %02X%02X [2 bytes]\tCMP  Compare Memory with Accumulator (zeropage).\tA(%d) - Memory[%02X](%d) = (%d) NOT EQUAL\n", Opcode, Memory[PC+1], A, Memory[PC+1], Memory[Memory[PC+1]], tmp)
						fmt.Printf("\n\tOpcode %02X%02X [2 bytes]\tCMP  Compare Memory with Accumulator (zeropage).\tA(%d) - %d = (%d) NOT EQUAL\n", Opcode, Memory[PC+1], A, Memory[PC+1], tmp)
					}
				}
				flags_Z(tmp)
				flags_N(tmp)
				flags_C_Subtraction(A,Memory[PC+1])

				PC += 2
				Beam_index += 3
			}


		// CMP  Compare Memory with Accumulator (immidiate)
		//
		//      A - M                            N Z C I D V
		//                                       + + + - - -
		//
		//      addressing    assembler    opc  bytes  cyles
		//      --------------------------------------------
		//      immidiate     CMP #oper     C9    2     2
		case 0xC9:
				tmp := A - Memory[PC+1]

				if Debug {
					if tmp == 0 {
						fmt.Printf("\n\tOpcode %02X%02X [2 bytes]\tCMP  Compare Memory with Accumulator (immidiate).\tA(%d) - %d = (%d) EQUAL\n", Opcode, Memory[PC+1], A, Memory[PC+1], tmp)
					} else {
						fmt.Printf("\n\tOpcode %02X%02X [2 bytes]\tCMP  Compare Memory with Accumulator (immidiate).\tA(%d) - %d = (%d) NOT EQUAL\n", Opcode, Memory[PC+1], A, Memory[PC+1], tmp)
					}
				}

				flags_Z(tmp)
				flags_N(tmp)
				flags_C_Subtraction(A,Memory[PC+1])

				PC += 2
				Beam_index += 2

		//-------------------------------------------------- DEC --------------------------------------------------//


		// DEC  Decrement Memory by One
		//
		//      M - 1 -> M                       N Z C I D V
		//                                       + + - - - -
		//
		//      addressing    assembler    opc  bytes  cyles
		//      --------------------------------------------
		//      zeropage      DEC oper      C6    2     5
		case 0xC6: // DEC

			if Debug {
				fmt.Printf("\n\tOpcode %02X%02X [2 bytes]\tDEC  Decrement Memory by One.\tMemory[%02X] -= 1 (%d)\n", Opcode, Memory[PC+1], Memory[PC+1], Memory[Memory[PC+1]] - 1 )
			}
			Memory[Memory[PC+1]] -= 1

			flags_Z(Memory[Memory[PC+1]])
			flags_N(Memory[Memory[PC+1]])

			PC += 2
			Beam_index += 5


		//-------------------------------------------------- AND --------------------------------------------------//

		// AND  AND Memory with Accumulator (immidiate)
		//
		//      A AND M -> A                     N Z C I D V
		//                                       + + - - - -
		//
		//      addressing    assembler    opc  bytes  cyles
		//      --------------------------------------------
		//      immidiate     AND #oper     29    2     2
		case 0x29: // AND (immidiate)

			if Debug {
				fmt.Printf("\n\tOpcode %02X%02X [2 bytes]\tAND  AND Memory with Accumulator (immidiate).\tA = A(%d) & Memory[%02X](%d)\t(%d)\n", Opcode, Memory[PC+1], A, PC+1, Memory[PC+1], A & Memory[PC+1] )
			}

			A = A & Memory[PC+1]

			flags_Z(A)
			flags_N(A)

			PC += 2
			Beam_index += 2


		//-------------------------------------------------- AND --------------------------------------------------//


		// EOR  Exclusive-OR Memory with Accumulator (immidiate)
		//
		//      A EOR M -> A                     N Z C I D V
		//                                       + + - - - -
		//
		//      addressing    assembler    opc  bytes  cyles
		//      --------------------------------------------
		//      immidiate     EOR #oper     49    2     2
		case 0x49:

			if Debug {
				fmt.Printf("\n\tOpcode %02X%02X [2 bytes]\tEOR  Exclusive-OR Memory with Accumulator (immidiate).\tA = A(%d) XOR Memory[%02X](%d)\t(%d)\n", Opcode, Memory[PC+1], A, PC+1, Memory[PC+1], A ^ Memory[PC+1] )
			}

			A = A ^ Memory[PC+1]

			flags_Z(A)
			flags_N(A)

			PC += 2
			Beam_index += 2



		//-------------------------------------------------- EOR --------------------------------------------------//


		// ASL  Shift Left One Bit (Memory or Accumulator) (accumulator)
		//
		//      C <- [76543210] <- 0             N Z C I D V
		//                                       + + + - - -
		//
		//      addressing    assembler    opc  bytes  cyles
		//      --------------------------------------------
		//      accumulator   ASL A         0A    1     2
		case 0x0A:

			if Debug {
				fmt.Printf("\n\tOpcode %02X [1 byte]\tASL  Shift Left One Bit (Memory or Accumulator) (accumulator).\tA = A(%d) Shift Left 1 bit\t(%d)\n", Opcode, A, A << 1 )
			}

			flags_C(A, A << 1)

			A = A << 1

			flags_N(A)
			flags_Z(A)

			PC += 1
			Beam_index += 2


		//-------------------------------------------------- BIT --------------------------------------------------//

		// BIT  Test Bits in Memory with Accumulator
		//
		//      bits 7 and 6 of operand are transfered to bit 7 and 6 of SR (N,V);
		//      the zeroflag is set to the result of operand AND accumulator.
		//
		//      A AND M, M7 -> N, M6 -> V        N Z C I D V
		//                                      M7 + - - - M6
		//
		//      addressing    assembler    opc  bytes  cyles
		//      --------------------------------------------
		//      absolute      BIT oper      2C    3     4
		case 0x2C:
			// Memory address
			tmp := uint16(Memory[PC+2])<<8 | uint16(Memory[PC+1])

			if Debug {
				fmt.Printf("\n\tOpcode %02X %02X%02X [3 bytes]\tBIT  Test Bits in Memory with Accumulator (absolute).\tA (%08b) AND Memory[%04X] (%08b) = %08b \tM7 -> N, M6 -> V\n", Opcode, Memory[PC+2], Memory[PC+1], A, tmp, Memory[tmp], A & Memory[tmp] )
			}
			// fmt.Printf("\n\n%08b\n\n",A & Memory[tmp])

			// Memory Address bit 7 (A) -> N (Negative)
			// A = 0x80
			if Debug {
				fmt.Printf("\n\tFlag N: %d -> ",P[7])
			}
			P[7] = A >> 7 & 0x1
			if Debug {
				fmt.Printf("%d\n",P[7])
			}

			// Memory Address bit 6 (A) -> V (oVerflow)
			// A = 0x40
			if Debug {
				fmt.Printf("\n\tFlag V: %d -> ",P[6])
			}
			P[6] = A >> 6 & 0x1
			if Debug {
				fmt.Printf("%d\n",P[6])
			}

			flags_Z(A & Memory[tmp])

			PC += 3
			Beam_index += 4


		default:
			fmt.Printf("\n\tOPCODE %X NOT IMPLEMENTED!\n\n", Opcode)
			os.Exit(2)

	}

	Cycle ++

}
