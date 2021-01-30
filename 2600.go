package main

import (
	"os"
	"log"
	"fmt"
	"Atari2600/VGS"
	"github.com/faiface/pixel/pixelgl"
)


// Function used by readROM to avoid 'bytesread' return
func ReadContent(file *os.File, bytes_number int) []byte {

	bytes := make([]byte, bytes_number)

	_, err := file.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}

	return bytes
}


// Read ROM and write it to the RAM
func readROM(filename string) {

	var (
		fileInfo os.FileInfo
		err      error
	)

	// Get ROM info
	fileInfo, err = os.Stat(filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Loading ROM:", filename)
	romsize := fileInfo.Size()
	fmt.Printf("Size in bytes: %d\n", romsize)

	// Open ROM file, insert all bytes into memory
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Call ReadContent passing the total size of bytes
	data := ReadContent(file, int(romsize))
	// Print raw data
	//fmt.Printf("%d\n", data)
	//fmt.Printf("%X\n", data)

	// 4KB roms
	if romsize == 4096  {
		// Load ROM to memory
		for i := 0; i < len(data); i++ {
			// F000 - FFFF // Cartridge ROM
			VGS.Memory[0xF000+i] = data[i]
		}
	}

	// 2KB roms (needs to duplicate it in memory)
	if romsize == 2048 {
		// Load ROM to memory
		for i := 0; i < len(data); i++ {
			// F000 - F7FF (2KB Cartridge ROM)
			VGS.Memory[0xF000+i] = data[i]
			// F800 - FFFF (2KB Mirror Cartridge ROM)
			VGS.Memory[0xF800+i] = data[i]
		}
	}

	// // Print Memory -  Fist 2kb
	// for i := 0xF7F0; i <= 0xF7FF; i++ {
	// 	fmt.Printf("%X ", VGS.Memory[i])
	// }
	// fmt.Println()
	// //
	// for i := 0xFFF0; i <= 0xFFFF; i++ {
	// 	fmt.Printf("%X ", VGS.Memory[i])
	// }
	// fmt.Println()

	//Print Memory
	// for i := 0; i < len(VGS.Memory); i++ {
	// 	fmt.Printf("%X ", VGS.Memory[i])
	// }
	// os.Exit(2)
}

func checkArgs() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s ROM_FILE\n\n", os.Args[0] )
		os.Exit(0)
	}
}

func testFile(filename string) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Printf("File '%s' not found.\n\n", os.Args[1])
		os.Exit(0)
	}
}

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
	// readROM(os.Args[1])
	// readROM("/Users/cassiano/go/src/Atari2600/TestPrograms/Bomber/2colorbg.bin")
	// readROM("/Users/cassiano/go/src/Atari2600/TestPrograms/Bomber/3rainbow.bin")
	readROM("/Users/cassiano/go/src/Atari2600/TestPrograms/Bomber/4playfieldborder.bin")
	// readROM("/Users/cassiano/go/src/Atari2600/TestPrograms/Bomber/5playerscoreboard.bin")
	// readROM("/Users/cassiano/go/src/Atari2600/TestPrograms/Bomber/8input.bin")
	// readROM("/Users/cassiano/go/src/Atari2600/TestPrograms/Bomber/105bomber-collision-fixed.bin")

	// Reset system
	VGS.Reset()

	// Start Window System and draw Graphics
	pixelgl.Run(VGS.Run)

}
