package Graphics

import (
	"fmt"
	"time"
	"Atari2600/CPU"
	"Atari2600/Global"
	"Atari2600/Input"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/font/basicfont"
)

var (
	// FPS
	textFPS		*text.Text	// On screen FPS counter
	textFPSstr	string		// String with the FPS counter
	drawCounter	= 0		// imd.Draw per second counter
	updateCounter	= 0		// Win.Updates per second counter

	// Screen messages
	textMessage	*text.Text	// On screen Message content
	cpuMessage  *text.Text	// In screen CPU components debug
	// Fonts
	atlas = text.NewAtlas(basicfont.Face7x13, text.ASCII)

	// Screen Size
	// screenWidth	= float64(1024)
	// screenHeight	= float64(768)
	screenWidth	= float64(640)
	screenHeight	= float64(480)

	width			= screenWidth  / Global.SizeX
	height			= screenHeight / Global.SizeY

	// Window Configuration
	// win				* pixelgl.Window
	imd				= imdraw.New(nil)
	cfg				= pixelgl.WindowConfig{}

	// Debug mode
	debug			bool = false
)

// const (
	// sizeX			float64	= 160.0 	// 68 color clocks (Horizontal Blank) + 160 color clocks (pixels)
	// sizeY			float64	= 192.0	// 3 Vertical Sync, 37 Vertical Blank, 192 Visible Area and 30 Overscan



// )


func renderGraphics() {
	cfg := pixelgl.WindowConfig{
		Title:  Global.WindowTitle,
		Bounds: pixel.R(0, 0, screenWidth, screenHeight),
		VSync:  false,
		Resizable: false,
		Undecorated: false,
		NoIconify: false,
		AlwaysOnTop: true,
	}
	var err error
	Global.Win, err = pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	// Disable Smooth
	Global.Win.SetSmooth(true)

	// Fullscreeen and video resolution - Retrieve all monitors
	monitors := pixelgl.Monitors()

	// Map the video modes available
	for i := 0; i < len(monitors); i++ {
		// Retrieve all video modes for a specific monitor.
		modes := monitors[i].VideoModes()
		for j := 0; j < len(modes); j++ {
			Global.Settings = append(Global.Settings, Global.Setting{
				Monitor: monitors[i],
				Mode:    &modes[j],
			})
		}

		// Determine monitor size in pixels to center the window
		Global.MonitorWidth, Global.MonitorHeight = monitors[i].Size()
		// fmt.Printf("-size: %v px, %v px\n", Global.MonitorWidth, Global.MonitorHeight)
	}

	// Complete monitor info
	// for i, m := range monitors {
	//
	// 		// fmt.Printf("monitor %v:\n", i)
	// 		//
	// 		// name := m.Name()
	// 		// fmt.Printf("-name: %v\n", name)
	// 		//
	// 		// bitDepthRed, bitDepthGreen, bitDepthBlue := m.BitDepth()
	// 		// fmt.Printf("-bitDepth: %v-bit red, %v-bit green, %v-bit blue\n",
	// 		// 	bitDepthRed, bitDepthGreen, bitDepthBlue)
	// 		//
	// 		// physicalSizeWidth, physicalSizeHeight := m.PhysicalSize()
	// 		// fmt.Printf("-physicalSize: %v mm, %v mm\n",
	// 		// 	physicalSizeWidth, physicalSizeHeight)
	// 		//
	// 		// positionX, positionY := m.Position()
	// 		// fmt.Printf("-position: %v, %v upper-left corner\n",
	// 		// 	positionX, positionY)
	// 		//
	// 		// refreshRate := m.RefreshRate()
	// 		// fmt.Printf("-refreshRate: %v Hz\n", refreshRate)
	//
	// 		sizeWidth, sizeHeight := m.Size()
	// 		fmt.Printf("-size: %v px, %v px\n",
	// 			sizeWidth, sizeHeight)
	//
	// 		// videoModes := m.VideoModes()
	// 		//
	// 		// for j, vm := range videoModes {
	// 		//
	// 		// 	fmt.Printf("-video mode %v: -width: %v px, height: %v px, refresh rate:%v Hz\n",
	// 		// 		j, vm.Width, vm.Height, vm.RefreshRate)
	// 		//
	// 		// }
	// 	}

	// Set Initial resolution
	Global.ActiveSetting = &Global.Settings[3]

	if Global.IsFullScreen {
		Global.Win.SetMonitor(Global.ActiveSetting.Monitor)
	} else {
		Global.Win.SetMonitor(nil)
	}
	Global.Win.SetBounds(pixel.R(0, 0, float64(Global.ActiveSetting.Mode.Width), float64(Global.ActiveSetting.Mode.Height)))

	// Center Window
	Global.CenterWindow()
	// winPos := Global.Win.GetPos()
	// winPos.X = (Global.MonitorWidth  - float64(Global.ActiveSetting.Mode.Width) ) / 2
	// winPos.Y = (Global.MonitorHeight - float64(Global.ActiveSetting.Mode.Height) ) / 2
	// Global.Win.SetPos(winPos)

	//Initialize FPS Text
	textFPS	= text.New(pixel.V(10, 470), atlas)
	//Initialize Messages Text
	// textMessage	= text.New(pixel.V(10, 10) , atlas)
	textMessage	= text.New(pixel.V(10, 10 ) , atlas)
	// Initialize CPU Debug Message
	cpuMessage = text.New(pixel.V(10, 150), atlas)
}


