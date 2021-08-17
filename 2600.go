package main

import (
	"fmt"

	"github.com/faiface/pixel/pixelgl"

	"Atari2600/VGS"
)

// func checkArgs() {
// 	if len(os.Args) != 2 {
// 		fmt.Printf("Usage: %s ROM_FILE\n\n", os.Args[0])
// 		os.Exit(0)
// 	}
// }

// func testFile(filename string) {
// 	if _, err := os.Stat(filename); os.IsNotExist(err) {
// 		fmt.Printf("File '%s' not found.\n\n", os.Args[1])
// 		os.Exit(0)
// 	}
// }

func main() {
	fmt.Printf("Atari 2600 Emulator\n")

	// Validate the Arguments
	//checkArgs()

	// Check if file exist
	//testFile(os.Args[1])

	// Set initial variables values
	VGS.Initialize()

	// Initialize Timers
	VGS.InitializeTimers()

	// Read ROM to the memory
	// CORE.ReadROM(flag.Arg(0))
	// VGS.ReadROM("/Users/cassiano/go/src/6502/TestPrograms/6502_functional_test.bin")
	// VGS.ReadROM("/Users/cassiano/go/src/6502/TestPrograms/6502_decimal_test.bin")
	// Read ROM to the memory
	// VGS.ReadROM(os.Args[1])
	// VGS.ReadROM("/Users/cassiano/go/src/Atari2600/TestPrograms/Bomber/2colorbg.bin")
	VGS.ReadROM("/Users/cassiano/go/src/Atari2600/TestPrograms/Bomber/3rainbow.bin")
	// VGS.ReadROM("/Users/cassiano/go/src/Atari2600/TestPrograms/Bomber/4playfieldborder.bin")
	// VGS.ReadROM("/Users/cassiano/go/src/Atari2600/TestPrograms/Bomber/5playerscoreboard.bin")
	// VGS.ReadROM("/Users/cassiano/go/src/Atari2600/TestPrograms/Bomber/6vertical.bin")
	// VGS.ReadROM("/Users/cassiano/go/src/Atari2600/TestPrograms/Bomber/7horizontal-fixed.bin")
	// VGS.ReadROM("/Users/cassiano/go/src/Atari2600/TestPrograms/Bomber/8input.bin")
	// VGS.ReadROM("/Users/cassiano/go/src/Atari2600/TestPrograms/Bomber/105bomber-collision-fixed.bin")
	// VGS.ReadROM("/Users/cassiano/go/src/Atari2600/TestPrograms/Surround.bin")
	// VGS.ReadROM("/Users/cassiano/go/src/Atari2600/TestPrograms/Pac-Man.bin")
	// VGS.ReadROM("/Users/cassiano/go/src/Atari2600/TestPrograms/cart.bin")

	// Reset system
	VGS.Reset()

	// Start Window System and draw Graphics
	pixelgl.Run(VGS.Run)

}
