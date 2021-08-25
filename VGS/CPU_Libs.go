package VGS

import (
	"fmt"
	"os"
)

// ---------------------------- Library Function ---------------------------- //

// // Function used by ReadROM to avoid 'bytesread' return
// func ReadContent(file *os.File, bytes_number int) []byte {

// 	bytes := make([]byte, bytes_number)

// 	_, err := file.Read(bytes)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return bytes
// }

// // Read ROM and write it to the RAM
// func ReadROM(filename string) {

// 	var (
// 		fileInfo os.FileInfo
// 		err      error
// 	)

// 	// Get ROM info
// 	fileInfo, err = os.Stat(filename)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Loading ROM:", filename)
// 	romsize := fileInfo.Size()
// 	fmt.Printf("Size in bytes: %d\n", romsize)

// 	// Open ROM file, insert all bytes into memory
// 	file, err := os.Open(filename)
// 	defer file.Close()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Call ReadContent passing the total size of bytes
// 	data := ReadContent(file, int(romsize))
// 	// Print raw data
// 	//fmt.Printf("%d\n", data)
// 	//fmt.Printf("%X\n", data)

// 	// // 4KB roms
// 	// if romsize == 4096 {
// 	// 	// Load ROM to memory
// 	// 	for i := 0; i < len(data); i++ {
// 	// 		// F000 - FFFF // Cartridge ROM
// 	// 		VGS.Memory[0xF000+i] = data[i]
// 	// 	}
// 	// }

// 	// // 2KB roms (needs to duplicate it in memory)
// 	// if romsize == 2048 {
// 	// 	// Load ROM to memory
// 	// 	for i := 0; i < len(data); i++ {
// 	// 		// F000 - F7FF (2KB Cartridge ROM)
// 	// 		VGS.Memory[0xF000+i] = data[i]
// 	// 		// F800 - FFFF (2KB Mirror Cartridge ROM)
// 	// 		VGS.Memory[0xF800+i] = data[i]
// 	// 	}
// 	// }

// 	if romsize <= 65536 {
// 		// Load ROM to memory
// 		for i := 0; i < len(data); i++ {
// 			// F000 - F7FF (2KB Cartridge ROM)
// 			// Memory[i] = data[i]
// 			// F800 - FFFF (2KB Mirror Cartridge ROM)
// 			Memory[i] = data[i]
// 		}
// 	} else {
// 		fmt.Printf("\nProgram bigger than 6502 addressable RAM (64KB). Exiting.\n\n")
// 		os.Exit(0)
// 	}

// 	// // Load ROM to memory
// 	// for i := 0; i < len(data); i++ {
// 	// 	// F000 - F7FF (2KB Cartridge ROM)
// 	// 	// Memory[i] = data[i]
// 	// 	// F800 - FFFF (2KB Mirror Cartridge ROM)
// 	// 	Memory[i+0x200] = data[i]
// 	// }

// 	// // Print Memory -  Fist 2kb
// 	// for i := 0xF7F0; i <= 0xF7FF; i++ {
// 	// 	fmt.Printf("%X ", VGS.Memory[i])
// 	// }
// 	// fmt.Println()
// 	// //
// 	// for i := 0xFFF0; i <= 0xFFFF; i++ {
// 	// 	fmt.Printf("%X ", VGS.Memory[i])
// 	// }
// 	// fmt.Println()

// 	// // Print Memory
// 	// for i := 0; i < len(VGS.Memory); i++ {
// 	// 	fmt.Printf("%X ", VGS.Memory[i])
// 	// }
// 	// os.Exit(2)
// }

// Memory Page Boundary cross detection
func MemPageBoundary(original_addr, new_addr uint16) byte {

	var extra_cycle byte = 0

	// Page Boundary Cross detected
	if original_addr>>8 != new_addr>>8 { // Get the High byte only to compare

		extra_cycle = 1

		if Debug {
			fmt.Printf("\tMemory Page Boundary Cross detected! Add 1 cycle.\tPC High byte: %02X\tBranch High byte: %02X\n", original_addr>>8, new_addr>>8)
		}

		// NO Page Boundary Cross detected
	} else {

		extra_cycle = 0

		if Debug {
			fmt.Printf("\tNo Memory Page Boundary Cross detected.\tPC High byte: %02X\tBranch High byte: %02X\n", original_addr>>8, new_addr>>8)
		}
	}

	return extra_cycle
}

