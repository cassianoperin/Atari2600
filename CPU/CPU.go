package CPU

import (
	"fmt"
	"os"
	"time"
)

// Components
var (
	//0000-002C - TIA (write)
	//0030-003D - TIA (read) - (sometimes mirrored at 0030-003D)
	//0080-00FF - RIOT (RAM) (128 bytes) -- Stack uses the last addresses
	//0280-0297 - RIOT (I/O, Timer)
	//F000-FFFF - Cartridge (ROM)
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

	// Memory Positions
	WSYNC		byte = 0x02		//---- ----   Wait for Horizontal Blank
	GRP0				byte = 0x1B		//xxxx xxxx   Graphics Register Player 0
	GRP1				byte = 0x1C		//xxxx xxxx   Graphics Register Player 1
	RESP0 			byte = 0x10		//---- ----   Reset Player 0
	RESP1 			byte = 0x11		//---- ----   Reset Player 1
	COLUP0			byte = 0x06		//xxxx xxx0   Color-Luminance Player 0
	CTRLPF			byte = 0x0A		//00xx 0xxx   Control Playfield, Ball, Collisions

	// CPU Variables
	Opcode		uint16		// CPU Operation Code
	Cycle		uint16		// CPU Cycle

	// Timers
	Clock			*time.Ticker	// CPU Clock // CPU: MOS Technology 6507 @ 1.19 MHz;
	ScreenRefresh		*time.Ticker	// Screen Refresh
	Second			= time.Tick(time.Second)			// 1 second to track FPS and draws




	// *************** Personal Control Flags *************** //
	// Beam index to control where to draw objects using cpu cycles
	Beam_index	byte = 0
	// Instruct Graphics to draw a new line
	DrawLine		bool = false
	// Instruct Graphics to draw Player 0 sprite
	DrawP0		bool = false
	DrawP0VertPosition	byte = 0
	// Instruct Graphics to draw Player 1 sprite
	DrawP1		bool = false
	DrawP1VertPosition	byte = 0

	// Pause
	Pause		bool = false

	//Debug
	debug 		bool = false
)


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

	// Initialize Vertifical position of objects
	// Made by CLEAN_START macro automatically
	// Memory[RESP0] = 23
	// Memory[RESP1] = 13

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
	fmt.Printf("\n\nCycle: %d\tOpcode: %02X\tPC: 0x%02X(%d)\tA: 0x%02X\tX: 0x%02X\tY: 0x%02X\tP: %d\tSP: %02X\tStack: [%d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d]\tGRP0: %08b\tCOLUP0: %02X\tCTRLPF: %08b", Cycle, Opcode, PC, PC, A, X, Y, P, SP, Memory[0xFF], Memory[0xFE], Memory[0xFD], Memory[0xFC], Memory[0xFB], Memory[0xFA], Memory[0xF9], Memory[0xF8], Memory[0xF7], Memory[0xF6], Memory[0xF5], Memory[0xF4], Memory[0xF3], Memory[0xF2], Memory[0xF1], Memory[0xF0], Memory[GRP0], Memory[COLUP0], Memory[CTRLPF] )
}


//-------------------------------------------------- Processor Flags --------------------------------------------------//

// Zero Flag
func flags_Z(value byte) {
	if debug {
		fmt.Printf("\n\tFlag Z: %d ->", P[1])
	}
	// Check if final value is 0
	if value == 0 {
		P[1] = 1
	} else {
		P[1] = 0
	}
	if debug {
		fmt.Printf(" %d", P[1])
	}
}

// Negative Flag
func flags_N(value byte) {
	if debug {
		fmt.Printf("\n\tFlag N: %d ->", P[7])
	}
	// Set Negtive flag to the the MSB of the value
	P[7] = value >> 7

	if debug {
		fmt.Printf(" %d | Value = %08b", P[7], value)
	}
}

