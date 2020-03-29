package main

import (
	"os"
	"log"
	"fmt"
	"Atari2600/Graphics"
	"Atari2600/CPU"
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

	// Load ROM from 0x200 address in memory
	for i := 0; i < len(data); i++ {
		// F000 - FFFF // Cartridge ROM
		CPU.Memory[0xF000+i] = data[i]
	}

	//Print Memory
	for i := 0; i < len(CPU.Memory); i++ {
		fmt.Printf("%X ", CPU.Memory[i])
	}
	// os.Exit(2)


}

func checkArgs() {

	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s ROM_FILE\n\n", os.Args[0] )
		os.Exit(2)
	}

}

func testFile(filename string) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Printf("File '%s' not found.\n\n", os.Args[1])
		os.Exit(2)
	}
}


func main() {
	fmt.Printf("Atari 2600 Emulator\n")

	// Validate the Arguments
	//checkArgs()

	// Check if file exist
	//testFile(os.Args[1])

	// Set initial variables values
	CPU.Initialize()

	// Read ROM to the memory
	//readROM(os.Args[1])
	readROM("/Users/cassiano/go/src/Atari2600/TestPrograms/cart.bin")


	// Reset system
	CPU.Reset()

	// Start Window System and draw Graphics
	pixelgl.Run(Graphics.Run)
}