// Decode Two's Complement
func DecodeTwoComplement(num byte) int8 {

	var sum int8 = 0

	for i := 0; i < 8; i++ {
		// Sum each bit and sum the value of the bit power of i (<<i)
		sum += (int8(num) >> i & 0x01) << i
	}

	return sum
}

// Decode opcode for debug messages
func debug_decode_opc(bytes uint16) string {

	var opc_string string

	// Decode opcode and operators
	for i := 0; i < int(bytes); i++ {
		if i == 1 {
			opc_string += fmt.Sprintf(" %02X", Memory[PC+uint16(i)])
		} else {
			opc_string += fmt.Sprintf("%02X", Memory[PC+uint16(i)])
		}
	}

	// Insert number of bytes into the string
	if bytes == 1 {
		opc_string += " [1 byte]"
	} else if bytes == 2 {
		opc_string += " [2 bytes]"
	} else {
		opc_string += " [3 bytes]"
	}

	return opc_string
}

// Decode opcode for debug messages
func Debug_decode_console(bytes byte, mem_addr uint16) (string, string, string) {

	var (
		opcode_string            string
		operand_string           string
		operand_bigendian_string string
	)

	// Operator (opcode)
	opcode_string = fmt.Sprintf("%02x", Memory[mem_addr])

	// Decode operators
	for i := 1; i < int(bytes); i++ {
		operand_string += fmt.Sprintf("%02x", Memory[mem_addr+uint16(i)])
	}

	// Decode operators (big endian)
	for i := int(bytes) - 1; i >= 1; i-- {
		operand_bigendian_string += fmt.Sprintf("%02x", Memory[mem_addr+uint16(i)])
	}

	return opcode_string, operand_string, operand_bigendian_string
}

// Print internal opcode cycle in debug mode
func debugInternalOpcCycle(opc_cycles byte) {
	if Debug {
		fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\n", Cycle, Opc_cycle_count, opc_cycles)
	}
}

// Print internal opcode cycle in debug mode - instructions with extra cycle
func debugInternalOpcCycleExtras(opc_cycles byte) {
	if Debug {
		fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\t(%d cycles + %d extra cycles)\n", Cycle, Opc_cycle_count, opc_cycles+Opc_cycle_extra, opc_cycles, Opc_cycle_extra)
	}
}

// Print internal opcode cycle in debug mode - Branches
func debugInternalOpcCycleBranch(opc_cycles byte) {
	if Debug {
		fmt.Printf("\tCPU Cycle: %d\t\tOpcode Cycle %d of %d\t(%d cycles + 1 cycle for branch + %d extra cycles for branch in different page)\n", Cycle, Opc_cycle_count, opc_cycles+Opc_cycle_extra+1, opc_cycles, Opc_cycle_extra)
	}
}

// Reset the internal opcode cycle counters
func resetIntOpcCycleCounters() {
	// Reset Opcode Cycle counter
	Opc_cycle_count = 1

	// Reset Opcode Extra Cycle counter
	Opc_cycle_extra = 0

	NewInstruction = true

	// Update IPS
	IPS++
}

// Data Bus - READ from Memory Operations
func dataBUS_Read(memAddr uint16) byte {

	var data_value byte

	// Read from TIA (14 bits + 2 unused, mirrored 3 more times)
	if memAddr < 64 {
		data_value = Memory_TIA_RO[memAddr]
		// TEMP - Read from other reserved TIA registers
	} else if memAddr < 128 {
		fmt.Printf("dataBUS_Read - Controlled Exit to map access to TIA READ Addresses. COULD BE MIRRORS!!!!!.\t EXITING\n")
		os.Exit(2)
		// Read from RIOT Memory Map (> 0x280)
	} else {
		data_value = Memory[memAddr]
	}

	// data_value := Memory[memAddr]

	return data_value
}

