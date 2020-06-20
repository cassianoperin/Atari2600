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
