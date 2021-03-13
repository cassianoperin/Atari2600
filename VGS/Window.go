package VGS

import (
	"os"
	"fmt"
	"time"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	// "golang.org/x/image/colornames"
	// "github.com/faiface/pixel/imdraw"

)

var (
	// second_timer		= time.Tick(time.Nanosecond)			// 1 second to track FPS and draws
	// second_timer		= time.Tick(time.Second / 30)			// 1 second to track FPS and draws
	// second_timer		= time.NewTicker(time.Nanosecond)			// 1 second to track FPS and draws
 	// cycle int = 0
	)

// func limpa(teste *pixelgl.Window) {
// 	teste.Clear(colornames.Red)
//
// }

func Run() {

	// Set up render system
	// Initial Pixel Size
	width		= screenWidth  / sizeX
	height		= screenHeight / sizeY

	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  false,
		Resizable: false,
		Undecorated: false,
		NoIconify: false,
		AlwaysOnTop: false,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	// Disable Smooth
	win.SetSmooth(true)

	// renderGraphics(win)

	if Debug {
		// XXX
		// startDebug(win)
	}




	// limpa(win)
	// win.Clear(colornames.Skyblue)

	// for !win.Closed() {
	//
	// 	select {
	// 		case <- second_timer: // Second
	// 			win.Update()
	// 			fmt.Println(cycle)
	// 			cycle = 0
	//
	// 		default:
	// 		// No timer to handle
	// 	}
	//
	// 	cycle++
	// }



	for !win.Closed() {

		// // Esc to quit program
		// if win.JustPressed(pixelgl.KeyEscape) {
		// 	break
		// }

			// Internal Loop to avoid slowness of !win.Closed() loop
			for i:=0 ; i < 50000 ; i++ {

				// Esc to quit program
				if win.JustPressed(pixelgl.KeyEscape) {
					os.Exit(0)
				}


				select {
					case <- second_timer: // Second
						win.SetTitle(fmt.Sprintf("%s IPS: %d| FPS: %d | Draws: %d", cfg.Title, counter_IPS, counter_FPS, counter_DPS))
						counter_IPS = 0
						counter_FPS = 0
						counter_DPS  = 0

					default:
						// No timer to handle
				}

				// Handle RIOT Timer
				riot_timer_counter ++

				if riot_timer_counter == riot_timer_mult {
					old_timer = riot_timer // used to know if returned to 255
					riot_timer --
					riot_timer_counter = 0

					if riot_timer > old_timer {
						// fmt.Println("Zerou!")
						Memory[TIMINT] = 128
						riot_timer_mult = 1
					}
					Memory[INTIM] = riot_timer
				}
				// fmt.Println(riot_timer)

				// fmt.Printf("riot_timer_counter: %d\t\told_timer: %d\triot_timer: %d\tMemory[INTIM] : %d\tMemory[TIMINT]: %d\n", riot_timer_counter, old_timer, riot_timer, Memory[INTIM], Memory[TIMINT])



				// Handle Input
				Keyboard(win)



				// select {
				// case <- clock_timer.C:

						if !Pause {
							// Time measurement - CPU Cycle
							if debugTiming {
								debugTiming_StartCycle = time.Now()
							}

							// Runs the interpreter
							if CPU_Enabled {
								CPU_Interpreter()
							} else {
								// Increment Cycle
								counter_F_Cycle ++
							}

							// Draw the pixels on the monitor accordingly to beam update (1 CPU cycle = 3 TIA color clocks)
							TIA( TIA_Update, win )
							// fmt.Printf("Cycle: %d\t\tLine: %d\n", counter_F_Cycle, line)

							// // Update Debug Screen
							// if Debug {
							// 	updateDebug()
							//
							// 	if dbg_break {
							// 		if counter_F_Cycle == dbg_break_cycle {
							// 			Pause = true
							// 		}
							// 	}
							// }

						// Refresh screen if in Pause mode
						} else {
							win.Update()
						}

			// 	default:
			// 		// No timer to handle
			// }


				select {
					case <- messagesClock_timer.C:
						// After some time, stop showing the message
						ShowMessage = false
					default:
						// No timer to handle
				}


				// ---------------- Reset Physical Switches ---------------- //
				// Reset switch not enabled (put 1 on position 0 of SWCHB)
				Memory[SWCHB] |= (1 << 0)
				// Game select switch not enabled (put 1 on position 1 of SWCHB)
				Memory[SWCHB] |= (1 << 1)

			}

	}
}
