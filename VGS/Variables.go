package VGS

import (
	"time"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
)

// Fullscreen / Video Modes
type Setting struct {
	Mode	*pixelgl.VideoMode
	Monitor	*pixelgl.Monitor
}

var (

	// ------------------------ Hardware Components ------------------------- //
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

	// --------------------------- CPU Variables ---------------------------- //
	opcode			byte		// CPU Operation Code

	// ------------------------------ Counters ------------------------------ //
	counter_Frame		uint16		// Frames Counter
	counter_F_Cycle		uint16		// Frame Cycles
	opc_cycle_count		byte		// Opcode cycle counter
	opc_cycle_extra		byte		// Opcode extra cycle
	counter_IPS			uint16		// Instructions per second
	counter_FPS			uint16		// Frames per second
	counter_DPS			uint16		// Draws per second

	// -------------------------- Memory Variables -------------------------- //
	memAddr			uint16		// Receive the memory address needed by the opcode
	memMode			string		// Receive the addressing mode used in the debug
	memValue		int8		// Receive the memory value needed by branches. Calculated in the first opc cycle to check for extra cycles, used in the last to perform the operation

	// ------------------------------- Timers ------------------------------- //
	clock_timer			*time.Ticker	// CPU Clock // CPU: MOS Technology 6507 @ 1.19 MHz;
	screenRefresh_timer	*time.Ticker	// Screen Refresh
	second_timer		= time.Tick(time.Second)			// 1 second to track FPS and draws
	messagesClock_timer	*time.Ticker		// Clock used to display messages on screen

	// ------------------------------- Beamer ------------------------------- //
	beamIndex	byte 		// Beam index to control where to draw objects using cpu cycles
	old_beamIndex	byte 	// Used to draw the beam updates every cycle on the CRT

	// -------------------------------- TIA --------------------------------- //
	line		int			// Line draw control
	line_max	int			// Line draw control
	TIA_Update	int8		// Tells Graphics that a TIA register was changed (values >= 0 (addresses) will be detected)
	visibleArea	bool		// Just draw in visible area
	VSYNC_passed		bool = false	// Workaround to avoid  WSYNC before VSYNC

	// ---------------------- Debug Timing Measurement ---------------------- //
	debugTiming 			bool
	debugTiming_Limit		float64
	debugTiming_StartCycle	time.Time
	debugTiming_StartTIA	time.Time
	debugTiming_StartTIA_BG	time.Time

	// ----------------------------- Messaging ------------------------------ //
	atlas = text.NewAtlas(basicfont.Face7x13, text.ASCII)	// Font
	textMessage		*text.Text	// On screen Message content
	cpuMessage  	*text.Text	// In screen CPU components debug
	// textFPS		*text.Text	// On screen FPS counter
	// textFPSstr	string		// String with the FPS counter
	// drawCounter	= 0		// imd.Draw per second counter
	// updateCounter	= 0		// win.Updates per second counter
	ShowMessage		bool
	TextMessageStr	string

	// -------------------- Players Vertical Positioning -------------------- //
	XPositionP0			byte
	XFinePositionP0			int8
	XPositionP1			byte
	XFinePositionP1			int8

	// ------------------------ Collision Detection ------------------------- //
	CD_debug			bool	= true
	CD_P0_P1			[160]byte
	CD_P0_P1_collision_detected			bool	= false	// Set when collision is detected
	CD_P0_PF			[160]byte
	CD_P0_PF_collision_detected			bool	= false	// Set when collision is detected

	// ----------------------------- Playfield ------------------------------ //
	// PF0(4,5,6,7) | PF1 (7,6,5,4,3,2,1,0) | PF2 (0,1,2,3,4,5,6,7)
	playfield			[40]byte			//Improve to binary
	pixelSize			float64 = 4.0		// 80 lines (half screen) / 20 PF0, PF1 and PF2 bits

	// ------------------------------ Graphics ------------------------------ //
	debugGraphics	bool	// Graphics Debug mode
	// Screen Size
	sizeX			float64	= 160.0 	// 68 color clocks (Horizontal Blank) + 160 color clocks (pixels)
	sizeY			float64	= 192.0		// 3 Vertical Sync, 37 Vertical Blank, 192 Visible Area and 30 Overscan
	sizeYused		float64	= 1.0	// Percentage of the Screen Heigh used by the emulator // 1.0 = 100%, 0.0 = 0%
	sizeXused		float64	= 1.0	// Percentage of the Screen Width used by the emulator // 1.0 = 100%, 0.0 = 0%
	// Window Resolution
	screenWidth		float64 = 1024
	screenHeight	float64 = 768
	// Pixel size
	width			float64
	height			float64

	// ------------------------------- Window ------------------------------- //
	win					*pixelgl.Window
	imd					= imdraw.New(nil)
	cfg					= pixelgl.WindowConfig{}
	windowTitle			string = "Atari 2600"
	settings			[]Setting
	activeSetting		*Setting
	isFullScreen		= false		// Fullscrenn flag
	resolutionCounter	int = 0		// Index of the available video resolution supported
	// Monitor Size (to center Window)
	monitorWidth	float64
	monitorHeight	float64

	// --------------------------- Debug Interface -------------------------- //
	debug_opc_text	string
	dbg_opc_bytes		uint16
	dbg_show_message	string	// Debug opcode detail messages
	// Opcode Message Block
	dbg_opc_messages []string
	// Running opcode flag - Used to advance entire opcode and not just a cycle
	dbg_running_opc	bool
	// Breakpoint
	dbg_break		bool
	dbg_break_cycle	uint16

	// Pause
	Pause		bool = true

	// Debug
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