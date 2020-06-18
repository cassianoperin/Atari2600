package CPU

import	"fmt"

// Relative
func addr_mode_Relative(offset uint16) int8 {

	// Branches needs the Two Complement of the offset value
	value := DecodeTwoComplement(Memory[offset])

	if Debug {
		fmt.Printf("\n\tRelative addressing mode.\tOffset(PC+1): %d\tValue : %d", offset, value)
	}

	return value

}
