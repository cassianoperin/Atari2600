package CPU

import	"fmt"

// Relative
func addr_mode_Relative(offset uint16) uint16 {

	// Branches needs the Two Complement of the offset value
	value := uint16( DecodeTwoComplement(Memory[offset]) )

	if Debug {
		fmt.Printf("\n\tRelative addressing mode.\tMemory[%02X]\tValue obtained: %d", offset, value)
	}

	return value
}


// Zeropage
func addr_mode_Zeropage(offset uint16) (uint16, string) {

	value	:= Memory[ Memory[offset] ]
	memAddr 	:= uint16(Memory[offset])
	mode		:= "Zeropage"

	if Debug {
		fmt.Printf("\n\t%s addressing mode.\tMemory[%02X]\tValue obtained: %d", mode, memAddr, value)
	}

	return memAddr, mode
}


// // Zeropage DOOOOOOIS
// // NECESSÁRIO PARA 7-horizontal no LDY
// func addr_mode_Zeropage2(offset uint16) (uint16, string) {
//
// 	value	:= Memory[ offset ]
// 	memAddr 	:= offset
// 	mode		:= "Zeropage 2"
//
// 	if Debug {
// 		fmt.Printf("\n\t%s addressing mode.\tMemory[%02X]\tValue obtained: %d", mode, memAddr, value)
// 	}
//
// 	return memAddr, mode
// }


// Immediate
func addr_mode_Immediate(offset uint16) (uint16, string) {

	value	:= Memory[offset]
	memAddr	:= offset
	mode		:= "Immediate"

	if Debug {
		fmt.Printf("\n\t%s addressing mode.\tMemory[%02X]\tValue obtained: %d", mode, memAddr, value)
	}

	return memAddr, mode
}


// Absolute
func addr_mode_Absolute(offset uint16) (uint16, string) {

	memAddr := uint16(Memory[offset+1])<<8 | uint16(Memory[offset])
	value := Memory[memAddr]
	mode		:= "Absolute"

	if Debug {
		fmt.Printf("\n\t%s addressing mode.\tMemory[%02X]\t\tValue obtained: %d", mode, memAddr, value)
	}
	return memAddr, mode
}

// Absolute,Y
func addr_mode_AbsoluteY	(offset uint16) (uint16, string) {

	memAddr := ( uint16(Memory[offset+1])<<8 | uint16(Memory[offset]) ) + uint16(Y)
	value := Memory[memAddr]
	mode		:= "Absolute,Y"

	if Debug {
		fmt.Printf("\n\t%s addressing mode.\tMemory[%02X]\t\tValue obtained: %d", mode, memAddr, value)
	}
	return memAddr, mode
}


// Indirect,Y
func addr_mode_IndirectY	(offset uint16) (uint16, string) {

	memAddr  := ( uint16(Memory[Memory[offset] + 1])<<8 | uint16(Memory[Memory[offset]]) ) + uint16(Y)
	value := Memory[memAddr]
	mode		:= "Indirect,Y"

	if Debug {
		fmt.Printf("\n\t%s addressing mode.\tMemory[%02X]\t\tValue obtained: %d", mode, memAddr, value)
	}
	return memAddr, mode
}