// Infinte Loop
func Run() {

	CD_P0_P1		= [160]byte{}

	// Set up render system
	renderGraphics()

	// Main Infinite Loop
	for !Global.Win.Closed() {


		// Esc to quit program
		if Global.Win.Pressed(pixelgl.KeyEscape) {
			break
		}

		// Every Cycle Control the clock!!!
		select {
			case <-CPU.Clock.C:

				width		= screenWidth/Global.SizeX
				height		= screenHeight/Global.SizeY * Global.SizeYused	// Define the heigh of the pixel, considering the percentage of screen reserved for emulator


				// Handle Input
				Input.Keyboard()

				if !CPU.Pause {
					// Time measurement - CPU Cycle
					if CPU.DebugTiming {
						CPU.StartCycle = time.Now()
					}

					// Runs the interpreter
					CPU.Interpreter()

					// Draw the pixels on the monitor accordingly to beam update (1 CPU cycle = 3 TIA color clocks)
					TIA( CPU.TIA_Update )

					// Reset Controllers Buttons to 1 (not pressed)
					CPU.Memory[CPU.SWCHA] = 0xFF //1111 11111

					// Time measurement - CPU Cycle
					if CPU.DebugTiming {
						elapsed := time.Since(CPU.StartCycle)
						if elapsed.Seconds() > CPU.DebugTimingLimit {
							fmt.Printf("\nTiming: Opcode: %X\tEntire CYCLE took %f seconds\n", CPU.Opcode, elapsed.Seconds())
							// CPU.Pause = true
						}
					}
				}

				case <-CPU.Second: // Second
					Global.Win.SetTitle(fmt.Sprintf("%s IPS: %d| FPS: %d | Draws: %d", cfg.Title, CPU.IPS, frames, draws))
					CPU.IPS = 0
					frames = 0
					draws  = 0

				case <-CPU.ScreenRefresh.C:
					// When ticker run (60 times in a second, Refresh the screen)
					Global.Win.Update()

				case <-CPU.MessagesClock.C:
					// After some time, stop showing the message
					Global.ShowMessage = false

				default:
					// No timer to handle
			}


		// -------------------------- SBC, CARRY and Overflow tests -------------------------- //
		// CPU.P[0] = 1

		// CPU.A = 135
		// CPU.Memory[61832] = 15

		// Test 1
		// CPU.A = 0xF8
		// CPU.Memory[61832] = 0x0A
		// Test 2
		// CPU.A = 0x81
		// CPU.Memory[61832] = 0x07
		// Test 3
		// CPU.A = 0x07
		// CPU.Memory[61832] = 0x02
		// Test 4
		// CPU.A = 0x07
		// CPU.Memory[61832] = 0xFE // Hexadecimal de -2 (twos)
		// Test 5
		// CPU.A = 0x07
		// CPU.Memory[61832] = 0x09
		// Test 6
		// CPU.A = 0x07
		// CPU.Memory[61832] = 0x90
		// Test 7
		// CPU.A = 0x10
		// CPU.Memory[61832] = 0x90
		// Test 8
		// CPU.A = 0x10
		// CPU.Memory[61832] = 0x91

		// CPU.A = 80
		// CPU.Memory[61832] = 240
		// CPU.A = 80
		// CPU.Memory[61832] = 176
		// CPU.A = 208
		// CPU.Memory[61832] = 48
		//
		// fmt.Printf("\n\nOPERATION: A: %d (0x%02X) - %d (0x%02X)", CPU.A, CPU.A, CPU.Memory[61832], CPU.Memory[61832])
		// CPU.PC = 0xF187
		// CPU.Interpreter()
		// fmt.Printf("\n\nResult: %d (0x%02X)\tTwo's Complement: %d (0x%02X)\tOverflow: %d\tCarry: %d\n\n", CPU.A, CPU.A, CPU.DecodeTwoComplement(CPU.A), CPU.DecodeTwoComplement(CPU.A), CPU.P[6], CPU.P[0])
		//
		// os.Exit(2)
		// ----------------------------------------------------------------------------------- //

	}

}
