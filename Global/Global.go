// Global Variables used to to avoid circular dependencies
package Global

import (
	"github.com/faiface/pixel/pixelgl"
	// "github.com/faiface/pixel/text"
)

// Fullscreen / Video Modes
type Setting struct {
	Mode	*pixelgl.VideoMode
	Monitor	*pixelgl.Monitor
}

var (
	// ------------------------ Global Variables ------------------------ //
	// Game_signature		string	= ""	// Game Signature (identify games that needs legacy opcodes)

	// ----------------------- Graphics Variables ----------------------- //
	Win			*pixelgl.Window
	WindowTitle		string = "Atari 2600"
	// Color_theme		= 2
	// // Fullscreen / Video Modes
	// Texts			[]*text.Text
	// StaticText		*text.Text
	Settings		[]Setting
	ActiveSetting		*Setting
	IsFullScreen		= false		// Fullscrenn flag
	ResolutionCounter	int = 0		// Index of the available video resolution supported
	// FPS
	// ShowFPS			bool		// Show or hide FPS counter flag
	// On screen messages
	ShowMessage		bool
	TextMessageStr		string

	// Screen Size
	SizeX			float64	= 160.0 	// 68 color clocks (Horizontal Blank) + 160 color clocks (pixels)
	SizeY			float64	= 192.0	// 3 Vertical Sync, 37 Vertical Blank, 192 Visible Area and 30 Overscan
	SizeYused			float64	= 1.0	// Percentage of the Screen Heigh used by the emulator // 1.0 = 100%, 0.0 = 0%
	// Window Resolution
	ScreenWidth		float64 = 1024
	ScreenHeight		float64 = 768
	// Pixel size
	Width			float64
	Height			float64



	// Monitor Size (to center Window)
	MonitorWidth	float64
	MonitorHeight	float64

	//
	// ----------------------- SaveStates Variables ----------------------- //
	// SavestateFolder		string = "Savestates"
	//
	// // ------------------------ Sound Variables ------------------------- //
	// SpeakerPlaying		bool = false
	// SpeakerStopped		bool = false
	//
	// // ---------------------------- Hybrids ----------------------------- //
	// Hybrid_ETI_660_HW	bool = false

	// TIA
	// Workaround to avoid  WSYNC before VSYNC
	VSYNC_passed		bool = false
)

// Center Window Function
func CenterWindow() {
	winPos := Win.GetPos()
	winPos.X = (MonitorWidth  - float64(ActiveSetting.Mode.Width) ) / 2
	winPos.Y = (MonitorHeight - float64(ActiveSetting.Mode.Height) ) / 2
	Win.SetPos(winPos)
}
