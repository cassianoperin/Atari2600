package VGS

import (
	"fmt"
	"os"
	"time"

	CPU_6502 "github.com/cassianoperin/6502"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	// "Atari2600/CPU_6502"
)

var (
// second_timer		= time.Tick(time.Nanosecond)			// 1 second to track FPS and draws
// second_timer		= time.Tick(time.Second / 30)			// 1 second to track FPS and draws
// second_timer		= time.NewTicker(time.Nanosecond)			// 1 second to track FPS and draws
// cycle int = 0
)

func Run() {

	// Set up render system
	// Initial Pixel Size
	width = screenWidth / sizeX
	height = screenHeight / sizeY

	cfg := pixelgl.WindowConfig{
		Title:       "Pixel Rocks!",
		Bounds:      pixel.R(0, 0, 1024, 768),
		VSync:       false,
		Resizable:   false,
		Undecorated: false,
		NoIconify:   false,
		AlwaysOnTop: false,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	// Disable Smooth
	win.SetSmooth(true)

	// renderGraphics(win)

	// if CPU_6502.Debug {
	// 	// XXX
	// 	// startDebug(win)
	// }

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
		for i := 0; i < 1000; i++ {

			// Esc to quit program
			if win.JustPressed(pixelgl.KeyEscape) {
				os.Exit(0)
			}

			select {
			case <-second_timer: // Second
				win.SetTitle(fmt.Sprintf("%s IPS: %d| FPS: %d | Draws: %d", cfg.Title, CPU_6502.IPS, counter_FPS, counter_DPS))
				CPU_6502.IPS = 0
				counter_FPS = 0
				counter_DPS = 0

			default:
				// No timer to handle
			}

			// Handle RIOT Timer
			riot_timer_counter++

			if riot_timer_counter == riot_timer_mult {
				old_timer = riot_timer // used to know if returned to 255
				riot_timer--
				riot_timer_counter = 0

				if riot_timer > old_timer {
					// fmt.Println("Zerou!")
					CPU_6502.Memory[TIMINT] = 128
					riot_timer_mult = 1
				}
				CPU_6502.Memory[INTIM] = riot_timer
			}
			// fmt.Println(riot_timer)

			// fmt.Printf("riot_timer_counter: %d\t\told_timer: %d\triot_timer: %d\tMemory[INTIM] : %d\tMemory[TIMINT]: %d\n", riot_timer_counter, old_timer, riot_timer, Memory[INTIM], Memory[TIMINT])

			// Handle Input
			Keyboard(win)

			// select {
			// case <- clock_timer.C:

			if !CPU_6502.Pause {
				// Time measurement - CPU Cycle
				if debugTiming {
					debugTiming_StartCycle = time.Now()
				}

				// Runs the interpreter
				if CPU_6502.CPU_Enabled {

					// Increment the beam
					beamIndex++

					CPU_6502.CPU_Interpreter()

					// Set it all the times to be ignored
					TIA_Update = -1

					fmt.Printf("\n\n\n\n\n\n\tAddress BUS: %d\n\n\n\n", CPU_6502.AddressBUS)

					if CPU_6502.Opc_cycle_count == CPU_6502.Opc_cycles+CPU_6502.Opc_cycle_extra {
						fmt.Println("Update TIAAA (add extra")
						// EXPORTAR MEMADDR, e no último ciclo, atualizar tia?
						TIA_Update = int16(CPU_6502.AddressBUS)
					}

					fmt.Printf("\n\tBeam Index: %d\n", beamIndex)
					win.Update()

				} else {
					// Increment Cycle
					counter_F_Cycle++
				}

				// Draw the pixels on the monitor accordingly to beam update (1 CPU cycle = 3 TIA color clocks)
				TIA(TIA_Update, win)
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

			// select {
			// case <-messagesClock_timer.C:
			// 	// After some time, stop showing the message
			// 	ShowMessage = false
			// default:
			// 	// No timer to handle
			// }

			// ---------------- Reset Physical Switches ---------------- //
			// Reset switch not enabled (put 1 on position 0 of SWCHB)
			CPU_6502.Memory[SWCHB] |= (1 << 0)
			// Game select switch not enabled (put 1 on position 1 of SWCHB)
			CPU_6502.Memory[SWCHB] |= (1 << 1)

		}

	}
}