// Data Bus - WRITE to Memory Operations
func dataBUS_Write(memAddr uint16, data_value byte) byte {

	// Update TIA and RIOT
	if memAddr < 128 || memAddr > 0x280 && memAddr <= 0x29F {
		TIA_Update = int16(memAddr)
	}

	// TIA Addresses $00-$3F (write)
	if memAddr >= 0x00 && memAddr <= 0x3F {
		// Memory[memAddr] = data_value

		// TIA Mirrors
		// ***************************************************
		// * $0000-$003F = TIA Addresses $00-$3F (zero page) *
		// * ----------------------------------------------- *
		// *                                                 *
		// *     mirror: $xyz0                               *
		// *                                                 *
		// *     x = {even}                                  *
		// *     y = {anything}                              *
		// *     z = {0, 4}                                  *
		// *                                                 *
		// ***************************************************
		var tia_mirrors = []uint16{0x0000, 0x0040, 0x0100, 0x0140, 0x0200, 0x0240, 0x0300, 0x0340,
			0x0400, 0x0440, 0x0500, 0x0540, 0x0600, 0x0640, 0x0700, 0x0740, 0x0800, 0x0840, 0x0900,
			0x0940, 0x0A00, 0x0A40, 0x0B00, 0x0B40, 0x0C00, 0x0C40, 0x0D00, 0x0D40, 0x0E00, 0x0E40,
			0x0F00, 0x0F40, 0x2000, 0x2040, 0x2100, 0x2140, 0x2200, 0x2240, 0x2300, 0x2340, 0x2400,
			0x2440, 0x2500, 0x2540, 0x2600, 0x2640, 0x2700, 0x2740, 0x2800, 0x2840, 0x2900, 0x2940,
			0x2A00, 0x2A40, 0x2B00, 0x2B40, 0x2C00, 0x2C40, 0x2D00, 0x2D40, 0x2E00, 0x2E40, 0x2F00,
			0x2F40, 0x4000, 0x4040, 0x4100, 0x4140, 0x4200, 0x4240, 0x4300, 0x4340, 0x4400, 0x4440,
			0x4500, 0x4540, 0x4600, 0x4640, 0x4700, 0x4740, 0x4800, 0x4840, 0x4900, 0x4940, 0x4A00,
			0x4A40, 0x4B00, 0x4B40, 0x4C00, 0x4C40, 0x4D00, 0x4D40, 0x4E00, 0x4E40, 0x4F00, 0x4F40,
			0x6000, 0x6040, 0x6100, 0x6140, 0x6200, 0x6240, 0x6300, 0x6340, 0x6400, 0x6440, 0x6500,
			0x6540, 0x6600, 0x6640, 0x6700, 0x6740, 0x6800, 0x6840, 0x6900, 0x6940, 0x6A00, 0x6A40,
			0x6B00, 0x6B40, 0x6C00, 0x6C40, 0x6D00, 0x6D40, 0x6E00, 0x6E40, 0x6F00, 0x6F40, 0x8000,
			0x8040, 0x8100, 0x8140, 0x8200, 0x8240, 0x8300, 0x8340, 0x8400, 0x8440, 0x8500, 0x8540,
			0x8600, 0x8640, 0x8700, 0x8740, 0x8800, 0x8840, 0x8900, 0x8940, 0x8A00, 0x8A40, 0x8B00,
			0x8B40, 0x8C00, 0x8C40, 0x8D00, 0x8D40, 0x8E00, 0x8E40, 0x8F00, 0x8F40, 0xA000, 0xA040,
			0xA100, 0xA140, 0xA200, 0xA240, 0xA300, 0xA340, 0xA400, 0xA440, 0xA500, 0xA540, 0xA600,
			0xA640, 0xA700, 0xA740, 0xA800, 0xA840, 0xA900, 0xA940, 0xAA00, 0xAA40, 0xAB00, 0xAB40,
			0xAC00, 0xAC40, 0xAD00, 0xAD40, 0xAE00, 0xAE40, 0xAF00, 0xAF40, 0xC000, 0xC040, 0xC100,
			0xC140, 0xC200, 0xC240, 0xC300, 0xC340, 0xC400, 0xC440, 0xC500, 0xC540, 0xC600, 0xC640,
			0xC700, 0xC740, 0xC800, 0xC840, 0xC900, 0xC940, 0xCA00, 0xCA40, 0xCB00, 0xCB40, 0xCC00,
			0xCC40, 0xCD00, 0xCD40, 0xCE00, 0xCE40, 0xCF00, 0xCF40, 0xE000, 0xE040, 0xE100, 0xE140,
			0xE200, 0xE240, 0xE300, 0xE340, 0xE400, 0xE440, 0xE500, 0xE540, 0xE600, 0xE640, 0xE700,
			0xE740, 0xE800, 0xE840, 0xE900, 0xE940, 0xEA00, 0xEA40, 0xEB00, 0xEB40, 0xEC00, 0xEC40,
			0xED00, 0xED40, 0xEE00, 0xEE40, 0xEF00, 0xEF40}

		// Update TIA (0x0000-0x003F) and its mirrors
		for _, mirror := range tia_mirrors {
			// fmt.Printf("2**%d = %d\n", i, value)
			// fmt.Printf("%02X\n", mirror)

			Memory[mirror+memAddr] = data_value

			// fmt.Printf("%02X %02X\n\n", Memory[0x83], Memory[0x183])

		}

		// Memory[memAddr] = data_value

		// RIOT RAM $80-$FF
	} else if memAddr >= 0x80 && memAddr <= 0xFF {

		// Define the ram base position
		ram_base := memAddr - 0x80

		// fmt.Printf("ORIGINAL: %02X\t\tBASE: %02X\n", memAddr, ram_base)

		// RAM Mirrors
		// **************************************
		// * $0080-$00FF = RIOT RAM (zero page) *
		// * ---------------------------------- *
		// *                                    *
		// *     mirror: $xy80                  *
		// *                                    *
		// *     x = {even}                     *
		// *     y = {0,1,4,5,8,9,$C,$D}        *
		// *                                    *
		// **************************************
		var ram_mirrors = []uint16{0x0080, 0x0180, 0x0480, 0x0580, 0x0880, 0x0980,
			0x0C80, 0x0D80, 0x2080, 0x2180, 0x2480, 0x2580, 0x2880, 0x2980, 0x2C80,
			0x2D80, 0x4080, 0x4180, 0x4480, 0x4580, 0x4880, 0x4980, 0x4C80, 0x4D80,
			0x6080, 0x6180, 0x6480, 0x6580, 0x6880, 0x6980, 0x6C80, 0x6D80, 0x8080,
			0x8180, 0x8480, 0x8580, 0x8880, 0x8980, 0x8C80, 0x8D80, 0xA080, 0xA180,
			0xA480, 0xA580, 0xA880, 0xA980, 0xAC80, 0xAD80, 0xC080, 0xC180, 0xC480,
			0xC580, 0xC880, 0xC980, 0xCC80, 0xCD80, 0xE080, 0xE180, 0xE480, 0xE580,
			0xE880, 0xE980, 0xEC80, 0xED80}

		// Update RAM (0x0080-0x00FF) and its mirrors
		for _, mirror := range ram_mirrors {
			// fmt.Printf("2**%d = %d\n", i, value)
			// fmt.Printf("%02X\n", mirror)

			Memory[mirror+ram_base] = data_value

			// fmt.Printf("%02X %02X\n\n", Memory[0x83], Memory[0x183])

		}

		// RIOT WRITE ADDRESS
	} else if memAddr > 0x280 && memAddr <= 0x29F {

		// fmt.Printf("RIOT addr: %02X\n", memAddr)

		// Just update these 2 addresses because I'm filtering the Timer opcodes on STA, STX and STY
		// Update RIOT RW
		Memory_RIOT_RW[memAddr-0x280] = data_value
		// Update RIOT RW Mirror
		Memory_RIOT_RW[memAddr-0x280+8] = data_value

		// Print RIOT RW Memory values
		// for i := 0 ; i < len(Memory_RIOT_RW) ; i++ {
		// 	fmt.Printf("%d: %02X\n", i, Memory_RIOT_RW[i])
		// }

		// Update the Timer
		riot_update_timer(memAddr)

		// All other addresses uses regular Memory array
	} else {
		Memory[memAddr] = data_value
	}

	return data_value
}
