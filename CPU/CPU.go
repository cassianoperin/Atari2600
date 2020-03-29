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
	// fmt.Printf("Cycle: %d\tOpcode: %04X(%04X)\tPC: %d(0x%X)\tSP: %d\tStack: %d\tV: %d\tI: %d\tDT: %d\tST: %d\tKey: %d\n", Cycle, Opcode, Opcode & 0xF000, PC, PC,  SP, Stack, V, I, DelayTimer, SoundTimer, Key)
	fmt.Printf("\nCycle: %d\tOpcode: %02X\tPC: %d(0x%X)\tA: %d\tX: %d\tY: %d\tP: %d", Cycle, Opcode, PC, PC, A, X, Y, P)
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
			Memory[PC+1]	=  X
			PC	+= 2
			if debug {
				fmt.Printf("\n\tOpcode %02X\tLDX  Load Index X with Memory (immidiate)\tMemory[%02X]	=  X (%d)\n\n", Opcode, PC+1, X)
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

		default:
			fmt.Printf("\n\tOPCODE %X NOT IMPLEMENTED!\n\n", Opcode)
			os.Exit(2)

	}

	Cycle ++

}
