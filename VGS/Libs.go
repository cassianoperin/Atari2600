package VGS

import (
	"fmt"
	"log"
	"os"
)

// ---------------------------- Library Function ---------------------------- //

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
func ReadROM(filename string) {

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
	if romsize == 4096 {
		// Load ROM to memory
		for i := 0; i < len(data); i++ {
			// F000 - FFFF // Cartridge ROM
			Memory[0xF000+i] = data[i]
		}
	}

	// 2KB roms (needs to duplicate it in memory)
	if romsize == 2048 {
		// Load ROM to memory
		for i := 0; i < len(data); i++ {
			// F000 - F7FF (2KB Cartridge ROM)
			Memory[0xF000+i] = data[i]
			// F800 - FFFF (2KB Mirror Cartridge ROM)
			Memory[0xF800+i] = data[i]
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

// Memory Bus - Used by INC, STA, STY and STX to update memory and sinalize TIA about the actions
func memUpdate(memAddr uint16, value byte) {

	// TIA and RIOT
	if memAddr < 128 || memAddr > 0x280 && memAddr <= 0x29F {
		TIA_Update = int16(memAddr)
	}

	// RIOT WRITE ADDRESS
	if memAddr > 0x280 && memAddr <= 0x29F {

		// fmt.Printf("RIOT addr: %02X\n", memAddr)

		// Just update these 2 addresses because I'm filtering the Timer opcodes on STA, STX and STY
		// Update RIOT RW
		Memory_RIOT_RW[memAddr-0x280] = value
		// Update RIOT RW Mirror
		Memory_RIOT_RW[memAddr-0x280+8] = value

		// Print RIOT RW Memory values
		// for i := 0 ; i < len(Memory_RIOT_RW) ; i++ {
		// 	fmt.Printf("%d: %02X\n", i, Memory_RIOT_RW[i])
		// }

		// Update the Timer
		riot_update_timer(memAddr)

		// All other addresses uses regular Memory array
	} else {
		Memory[memAddr] = value
	}

}

// Just TIA can update the Read-only memory space
func update_Memory_TIA_RO(TIAmemAddress, value byte) {

	// TIA will write to first 16 bits and after mirror the values
	if TIAmemAddress >= 16 {
		fmt.Println("TIA - Attempt to write an address >= 16. Exiting.")
		os.Exit(2)
	}

	// TIA RO Memory has 4 mirror in its 64 bits
	for i := 0; i < 4; i++ {
		Memory_TIA_RO[TIAmemAddress+(byte(i)*16)] = value
	}

	// // Print TIA Read Only Memory values
	// for i := 0 ; i < len(Memory_TIA_RO) ; i++ {
	// 	fmt.Printf("%d: %02X\n", i, Memory_TIA_RO[i])
	// }

	// *********************
	// * TIA Documentation *
	// *********************
	//
	// --------------------------------------
	// TIA Addressing Notes for the Atari VCS
	// --------------------------------------
	// A12 is connected to /CS0 (Chip Select 0 - active low)
	// A7 is connected to /CS3  (Chip Select 3 - active low)
	// A[11:8] and A6 are not connected to the TIA
	//
	// VCC is connected to CS1  (Chip Select 1 - active high)
	// GND is connected to /CS2 (Chip Select 2 - active low)
	//
	// --------------------------------------------
	// $0000 - $003F = TIA.......(write)......(read)
	// --------------------------------------------
	// $0000 = TIA Address $00 - (VSYNC)......(CXM0P)
	// $0001 = TIA Address $01 - (VBLANK).....(CXM1P)
	// $0002 = TIA Address $02 - (WSYNC)......(CXP0FB)
	// $0003 = TIA Address $03 - (RSYNC)......(CXP1FB)
	// $0004 = TIA Address $04 - (NUSIZ0).....(CXM0FB)
	// $0005 = TIA Address $05 - (NUSIZ1).....(CXM1FB)
	// $0006 = TIA Address $06 - (COLUP0).....(CXBLPF)
	// $0007 = TIA Address $07 - (COLUP1).....(CXPPMM)
	// $0008 = TIA Address $08 - (COLUPF).....(INPT0)
	// $0009 = TIA Address $09 - (COLUBK).....(INPT1)
	// $000A = TIA Address $0A - (CTRLPF).....(INPT2)
	// $000B = TIA Address $0B - (REFP0)......(INPT3)
	// $000C = TIA Address $0C - (REFP1)......(INPT4)
	// $000D = TIA Address $0D - (PF0)........(INPT5)
	// $000E = TIA Address $0E - (PF1)........(UNDEFINED)
	// $000F = TIA Address $0F - (PF2)........(UNDEFINED)
	// $0010 = TIA Address $10 - (RESP0)......(CXM0P)
	// $0011 = TIA Address $11 - (RESP1)......(CXM1P)
	// $0012 = TIA Address $12 - (RESM0)......(CXP0FB)
	// $0013 = TIA Address $13 - (RESM1)......(CXP1FB)
	// $0014 = TIA Address $14 - (RESBL)......(CXM0FB)
	// $0015 = TIA Address $15 - (AUDC0)......(CXM1FB)
	// $0016 = TIA Address $16 - (AUDC1)......(CXBLPF)
	// $0017 = TIA Address $17 - (AUDF0)......(CXPPMM)
	// $0018 = TIA Address $18 - (AUDF1)......(INPT0)
	// $0019 = TIA Address $19 - (AUDV0)......(INPT1)
	// $001A = TIA Address $1A - (AUDV1)......(INPT2)
	// $001B = TIA Address $1B - (GRP0).......(INPT3)
	// $001C = TIA Address $1C - (GRP1).......(INPT4)
	// $001D = TIA Address $1D - (ENAM0)......(INPT5)
	// $001E = TIA Address $1E - (ENAM1)......(UNDEFINED)
	// $001F = TIA Address $1F - (ENABL)......(UNDEFINED)
	// $0020 = TIA Address $20 - (HMP0).......(CXM0P)
	// $0021 = TIA Address $21 - (HMP1).......(CXM1P)
	// $0022 = TIA Address $22 - (HMM0).......(CXP0FB)
	// $0023 = TIA Address $23 - (HMM1).......(CXP1FB)
	// $0024 = TIA Address $24 - (HMBL).......(CXM0FB)
	// $0025 = TIA Address $25 - (VDELP0).....(CXM1FB)
	// $0026 = TIA Address $26 - (VDELP1).....(CXBLPF)
	// $0027 = TIA Address $27 - (VDELBL).....(CXPPMM)
	// $0028 = TIA Address $28 - (RESMP0).....(INPT0)
	// $0029 = TIA Address $29 - (RESMP1).....(INPT1)
	// $002A = TIA Address $2A - (HMOVE)......(INPT2)
	// $002B = TIA Address $2B - (HMCLR)......(INPT3)
	// $002C = TIA Address $2C - (CXCLR)......(INPT4)
	// $002D = TIA Address $2D - (UNDEFINED)..(INPT5)
	// $002E = TIA Address $2E - (UNDEFINED)..(UNDEFINED)
	// $002F = TIA Address $2F - (UNDEFINED)..(UNDEFINED)
	// $0030 = TIA Address $30 - (UNDEFINED)..(CXM0P)
	// $0031 = TIA Address $31 - (UNDEFINED)..(CXM1P)
	// $0032 = TIA Address $32 - (UNDEFINED)..(CXP0FB)
	// $0033 = TIA Address $33 - (UNDEFINED)..(CXP1FB)
	// $0034 = TIA Address $34 - (UNDEFINED)..(CXM0FB)
	// $0035 = TIA Address $35 - (UNDEFINED)..(CXM1FB)
	// $0036 = TIA Address $36 - (UNDEFINED)..(CXBLPF)
	// $0037 = TIA Address $37 - (UNDEFINED)..(CXPPMM)
	// $0038 = TIA Address $38 - (UNDEFINED)..(INPT0)
	// $0039 = TIA Address $39 - (UNDEFINED)..(INPT1)
	// $003A = TIA Address $3A - (UNDEFINED)..(INPT2)
	// $003B = TIA Address $3B - (UNDEFINED)..(INPT3)
	// $003C = TIA Address $3C - (UNDEFINED)..(INPT4)
	// $003D = TIA Address $3D - (UNDEFINED)..(INPT5)
	// $003E = TIA Address $3E - (UNDEFINED)..(UNDEFINED)
	// $003F = TIA Address $3F - (UNDEFINED)..(UNDEFINED)
}

// Update RIOT
func update_Memory_RIOT_RO(TIAmemAddress, value byte) {

	// // TIA will write to first 16 bits and after mirror the values
	// if TIAmemAddress >= 16 {
	// 	fmt.Println("TIA - Attempt to write an address >= 16. Exiting.")
	// 	os.Exit(2)
	// }
	//
	// // TIA RO Memory has 4 mirror in its 64 bits
	// for i := 0 ; i < 4 ; i++ {
	// 	Memory_TIA_RO[TIAmemAddress + (byte(i) * 16)] = value
	// }

	// --------------------------------
	// RIOT Addresses: names taken from
	// Atari VCS Stella Documentation
	// --------------------------------
	// $0280 = (RIOT $00) - SWCHA  (read/write)
	// $0281 = (RIOT $01) - SWACNT (read/write)
	// $0282 = (RIOT $02) - SWCHB  (read/write) (*)
	// $0283 = (RIOT $03) - SWBCNT (read/write) (*)
	// $0284 = (RIOT $04) - INTIM (read), edge detect control (write)
	// $0285 = (RIOT $05) - read interrupt flag (read), edge detect control (write)
	// $0286 = (RIOT $06) - INTIM (read), edge detect control (write)
	// $0287 = (RIOT $07) - read interrupt flag (read), edge detect control (write)
	// $0288 = (RIOT $08) - SWCHA  (read/write)
	// $0289 = (RIOT $09) - SWACNT (read/write)
	// $028A = (RIOT $0A) - SWCHB  (read/write) (*)
	// $028B = (RIOT $0B) - SWBCNT (read/write) (*)
	// $028C = (RIOT $0C) - INTIM (read) , edge detect control (write)
	// $028D = (RIOT $0D) - read interrupt flag (read), edge detect control (write)
	// $028E = (RIOT $0E) - INTIM (read) , edge detect control (write)
	// $028F = (RIOT $0F) - read interrupt flag (read), edge detect control (write)
	// $0290 = (RIOT $10) - SWCHA  (read/write)
	// $0291 = (RIOT $11) - SWACNT (read/write)
	// $0292 = (RIOT $12) - SWCHB  (read/write) (*)
	// $0293 = (RIOT $13) - SWBCNT (read/write) (*)
	// $0294 = (RIOT $14) - INTIM (read), TIM1T (write)
	// $0295 = (RIOT $15) - read interrupt flag (read), TIM8T (write)
	// $0296 = (RIOT $16) - INTIM (read), TIM64T (write)
	// $0297 = (RIOT $17) - read interrupt flag (read), TIM1024T (write)
	// $0298 = (RIOT $18) - SWCHA  (read/write)
	// $0299 = (RIOT $19) - SWACNT (read/write)
	// $029A = (RIOT $1A) - SWCHB  (read/write) (*)
	// $029B = (RIOT $1B) - SWBCNT (read/write) (*)
	// $029C = (RIOT $1C) - INTIM (read), TIM1T (write)
	// $029D = (RIOT $1D) - read interrupt flag (read), TIM8T (write)
	// $029E = (RIOT $1E) - INTIM (read), TIM64T (write)
	// $029F = (RIOT $1F) - read interrupt flag (read), TIM1024T (write)

}
