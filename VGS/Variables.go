package VGS

import (
	"time"

	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

// Fullscreen / Video Modes
type Setting struct {
	Mode    *pixelgl.VideoMode
	Monitor *pixelgl.Monitor
}

var (
	// ------------------------------ Graphics ------------------------------ //
	// Screen Size
	sizeX float64 = 160.0 // 68 color clocks (Horizontal Blank) + 160 color clocks (pixels)
	sizeY float64 = 192.0 // 3 Vertical Sync, 37 Vertical Blank, 192 Visible Area and 30 Overscan
	// Window Resolution
	screenWidth  float64 = 1024
	screenHeight float64 = 768
	// Pixel size
	width  float64
	height float64

	// ------------------------------ Counters ------------------------------ //
	counter_F_Cycle uint16 // Frame Cycles
	counter_DPS     uint16 // Draws per second
	counter_VSYNC   byte   // Count the tree VSYNCs each frame
	counter_FPS     uint16 // Frames per second
	counter_Frame   uint16 // Frames Counter

	// ------------------------------- Window ------------------------------- //
	// win					*pixelgl.Window
	imd                      = imdraw.New(nil)
	cfg                      = pixelgl.WindowConfig{}
	windowTitle       string = "Atari 2600"
	settings          []Setting
	activeSetting     *Setting
	isFullScreen          = false // Fullscrenn flag
	resolutionCounter int = 0     // Index of the available video resolution supported
	// Monitor Size (to center Window)
	monitorWidth  float64
	monitorHeight float64

	// ----------------------------- RIOT Timer ----------------------------- //
	riot_timer         byte
	riot_timer_counter uint16
	riot_timer_mult    uint16
	old_timer          byte

	// ------------------------------- Timers ------------------------------- //
	second_timer = time.Tick(time.Second)

	// ------------------------------- Beamer ------------------------------- //
	beamIndex byte // Beam index to control where to draw objects using cpu cycles

	// -------------------------------- TIA --------------------------------- //
	line       int   // Line draw control
	line_max   int   // Line draw control
	TIA_Update int16 // Tells Graphics that a TIA register was changed (values >= 0 (addresses) will be detected)

	// // ---------------------- Debug Timing Measurement ---------------------- //
	debugTiming             bool
	debugTiming_Limit       float64
	debugTiming_StartCycle  time.Time
	debugTiming_StartTIA    time.Time
	debugTiming_StartTIA_BG time.Time

	// // ----------------------------- Messaging ------------------------------ //
	// atlas       = text.NewAtlas(basicfont.Face7x13, text.ASCII) // Font
	// textMessage *text.Text                                      // On screen Message content
	// cpuMessage  *text.Text                                      // In screen CPU components debug
	// // textFPS		*text.Text	// On screen FPS counter
	// // textFPSstr	string		// String with the FPS counter
	// // drawCounter	= 0		// imd.Draw per second counter
	// // updateCounter	= 0		// win.Updates per second counter
	// ShowMessage    bool
	// TextMessageStr string

	// -------------------- Players Vertical Positioning -------------------- //
	XPositionP0     byte
	XFinePositionP0 int8
	P0_bit          byte = 0
	P1_bit          byte = 0

	XPositionP1     byte
	XFinePositionP1 int8

	// ------------------------ Collision Detection ------------------------- //
	collision_PF [161]byte // 161 because pixel_position starts with 1 and goes until 160 (Ignore the position 0 in this slice)
	collision_P0 [161]byte // 161 because pixel_position starts with 1 and goes until 160 (Ignore the position 0 in this slice)
	collision_P1 [161]byte // 161 because pixel_position starts with 1 and goes until 160 (Ignore the position 0 in this slice)

	// ----------------------------- Playfield ------------------------------ //
	// PF(4,5,6,7) | PF1 (7,6,5,4,3,2,1,0) | PF2 (0,1,2,3,4,5,6,7)
	pf0_bit        byte = 4 // PF0 bit index
	pf1_bit        byte = 7 // PF1 bit index
	pf2_bit        byte = 0 // PF2 bit index
	pf0_mirror_bit byte = 7 // PF2 bit index
	pf1_mirror_bit byte = 0 // PF2 bit index
	pf2_mirror_bit byte = 7 // PF2 bit index

	// Debug
	debugGraphics bool = true  // Graphics Debug mode
	debugRIOT     bool = false // RIOT Debug mode
)

const (
	//-------------------------------------------------- Memory locations -------------------------------------------------//

	//0000-002C - TIA (write)
	//0030-003D - TIA (read) - (sometimes mirrored at 0030-003D)
	//0080-00FF - RIOT (RAM) (128 bytes) -- Stack uses the last addresses
	//0280-0297 - RIOT (I/O, Timer)
	//F000-FFFF - Cartridge (ROM)

	//------------------- 0000-002C - TIA (write)
	VSYNC  byte = 0x00 //0000 00x0   Vertical Sync Set-Clear
	VBLANK byte = 0x01 //xx00 00x0   Vertical Blank Set-Clear
	WSYNC  byte = 0x02 //---- ----   Wait for Horizontal Blank
	RSYNC  byte = 0x03 //---- ----   Reset Horizontal Sync Counter
	NUSIZ0 byte = 0x04 //00xx 0xxx   Number-Size player/missle 0
	NUSIZ1 byte = 0x05 //00xx 0xxx   Number-Size player/missle 1
	COLUP0 byte = 0x06 //xxxx xxx0   Color-Luminance Player 0
	COLUP1 byte = 0x07 //xxxx xxx0   Color-Luminance Player 1
	COLUPF byte = 0x08 //xxxx xxx0   Color-Luminance Playfield
	COLUBK byte = 0x09 //xxxx xxx0   Color-Luminance Background
	// CTRLPLF (8 bits register)
	// D0 = 0 Repeat the PF, D0 = 1 = Reflect the PF
	// D1 = Score == Color of the score will be the same as player
	// D2 = Priority == Player behind the playfield
	// D4-5 = Ball Size (1, 2, 4, 8)
	CTRLPF byte = 0x0A //00xx 0xxx   Control Playfield, Ball, Collisions
	REFP0  byte = 0x0B //0000 x000   Reflection Player 0
	REFP1  byte = 0x0C //0000 x000   Reflection Player 1
	PF0    byte = 0x0D //xxxx 0000   Playfield Register Byte 0
	PF1    byte = 0x0E //xxxx 0000   Playfield Register Byte 1
	PF2    byte = 0x0F //xxxx 0000   Playfield Register Byte 2
	RESP0  byte = 0x10 //---- ----   Reset Player 0
	RESP1  byte = 0x11 //---- ----   Reset Player 1
	GRP0   byte = 0x1B //xxxx xxxx   Graphics Register Player 0
	GRP1   byte = 0x1C //xxxx xxxx   Graphics Register Player 1
	HMP0   byte = 0x20 // xxxx 0000   Horizontal Motion Player 0
	HMP1   byte = 0x21 // xxxx 0000   Horizontal Motion Player 1
	HMM0   byte = 0x22 // xxxx 0000   Horizontal Motion Missle 0
	HMM1   byte = 0x23 // xxxx 0000   Horizontal Motion Missle 1
	HMBL   byte = 0x24 // xxxx 0000   Horizontal Motion Ball
	HMOVE  byte = 0x2A // ---- ----   Apply Horizontal Motion
	HMCLR  byte = 0x2B // ---- ----   Clear Horizontal Move Registers
	CXCLR  byte = 0x2C // ---- ----   Clear Collision Latches

	// ;-------------------------------------------------------------------------------
	//
	// 			SEG.U TIA_REGISTERS_READ
	// 			ORG TIA_BASE_READ_ADDRESS
	//
	// ;															bit 7   bit 6
	CXM0P  byte = 0x00 // xx00 0000     Read Collision  M0-P1   M0-P0
	CXM1P  byte = 0x01 // xx00 0000                     M1-P0   M1-P1
	CXP0FB byte = 0x02 // xx00 0000                     P0-PF   P0-BL
	CXP1FB byte = 0x03 // xx00 0000                     P1-PF   P1-BL
	CXM0FB byte = 0x04 // xx00 0000                     M0-PF   M0-BL
	CXM1FB byte = 0x05 // xx00 0000                     M1-PF   M1-BL
	CXBLPF byte = 0x06 // x000 0000                     BL-PF   -----
	CXPPMM byte = 0x07 // xx00 0000                     P0-P1   M0-M1
	INPT0  byte = 0x08 // x000 0000       Read Pot Port 0
	INPT1  byte = 0x09 // x000 0000       Read Pot Port 1
	INPT2  byte = 0x0A // x000 0000       Read Pot Port 2
	INPT3  byte = 0x0B // x000 0000       Read Pot Port 3
	// P0 action Button  INPT4.7   Button (0=Pressed, 1=Not pressed)
	INPT4 byte = 0x0C // x000 0000       Read Input (Trigger) 0
	// P1 action Button  INPT5.7   Button (0=Pressed, 1=Not pressed)
	INPT5 byte = 0x0D // x000 0000       Read Input (Trigger) 1

	//------------------- 0280-0297 - RIOT (I/O, Timer)
	SWCHA uint16 = 0x280 // Port A data register for joysticks: Bits 4-7 for player 1.  Bits 0-3 for player 2.
	//SWACNT      ds 1    ; $281      Port A data direction register (DDR)
	SWCHB uint16 = 0x282 // Port B data (console switches)
	// Data			BitSwitchBit		Meaning
	// D7			P1 difficulty		0 = amateur (B), 1 = pro (A)
	// D6			P0 difficulty		0 = amateur (B), 1 = pro (A)
	// D5/D4		(not used)
	// D3			color - B/W		0 = B/W, 1 = color
	// D2			(not used)
	// D1			game select		0 = switch pressed
	// D0			game reset		0 = switch pressed
	//SWBCNT      ds 1    ; $283      Port B DDR
	INTIM  uint16 = 0x284 // Timer output
	TIMINT uint16 = 0x285 // Timer interrupt flag, which is 0 if the timer hasn't passed 0 yet, and is set to 128 (bit 7 on)
	// once the timer has passed 0. (Actually, the TIMINT register contains another flag in bit 6,
	// but it isn't used by the Atari 2600, so TIMINT will either be 0 or 128.
	// That way, you can set the timer to a value greater than 127, and you won't have to worry about INTIM starting out with a "negative" value.
	TIM1T  uint16 = 0x294 // Set 1 clock interval
	TIM8T  uint16 = 0x295 // Set 8 clock interval
	TIM64T uint16 = 0x296 // Set 64 clock interval
	T1024T uint16 = 0x297 // Set 1024 clock interval

)
