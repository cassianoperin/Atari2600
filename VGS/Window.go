package VGS

import (
	"fmt"
	"time"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
	"github.com/faiface/pixel/pixelgl"
)

func renderGraphics() {

	// Initial Pixel Size
	width		= screenWidth  / sizeX
	height		= screenHeight / sizeY

	cfg := pixelgl.WindowConfig{
		Title:  windowTitle,
		Bounds: pixel.R(0, 0, screenWidth, screenHeight),
		VSync:  false,
		Resizable: false,
		Undecorated: false,
		NoIconify: false,
		AlwaysOnTop: false,
	}
	var err error
	win, err = pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	// Disable Smooth
	win.SetSmooth(true)

	// Fullscreeen and video resolution - Retrieve all monitors
	monitors := pixelgl.Monitors()

	// Map the video modes available
	for i := 0; i < len(monitors); i++ {
		// Retrieve all video modes for a specific monitor.
		modes := monitors[i].VideoModes()
		for j := 0; j < len(modes); j++ {
			settings = append(settings, Setting{
				Monitor: monitors[i],
				Mode:    &modes[j],
			})
		}

		// Determine monitor size in pixels to center the window
		monitorWidth, monitorHeight = monitors[i].Size()
		// fmt.Printf("-size: %v px, %v px\n", MonitorWidth, MonitorHeight)
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
	activeSetting = &settings[3]

	if isFullScreen {
		win.SetMonitor(activeSetting.Monitor)
	} else {
		win.SetMonitor(nil)
	}
	win.SetBounds(pixel.R(0, 0, float64(activeSetting.Mode.Width), float64(activeSetting.Mode.Height)))

	// Center Window
	// CenterWindow()
	// winPos := win.GetPos()
	// winPos.X = (MonitorWidth  - float64(activeSetting.Mode.Width) ) / 2
	// winPos.Y = (MonitorHeight - float64(activeSetting.Mode.Height) ) / 2
	// win.SetPos(winPos)

	//Initialize FPS Text
	// textFPS	= text.New(pixel.V(10, 470), atlas)
	//Initialize Messages Text
	// textMessage	= text.New(pixel.V(10, 10) , atlas)
	textMessage	= text.New(pixel.V(10, 10 ) , atlas)
	// Initialize CPU Debug Message
	cpuMessage = text.New(pixel.V(10, 150), atlas)
}


// Infinte Loop
func RunInfiniteLoop() {

	// Set up render system
	renderGraphics()

	if Debug {
		startDebug()
	}

	// Main Infinite Loop
	for !win.Closed() {

		// Esc to quit program
		if win.Pressed(pixelgl.KeyEscape) {
			break
		}

		// Every Cycle Control the clock!!!
		select {
		case <- clock_timer.C:

				// Handle Input
				Keyboard()

				if !Pause {
					// Time measurement - CPU Cycle
					if debugTiming {
						debugTiming_StartCycle = time.Now()
					}

					// Runs the interpreter
					CPU_Interpreter()

					// Draw the pixels on the monitor accordingly to beam update (1 CPU cycle = 3 TIA color clocks)
					TIA( TIA_Update )

					// Reset Controllers Buttons to 1 (not pressed)
					Memory[SWCHA] = 0xFF //1111 11111

					// Time measurement - CPU Cycle
					if debugTiming {
						elapsed := time.Since(debugTiming_StartCycle)
						if elapsed.Seconds() > debugTiming_Limit {
							fmt.Printf("\nTiming: Opcode: %X\tEntire CYCLE took %f seconds\n", opcode, elapsed.Seconds())
							// Pause = true
						}
					}

					// Update Debug Screen
					if Debug {
						updateDebug()
					}

				}

				case <- second_timer: // Second
					win.SetTitle(fmt.Sprintf("%s IPS: %d| FPS: %d | Draws: %d", cfg.Title, counter_IPS, counter_FPS, counter_DPS))
					counter_IPS = 0
					counter_FPS = 0
					counter_DPS  = 0

				case <- screenRefresh_timer.C:
					// When ticker run (60 times in a second, Refresh the screen)
					win.Update()

				case <- messagesClock_timer.C:
					// After some time, stop showing the message
					ShowMessage = false

				default:
					// No timer to handle
			}


		// -------------------------- SBC, CARRY and Overflow tests -------------------------- //
		// P[0] = 1

		// A = 135
		// Memory[61832] = 15

		// Test 1
		// A = 0xF8
		// Memory[61832] = 0x0A
		// Test 2
		// A = 0x81
		// Memory[61832] = 0x07
		// Test 3
		// A = 0x07
		// Memory[61832] = 0x02
		// Test 4
		// A = 0x07
		// Memory[61832] = 0xFE // Hexadecimal de -2 (twos)
		// Test 5
		// A = 0x07
		// Memory[61832] = 0x09
		// Test 6
		// A = 0x07
		// Memory[61832] = 0x90
		// Test 7
		// A = 0x10
		// Memory[61832] = 0x90
		// Test 8
		// A = 0x10
		// Memory[61832] = 0x91

		// A = 80
		// Memory[61832] = 240
		// A = 80
		// Memory[61832] = 176
		// A = 208
		// Memory[61832] = 48
		//
		// fmt.Printf("\n\nOPERATION: A: %d (0x%02X) - %d (0x%02X)", A, A, Memory[61832], Memory[61832])
		// PC = 0xF187
		// Interpreter()
		// fmt.Printf("\n\nResult: %d (0x%02X)\tTwo's Complement: %d (0x%02X)\tOverflow: %d\tCarry: %d\n\n", A, A, DecodeTwoComplement(A), DecodeTwoComplement(A), P[6], P[0])
		//
		// os.Exit(2)
		// ----------------------------------------------------------------------------------- //

		// // Avoid high CPU usage for the loop
		// if Pause {
		// 	time.Sleep(50 * time.Millisecond)
		// }

	}

}

// Center Window Function
func CenterWindow() {
	winPos := win.GetPos()
	winPos.X = (monitorWidth  - float64(activeSetting.Mode.Width) ) / 2
	winPos.Y = (monitorHeight - float64(activeSetting.Mode.Height) ) / 2
	win.SetPos(winPos)
}