// Carry Flag
func flags_C(value1, value2 byte) {
	if debug {
		fmt.Printf("\n\tFlag C: %d ->", P[0])
	}



	// Check if final value is 0
	if value1 >= value2 {
		P[0] = 1
	} else {
		P[0] = 0
	}


	if debug {
		fmt.Printf(" %d", P[0])
	}
}

// oVerflow Flag for ADC
func Flags_V_ADC(value1, value2 byte) {
	var (
		carry_bit		[8]byte
		carry_OUT 	byte = 0
	)
	// fmt.Printf("\n  %08b\t%d",value1,value1)
	// fmt.Printf("\n  %08b\t%d",value2,value2)

	if debug {
		fmt.Printf("\n\tFlag V: %d ->", P[6])
	}

	// Make the magic
	for i:=0 ; i <= 7 ; i++{
		// sum the bit from value one + bit from value 2 + carry value
		tmp := (value1 >> byte(i) & 0x01) + (value2 >> byte(i) & 0x01) + carry_bit[i]
		if tmp >= 2 {
			// set the carry out
			if i == 7 {
				carry_OUT = 1
			} else {
				carry_bit[i+1] = 1
			}
		}
	}

	// fmt.Printf("\n%d ",carry_OUT)
	// for i:=7 ; i >=0 ; i--{
	// 	fmt.Printf("%d",carry_bit[i])
	// }
	// fmt.Printf("\n  %08b (soma)\tDecimal: %d",value1+value2,value1+value2)

	// Formula to calculate: V = C6 xor C7
	P[6] = carry_bit[7] ^ carry_OUT
	// fmt.Printf("\nV: %d", P[6])

	if debug {
		fmt.Printf(" %d", P[6])
	}
}

// oVerflow Flag for SBC
func Flags_V_SBC(value1, value2 byte) {
	var (
		carry_bit		[8]byte
		carry_OUT 	byte = 0
	)
	// fmt.Printf("\n  %08b\t%d",value1,value1)
	// fmt.Printf("\n  %08b\t%d",value2,value2)

	// Since internall subtraction is just addition of the ones-complement
	// N can simply be replaced by 255-N in the formulas of Flags_V_ADC
	value2 = 255-value2
	// Set the carry flag on bit 0 of carry_bit Array to bring the carry if exists
	carry_bit[0] = P[0]

	if debug {
		fmt.Printf("\n\tFlag V: %d ->", P[6])
	}

	// Make the magic
	for i:=0 ; i <= 7 ; i++{
		// sum the bit from value one + bit from value 2 + carry value
		tmp := (value1 >> byte(i) & 0x01) + (value2 >> byte(i) & 0x01) + carry_bit[i]
		if tmp >= 2 {
			// set the carry out
			if i == 7 {
				carry_OUT = 1
			} else {
				carry_bit[i+1] = 1
			}
		}
	}

	// fmt.Printf("\n%d ",carry_OUT)
	// for i:=7 ; i >=0 ; i--{
	// 	fmt.Printf("%d",carry_bit[i])
	// }
	// fmt.Printf("\n  %08b (soma)\tDecimal: %d",value1+value2,value1+value2)

	// Formula to calculate: V = C6 xor C7
	P[6] = carry_bit[7] ^ carry_OUT
	// fmt.Printf("\nV: %d", P[6])

	if debug {
		fmt.Printf(" %d", P[6])
	}
}



