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

	// CPU Variables
	Opcode		uint16		// CPU Operation Code
	Cycle		uint16		// CPU Cycle

	// Timers
	Clock			*time.Ticker	// CPU Clock // CPU: MOS Technology 6507 @ 1.19 MHz;
	ScreenRefresh		*time.Ticker	// Screen Refresh
	Second			= time.Tick(time.Second)			// 1 second to track FPS and draws

	debug 		bool = true

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
	ScreenRefresh	= time.NewTicker(time.Second / 60)	// 60Hz Clock for screen refresh rate

}

// Reset Vector // 0xFFFC | 0xFFFD (Little Endian)
func Reset() {
	// Read the Opcode from PC+1 and PC bytes (Little Endian)
	PC = uint16(Memory[0xFFFD])<<8 | uint16(Memory[0xFFFC])
	fmt.Printf("\n%04X\n",PC)
}

func Show() {
	fmt.Printf("\nCycle: %d\tOpcode: %02X\tPC: 0x%X(%d)\tA: %d\tX: 0x%02X\tY: %02X\tP: %d\tSP: %02X\tStack: [%d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d]", Cycle, Opcode, PC, PC, A, X, Y, P, SP, Memory[0xFF], Memory[0xFE], Memory[0xFD], Memory[0xFC], Memory[0xFB], Memory[0xFA], Memory[0xF9], Memory[0xF8], Memory[0xF7], Memory[0xF6], Memory[0xF5], Memory[0xF4], Memory[0xF3], Memory[0xF2], Memory[0xF1], Memory[0xF0] )
}

// CPU Interpreter
func Interpreter() {

	// Read the Next Instruction to be executed
	Opcode = uint16(Memory[PC])

	// Print Cycle and Debug Information
	Show()

	// Map Opcode Family
	switch Opcode {

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
			PC	+= 1
			if debug {
				fmt.Printf("\n\tOpcode %02X\tSEI  Set Interrupt Disable Status\tP[2]=1\n\n", Opcode)
			}

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
			PC	+= 1
			if debug {
				fmt.Printf("\n\tOpcode %02X\tCLD  Clear Decimal Mode\tP[3]=0\n\n", Opcode)
			}

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
			PC	+= 2
			if debug {
				fmt.Printf("\n\tOpcode %02X\tLDX  Load Index X with Memory (immidiate)\tX = Memory[%02X] (%d)\n\n", Opcode, PC+1, X)
			}

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
			PC	+= 1
			if debug {
				fmt.Printf("\n\tOpcode %02X\tTXA  Transfer Index X to Accumulator\tA = X (%d)\n\n", Opcode, X)
			}

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
			PC	+= 1
			if debug {
				fmt.Printf("\n\tOpcode %02X\tTAY  Transfer Accumulator to Index Y\tA = X (%d)\n\n", Opcode, X)
			}

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
			PC	+= 1
			if debug {
				fmt.Printf("\n\tOpcode %02X\tDEX  Decrement Index X by One\tX-- (%d)\n\n", Opcode, X)
			}

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
			PC	+= 1
			if debug {
				fmt.Printf("\n\tOpcode %02X\tTXS  Transfer Index X to Stack Register\tSP = X (%d)\n\n", Opcode, SP)
			}

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
			PC	+= 1
			if debug {
				fmt.Printf("\n\tOpcode %02X\tPHA  Push Accumulator on Stack\tMemory[%02X] = A (%d) | SP--\n\n", Opcode, SP ,Memory[SP])
			}
			SP--

		// BNE  Branch on Result not Zero (Accumulator)
		//
		//      branch on Z = 0                  N Z C I D V
		//                                       - - - - - -
		//
		//      addressing    assembler    opc  bytes  cyles
		//      --------------------------------------------
		//      relative      BNE oper      D0    2     2**
		case 0xD0: // BNE

			// Test A and set or no the Zero Flag
			if A == 0 {
				P[1] = 1
			} else {
				P[1] = 0
			}

			// 1    Z     Zero          (0=Nonzero, 1=Zero)
			if P[1] == 1 {
				PC	+= 2
				if debug {
					fmt.Printf("\n\tOpcode %02X\tBNE  Branch on Result not Zero (Accumulator)\tZero Flag(P1) = %d | PC += 2\n\n", Opcode ,Memory[SP])
				}

			} else {
				PC = uint16(Memory[PC+1])
				if debug {
					fmt.Printf("\n\tOpcode %02X\tBNE  Branch on Result not Zero (Accumulator)\tZero Flag(P1) = %d | PC = Jump to Memory[%02X] (%02X)\n\n", Opcode ,Memory[SP], PC+1, Memory[PC+1])
				}
			}

			// STX  Store Index X in Memory (zeropage)
			//
			//      X -> M                           N Z C I D V
			//                                       - - - - - -
			//
			//      addressing    assembler    opc  bytes  cyles
			//      --------------------------------------------
			//      zeropage      STX oper      86    2     3
			case 0x86: // STX
				Memory[PC+1]	= X

				PC	+= 2
				if debug {
					fmt.Printf("\n\tOpcode %02X\tSTX  Store Index X in Memory (zeropage)\tMemory[%02X] = X (%d)\n\n", Opcode, PC+1, X)
				}

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

				PC	+= 2
				if debug {
					fmt.Printf("\n\tOpcode %02X\tLDA  Load Accumulator with Memory (immidiate)\tA = Memory[%02X] (%d)\n\n", Opcode, PC+1, A)
				}


		default:
			fmt.Printf("\n\tOPCODE %X NOT IMPLEMENTED!\n\n", Opcode)
			os.Exit(2)

	}

	Cycle ++

}