// CPU Interpreter
func Interpreter() {

	// Read the Next Instruction to be executed
	Opcode = uint16(Memory[PC])

	// Print Cycle and Debug Information
	if debug {
		Show()
	}

	// Map Opcode
	switch Opcode {

		//-------------------------------------------------- Unique Memory Addressing --------------------------------------------------//

		//SEI  Set Interrupt Disable Status
		//
		// 1 -> I                           N Z C I D V
		//                                  - - - 1 - -
		//
		// addressing    assembler    opc  bytes  cyles
		// --------------------------------------------
		// implied       SEI           78    1     2
		case 0x78: // SEI
			P[2]	=  1
			PC += 1
			if debug {
				fmt.Printf("\n\tOpcode %02X [1 byte]\tSEI  Set Interrupt Disable Status.\tP[2]=1\n", Opcode)
			}
			Beam_index += 2


		// SEC  Set Carry Flag
		//
		// 1 -> C                           N Z C I D V
		//                                  - - 1 - - -
		//
		// addressing    assembler    opc  bytes  cyles
		// --------------------------------------------
		// implied       SEC           38    1     2
		case 0x38: // SEC
			P[0]	=  1
			PC += 1
			if debug {
				fmt.Printf("\n\tOpcode %02X [1 byte]\tSEC  Set Carry Flag.\tP[0]=1\n", Opcode)
			}
			Beam_index += 2


		// CLD  Clear Decimal Mode
		//
		//      0 -> D                           N Z C I D V
		//                                       - - - - 0 -
		//
		//      addressing    assembler    opc  bytes  cyles
		//      --------------------------------------------
		//      implied       CLD           D8    1     2
		case 0xD8: // CLD
			P[3]	=  0
			PC += 1
			if debug {
				fmt.Printf("\n\tOpcode %02X [1 byte]\tCLD  Clear Decimal Mode.\tP[3]=0\n", Opcode)
			}
			Beam_index += 2


		// TXA  Transfer Index X to Accumulator
		//
		//      X -> A                           N Z C I D V
		//                                       + + - - - -
		//
		//      addressing    assembler    opc  bytes  cyles
		//      --------------------------------------------
		//      implied       TXA           8A    1     2
		case 0x8A: // TXA
			A = X
			if debug {
				fmt.Printf("\n\tOpcode %02X [1 byte]\tTXA  Transfer Index X to Accumulator.\tA = X (%d)\n", Opcode, X)
			}
			PC += 1
			flags_Z(A)
			flags_N(A)
			Beam_index += 2


		// TAY  Transfer Accumulator to Index Y
		//
		//      A -> Y                           N Z C I D V
		//                                       + + - - - -
		//
		//      addressing    assembler    opc  bytes  cyles
		//      --------------------------------------------
		//      implied       TAY           A8    1     2
		case 0xA8: // TAY
			Y = A
			if debug {
				fmt.Printf("\n\tOpcode %02X [1 byte]\tTAY  Transfer Accumulator to Index Y.\tA = X (%d)\n", Opcode, X)
			}
			PC += 1
			flags_Z(Y)
			flags_N(Y)
			Beam_index += 2


		// DEX  Decrement Index X by One
		//
		//      X - 1 -> X                       N Z C I D V
		//                                       + + - - - -
		//
		//      addressing    assembler    opc  bytes  cyles
		//      --------------------------------------------
		//      implied       DEC           CA    1     2
		case 0xCA: // DEX
			X--
			if debug {
				fmt.Printf("\n\tOpcode %02X [1 byte]\tDEX  Decrement Index X by One.\tX-- (%d)\n", Opcode, X)
			}
			PC += 1
			flags_Z(X)
			flags_N(X)
			Beam_index += 2

		// DEY  Decrement Index Y by One
		//
		//      Y - 1 -> Y                       N Z C I D V
		//                                       + + - - - -
		//
		//      addressing    assembler    opc  bytes  cyles
		//      --------------------------------------------
		//      implied       DEC           88    1     2
		case 0x88:
			Y--
			if debug {
				fmt.Printf("\n\tOpcode %02X [1 byte]\tDEY  Decrement Index Y by One.\tY-- (%d)\n", Opcode, Y)
			}
			PC += 1
			flags_Z(Y)
			flags_N(Y)
			Beam_index += 2


		// TXS  Transfer Index X to Stack Register
		//
		//      X -> SP                          N Z C I D V
		//                                       - - - - - -
		//
		//      addressing    assembler    opc  bytes  cyles
		//      --------------------------------------------
		//      implied       TXS           9A    1     2
		case 0x9A: // TXS
			SP	= X
			if debug {
				fmt.Printf("\n\tOpcode %02X [1 bytes]\tTXS  Transfer Index X to Stack Pointer.\tSP = X (%d)\n", Opcode, SP)
			}
			PC += 1
			Beam_index += 2


		// PHA  Push Accumulator on Stack
		//
		//      push A                           N Z C I D V
		//                                       - - - - - -
		//
		//      addressing    assembler    opc  bytes  cyles
		//      --------------------------------------------
		//      implied       PHA           48    1     3
		case 0x48: // PHA
			Memory[SP]	= A
			if debug {
				fmt.Printf("\n\tOpcode %02X [1 byte]\tPHA  Push Accumulator on Stack.\tMemory[%02X] = A (%d) | SP--\n", Opcode, SP, Memory[SP])
			}
			PC += 1
			SP--
			Beam_index += 3


		// BRK  Force Break
		//
		//      interrupt,                       N Z C I D V
		//      push PC+2, push SR               - - - 1 - -
		//
		//      addressing    assembler    opc  bytes  cyles
		//      --------------------------------------------
		//      implied       BRK           00    1     7
		case 0x00: //BRK
			if debug {
				fmt.Printf("\n\tOpcode %02X [1 byte]\tBRK  Force Break.\tBREAK!\n", Opcode)
			}
			// IRQ Enabled
			P[2] = 1
			Break()
			Beam_index += 7


		// INY  Increment Index Y by One
		//
		//      Y + 1 -> Y                       N Z C I D V
		//                                       + + - - - -
		//
		//      addressing    assembler    opc  bytes  cyles
		//      --------------------------------------------
		//      implied       INY           C8    1     2
		case 0xC8:
			Y++
			if debug {
				fmt.Printf("\n\tOpcode %02X [1 byte]\tINY  Increment Index Y by One (%02X)\n", Opcode, Y)
			}
			flags_Z(Y)
			flags_N(Y)
			PC++
			Beam_index += 2


		// INC  Increment Memory by One
		//
		//      M + 1 -> M                       N Z C I D V
		//                                       + + - - - -
		//
		//      addressing    assembler    opc  bytes  cyles
		//      --------------------------------------------
		//      zeropage      INC oper      E6    2     5
		case 0xE6:

			if debug {
				fmt.Printf("\n\tOpcode %02X%02X [2 bytes]\tINC  Increment Memory[%02X](%d) by One (%d)\n", Opcode, Memory[PC+1], Memory[PC+1], Memory[Memory[PC+1]], Memory[Memory[PC+1]] + 1)
			}

			Memory[Memory[PC+1]] = Memory[Memory[PC+1]] + 1

			flags_Z(Memory[Memory[PC+1]])
			flags_N(Memory[Memory[PC+1]])
			PC+=2
			Beam_index += 5


		//-------------------------------------------------- Branches --------------------------------------------------//


		// BNE  Branch on Result not Zero
		//
		//      branch on Z = 0                  N Z C I D V
		//                                       - - - - - -
		//
		//      addressing    assembler    opc  bytes  cyles
		//      --------------------------------------------
		//      relative      BNE oper      D0    2     2**
		// ** The offset is a signed byte, so it can jump a maximum of 127 bytes forward, or 128 bytes backward. **
		case 0xD0: // BNE

			// If P[1] = 1 (Zero Flag)
			if P[1] == 1 {

				if debug {
					fmt.Printf("\n\tOpcode %02X%02X [2 bytes]\tBNE  Branch on Result not Zero.\t| Zero Flag(P1) = %d | PC += 2\n", Opcode, Memory[PC+1], P[1])
				}
				PC += 2

			} else {
				offset := DecodeTwoComplement(Memory[PC+1])
				//fmt.Printf("\tOffset(%02X): %d \n", Memory[PC+1], offset)

				// Increment PF + the offset
				PC += 2 + uint16(offset)
				if debug {
					fmt.Printf("\n\tOpcode %02X%02X [2 bytes]\tBNE  Branch on Result not Zero.\tZero Flag(P1) = %d | PC = Jump to Memory[%02X] (%02X)\n", Opcode, Memory[PC+1], Memory[SP], PC, Memory[PC])
					Beam_index += 1
				}
			}
			Beam_index += 2 // ************** PODE VARIAR!!! IMPLEMENTAR **************


		// BCC  Branch on Carry Clear
		//
		//      branch on C = 0                  N Z C I D V
		//                                       - - - - - -
		//
		//      addressing    assembler    opc  bytes  cyles
		//      --------------------------------------------
		//      relative      BCC oper      90    2     2**
		case 0x90: // BCC
			// If carry is clear
			if P[0] == 0 {
				if debug {
					fmt.Printf("\n\tOpcode %02X%02X [2 bytes]\tBCC  Branch on Carry Clear (relative).\tCarry EQUAL 0, JUMP TO %04X\n", Opcode, Memory[PC+1], PC+2+uint16(Memory[PC+1]))
				}

				// PC+=2 to step to next instruction + the number of bytes to jump on carry clear
				PC+=2+uint16(DecodeTwoComplement(Memory[PC+1]))

			// If carry is set
			} else {
				if debug {
					fmt.Printf("\n\tOpcode %02X%02X [2 bytes]\tBCC  Branch on Carry Clear (relative).\tCarry NOT EQUAL 0, PC+2 \n", Opcode, Memory[PC+1])
				}
				PC += 2

			}

			Beam_index += 2


		// BCS  Branch on Carry Set
		//
		//      branch on C = 1                  N Z C I D V
		//                                       - - - - - -
		//
		//      addressing    assembler    opc  bytes  cyles
		//      --------------------------------------------
		//      relative      BCS oper      B0    2     2**
		case 0xB0: // BCC
			// If carry is clear
			if P[0] == 1 {
				if debug {
					fmt.Printf("\n\tOpcode %02X%02X [2 bytes]\tBCS  Branch on Carry Set (relative).\tCarry EQUAL 1, JUMP TO %04X\n", Opcode, Memory[PC+1], PC+2+uint16(Memory[PC+1]))
				}

				// PC+=2 to step to next instruction + the number of bytes to jump on carry clear
				PC+=2+uint16(DecodeTwoComplement(Memory[PC+1]))

			// If carry is set
			} else {
				if debug {
					fmt.Printf("\n\tOpcode %02X%02X [2 bytes]\tBCS  Branch on Carry Set (relative).\tCarry NOT EQUAL 1, PC+2 \n", Opcode, Memory[PC+1])
				}
				PC += 2
			}

			Beam_index += 2


		//-------------------------------------------------- LDX --------------------------------------------------//


		// LDX  Load Index X with Memory
		//
		//      M -> X                           N Z C I D V
		//                                       + + - - - -
		//
		//      addressing    assembler    opc  bytes  cyles
		//      --------------------------------------------
		//      immidiate     LDX #oper     A2    2     2
		case 0xA2: // LDX immidiate
			X = Memory[PC+1]
			if debug {
				fmt.Printf("\n\tOpcode %02X%02X [2 bytes]\tLDX  Load Index X with Memory (immidiate).\tX = Memory[%02X] (%d)\n", Opcode, Memory[PC+1], PC+1, X)
			}
			PC += 2

			flags_Z(X)
			flags_N(X)
			Beam_index += 2


		//-------------------------------------------------- STX --------------------------------------------------//


		// STX  Store Index X in Memory (zeropage)
		//
		//      X -> M                           N Z C I D V
		//                                       - - - - - -
		//
		//      addressing    assembler    opc  bytes  cyles
		//      --------------------------------------------
		//      zeropage      STX oper      86    2     3
		case 0x86: // STX

			Memory[Memory[PC+1]] = X
			if debug {
				fmt.Printf("\n\tOpcode %02X%02X [2 bytes]\tSTX  Store Index X in Memory (zeropage).\tMemory[%02X] = X (%d)\n", Opcode, Memory[PC+1], Memory[PC+1], X)
			}

			PC += 2
			Beam_index += 3

		//-------------------------------------------------- LDA --------------------------------------------------//


		// LDA  Load Accumulator with Memory (immidiate)
		//
		//      M -> A                           N Z C I D V
		//                                       + + - - - -
		//
		//      addressing    assembler    opc  bytes  cyles
		//      --------------------------------------------
		//      immidiate     LDA #oper     A9    2     2
		case 0xA9: // LDA (immidiate)

			A = Memory[PC+1]
			if debug {
				fmt.Printf("\n\tOpcode %02X%02X [2 bytes]\tLDA  Load Accumulator with Memory (immidiate).\tA = Memory[%02X] (%d)\n", Opcode, Memory[PC+1], PC+1, A)
			}

			flags_Z(A)
			flags_N(A)
			PC += 2
			Beam_index += 2


		// LDA  Load Accumulator with Memory (zeropage)
		//
		//      M -> A                           N Z C I D V
		//                                       + + - - - -
		//
		//      addressing    assembler    opc  bytes  cyles
		//      --------------------------------------------
		//      zeropage      LDA oper      A5    2     3
		case 0xA5: // LDA (zeropage)

			A = Memory[Memory[PC+1]]
			if debug {
				fmt.Printf("\n\tOpcode %02X%02X [2 bytes]\tLDA  Load Accumulator with Memory (zeropage).\tA = Memory[%02X] (%d)\n", Opcode, Memory[PC+1], Memory[PC+1], A)
			}

			flags_Z(A)
			flags_N(A)
			PC += 2
			Beam_index += 3


		// LDA  Load Accumulator with Memory (absolute,Y)
		//
		//      M -> A                           N Z C I D V
		//                                       + + - - - -
		//
		//      addressing    assembler    opc  bytes  cyles
		//      --------------------------------------------
		//      absolute,Y    LDA oper,Y    B9    3     4*
		case 0xB9: // LDA (absolute,Y)
			tmp := uint16(Memory[PC+2])<<8 | uint16(Memory[PC+1])
			A = Memory[tmp + uint16(Y)]
			// fmt.Printf("\n\n\n%08b", A)
			// fmt.Printf("\n\n\n%d", A)
			if debug {
				fmt.Printf("\n\tOpcode %02X%02X%02X [3 bytes]\tLDA  Load Accumulator with Memory (absolute,Y).\tA = Memory[%04X + Y(%d)]  (%d)\n", Opcode, Memory[PC+2], Memory[PC+1], tmp, Y, A)
			}
			PC += 3
			flags_Z(A)
			flags_N(A)
			Beam_index += 4

		//-------------------------------------------------- STA --------------------------------------------------//


		// STA  Store Accumulator in Memory (zeropage,X)
		//
		//      A -> M                           N Z C I D V
		//                                       - - - - - -
		//
		//      addressing    assembler    opc  bytes  cyles
		//      --------------------------------------------
		//      zeropage,X    STA oper,X    95    2     4
		case 0x95: // STA (zeropage, X)
			Memory[Memory[PC+1]] = A
			if debug {
				fmt.Printf("\n\tOpcode %02X%02X [2 bytes]\tSTA  Store Accumulator in Memory (zeropage, X).\tMemory[%02X] = A (%d)\n", Opcode,Memory[PC+1], Memory[PC+1], Memory[Memory[PC+1]] )
			}
			PC += 2
			Beam_index += 4


		// STA  Store Accumulator in Memory (zeropage)
		//
		//      A -> M                           N Z C I D V
		//                                       - - - - - -
		//
		//      addressing    assembler    opc  bytes  cyles
		//      --------------------------------------------
		//      zeropage      STA oper      85    2     3
		case 0x85: // STA (zeropage)

			Memory[Memory[PC+1]] = A
			if debug {
				fmt.Printf("\n\tOpcode %02X%02X [2 bytes]\tSTA  Store Accumulator in Memory (zeropage).\tMemory[%02X] = A (%d)\n", Opcode, Memory[PC+1], Memory[PC+1], Memory[Memory[PC+1]] )
			}

			Beam_index += 3

			// Wait for a new line and authorize graphics to draw the line
			// Wait for Horizontal Blank to draw the new line
			if Memory[PC+1] == WSYNC {
				DrawLine = true

				if Memory[GRP0] != 0 {
					DrawP0 = true
					// Pause = true
				}

				if Memory[GRP1] != 0 {
					DrawP1 = true
				}

			// if Memory[PC+1] == GRP0 {
			// 	if Memory[GRP0] != 0 {
			// 		// for i:=0 ; i <=7 ; i++{
			// 		// 	bit := CPU.Memory[GRP0] >> 7-byte(i) & 0x01
			// 		// fmt.Printf("\n\n\n\n\n\n\n\n\n\n\n\n%08b\n", Memory[GRP0])
			// 		// fmt.Printf("\nmem: %08b\n",Memory[0x1B])
			// 		// os.Exit(2)
			// 		DrawP0 = true
			// 		// DrawP0VertPosition = Beam_index
			// 		// fmt.Printf("\nDraw P0")
			// 	}
			//
			// }

			// if Memory[PC+1] == GRP1 {
			// 	if Memory[GRP1] != 0 {
			// 		// for i:=0 ; i <=7 ; i++{
			// 		// 	bit := CPU.Memory[GRP0] >> 7-byte(i) & 0x01
			// 		// fmt.Printf("\n\n\n\n\n\n\n\n\n\n\n\n%08b\n", Memory[GRP0])
			// 		// fmt.Printf("\nmem: %08b\n",Memory[0x1B])
			// 		// os.Exit(2)
			// 		DrawP1 = true
			// 		// DrawP1VertPosition = Beam_index
			// 		// fmt.Printf("\nDraw P0")
			// 	}

			}

			PC += 2


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
			if debug {
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
			if debug {
				fmt.Printf("\n\tOpcode %02X%02X [2 bytes]\tLDY  Load Index y with Memory (zeropage).\tY = Memory[%02X](%d)\t(%d)\n", Opcode, Memory[PC+1], PC+1, Memory[PC+1], Y)
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
			if debug {
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
		case 0x4C: // JMP (absolute)
			tmp := uint16(Memory[PC+2])<<8 | uint16(Memory[PC+1])
			//fmt.Printf("\n%04X\n",tmp)
			if debug {
				fmt.Printf("\n\tOpcode %02X%02X%02X [3 bytes]\tJMP  Jump to New Location (absolute).\tPC = Memory[%d](%d)\n", Opcode, Memory[PC+2], Memory[PC+1], tmp, Memory[tmp])
			}
			PC = tmp
			Beam_index += 3

		// ISB (INC FOLLOWED BY SBC - IMPLEMENT IT!!!!!!)
		// FF (Filled ROM)
		case 0xFF:
			if debug {
				fmt.Printf("\n\tOpcode %02X [1 byte]\tFilled ROM.\tPC incremented.\n", Opcode)
			}
			PC +=1


		//-------------------------------------------------- CPY --------------------------------------------------//


		// CPY  Compare Memory and Index Y
		//
		//      Y - M                            N Z C I D V
		//                                       + + + - - -
		//
		//      addressing    assembler    opc  bytes  cyles
		//      --------------------------------------------
		//      immidiate     CPY #oper     C0    2     2
		//
		//
		case 0xC0: //CPY

			tmp := Y - Memory[PC+1]

			if debug {
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
		case 0xC4: //CPY

			tmp := Y - Memory[Memory[PC+1]]

			if debug {
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

		// SBC  Subtract Memory from Accumulator with Borrow
		//
	     // A - M - C -> A                   N Z C I D V
	     //                                  + + + - - +
		//
	     // addressing    assembler    opc  bytes  cyles
	     // --------------------------------------------
	     // zeropage      SBC oper      E5    2     3
		case 0xE5: // SBC
			// Memory[Memory[PC+1]] = X
			if debug {
				fmt.Printf("\n\tNEED TO IMPLEMENT OVERFLOW!!! - Opcode %02X%02X [2 bytes]\tSBC  Subtract Memory from Accumulator with Borrow (zeropage).\tA = A(%d) - Memory[Memory[%02X]](%d) - (Carry(%d)-1)= %d\n", Opcode, Memory[PC+1], A, PC+1, Memory[Memory[PC+1]], P[0] , A - Memory[Memory[PC+1]] - (1-P[0]))
			}

			// A-M-(1-Carry)
			A = A - Memory[Memory[PC+1]] - (1-P[0])

			flags_Z(A)
			flags_N(A)
			// flags_C will be cleared if overflow in bit 7
			// FALTA TRATAR OVERFLOW

			PC += 2
			Beam_index += 3

//** IMPLEMENT OVERFLOW **//
		// SBC  Subtract Memory from Accumulator with Borrow (immidiate)
		//
		//      A - M - C -> A                   N Z C I D V
		//                                       + + + - - +
		//
		//      addressing    assembler    opc  bytes  cyles
		//      --------------------------------------------
		//      immidiate     SBC #oper     E9    2     2
		case 0xE9: // SBC
			// Memory[Memory[PC+1]] = X
			if debug {
				fmt.Printf("\n\tNEED TO IMPLEMENT OVERFLOW!!! - Opcode %02X%02X [2 bytes]\tSBC  Subtract Memory from Accumulator with Borrow (immidiate).\tA = A(%d) - Memory[%02X](%d) - (Carry(%d)-1)= %d\n", Opcode, Memory[PC+1], A, PC+1, Memory[PC+1], P[0] , A - Memory[PC+1] - (1-P[0]))
			}

			flags_C(A,Memory[PC+1])
			Flags_V_SBC(A,Memory[PC+1])

			// A-M-(1-Carry)
			A = A - Memory[PC+1] - (1-P[0])

			flags_Z(A)
			flags_N(A)




			// Clear Carry if overflow in bit 7
			if P[6] == 1 {
				P[0] = 0
			}

			PC += 2
			Beam_index += 2
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
		case 0xC5: // CMP

			tmp := A - Memory[Memory[PC+1]]

			if debug {
				if tmp == 0 {
					fmt.Printf("\n\tOpcode %02X%02X [2 bytes]\tCMP  Compare Memory with Accumulator (zeropage).\tA(%d) - Memory[%02X](%d) = (%d) EQUAL\n", Opcode, Memory[PC+1], A, PC+1, Memory[PC+1], tmp)
				} else {
					fmt.Printf("\n\tOpcode %02X%02X [2 bytes]\tCMP  Compare Memory with Accumulator (zeropage).\tA(%d) - Memory[%02X](%d) = (%d) NOT EQUAL\n", Opcode, Memory[PC+1], A, PC+1, Memory[PC+1], tmp)
				}
			}

			flags_Z(tmp)
			flags_N(tmp)
			flags_C(A,Memory[Memory[PC+1]])

			PC += 2
			Beam_index += 3


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

			if debug {
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

			if debug {
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

			if debug {
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

			if debug {
				fmt.Printf("\n\tOpcode %02X [1 byte]\tASL  Shift Left One Bit (Memory or Accumulator) (accumulator).\tA = A(%d) Shift Left 1 bit\t(%d)\n", Opcode, A, A << 1 )
			}

			flags_C(A, A << 1)

			A = A << 1

			flags_N(A)
			flags_Z(A)

			PC += 1
			Beam_index += 2





		default:
			fmt.Printf("\n\tOPCODE %X NOT IMPLEMENTED!\n\n", Opcode)
			os.Exit(2)

	}

	Cycle ++

}
