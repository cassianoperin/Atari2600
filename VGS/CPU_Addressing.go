package VGS

import	"fmt"

// Relative
func addr_mode_Relative(offset uint16) int8 {

	// Branches needs the Two Complement of the offset value
	value := DecodeTwoComplement(Memory[offset])
	memAddr 	:= offset
	mode		:= "Relative"

	if Debug {
		fmt.Printf("\t%s addressing mode.\tMemory[%02X]\tValue obtained: %d (SIGNED value)\n", mode, memAddr, value)
	}

	return value
}


// Zeropage
func addr_mode_Zeropage(offset uint16) (uint16, string) {

	value	:= Memory[ Memory[offset] ]
	memAddr 	:= Memory[offset]
	mode		:= "Zeropage"

	if Debug {
		fmt.Printf("\t%s addressing mode.\tMemory[%02X]\tValue obtained: %d\n", mode, memAddr, value)
	}

	return uint16(memAddr), mode
}


// Zeropage,X
func addr_mode_ZeropageX(offset uint16) (uint16, string) {

	value	:= Memory[ Memory[offset] + X ]
	memAddr 	:= Memory[offset] + X
	mode		:= "Zeropage,X"

	if Debug {
		fmt.Printf("\t%s addressing mode.\tMemory[%02X]\tValue obtained: %d\n", mode, memAddr, value)
	}

	return uint16(memAddr), mode
}

// Immediate
func addr_mode_Immediate(offset uint16) (uint16, string) {

	value	:= Memory[offset]
	memAddr	:= offset
	mode		:= "Immediate"

	if Debug {
		fmt.Printf("\t%s addressing mode.\tMemory[%02X]\tValue obtained: %d\n", mode, memAddr, value)
	}

	return memAddr, mode

}


// Absolute
func addr_mode_Absolute(offset uint16) (uint16, string) {

	memAddr := uint16(Memory[offset+1])<<8 | uint16(Memory[offset])
	value := Memory[memAddr]
	mode		:= "Absolute"

	if Debug {
		fmt.Printf("\t%s addressing mode.\tMemory[%02X]\t\tValue obtained: %d\n", mode, memAddr, value)
	}
	return memAddr, mode
}


// Absolute,Y
func addr_mode_AbsoluteY	(offset uint16) (uint16, string) {

	memAddr := ( uint16(Memory[offset+1])<<8 | uint16(Memory[offset]) ) + uint16(Y)
	value := Memory[memAddr]
	mode		:= "Absolute,Y"

	if Debug {
		fmt.Printf("\t%s addressing mode.\t\tMemory[%02X]\t\tValue obtained: %d\n", mode, memAddr, value)
	}
	return memAddr, mode
}

// Absolute,X
func addr_mode_AbsoluteX	(offset uint16) (uint16, string) {

	memAddr := ( uint16(Memory[offset+1])<<8 | uint16(Memory[offset]) ) + uint16(X)
	value := Memory[memAddr]
	mode		:= "Absolute,X"

	if Debug {
		fmt.Printf("\t%s addressing mode.\t\tMemory[%02X]\t\tValue obtained: %d\n", mode, memAddr, value)
	}
	return memAddr, mode
}


// Indirect,Y
func addr_mode_IndirectY	(offset uint16) (uint16, string) {

	memAddr  := ( uint16(Memory[Memory[offset+1]])<<8 | uint16(Memory[Memory[offset]]) ) + uint16(Y)
	value := Memory[memAddr]
	mode		:= "Indirect,Y"

	if Debug {
		fmt.Printf("\t%s addressing mode.\tMemory[%04X]\t\tValue obtained: %02X\n", mode, memAddr, value)
	}
	return memAddr, mode
}

// Indirect,X
func addr_mode_IndirectX	(offset uint16) (uint16, string) {

	memAddr  := ( uint16(Memory[Memory[offset+1]])<<8 | uint16(Memory[Memory[offset]]) ) + uint16(X)
	value := Memory[memAddr]
	mode		:= "Indirect,X"

	if Debug {
		fmt.Printf("\t%s addressing mode.\tMemory[%04X]\t\tValue obtained: %02X\n", mode, memAddr, value)
	}
	return memAddr, mode
}


// ---------------------------- Library Function ---------------------------- //

// Memory Page Boundary cross detection
func MemPageBoundary(Address1, Address2 uint16) bool {

	var cross bool = false

	// Get the High byte only to compare
	// Page Boundary Cross detected
	if Address1 >> 8 != Address2 >> 8 {
		cross = true

		if Debug {
			fmt.Printf("\tMemory Page Boundary Cross detected! Add 1 cycle.\tPC High byte: %02X\tBranch High byte: %02X\n",Address1 >>8, Address2 >>8)
		}
	// NO Page Boundary Cross detected
	} else {
		if Debug {
			fmt.Printf("\tNo Memory Page Boundary Cross detected.\tPC High byte: %02X\tBranch High byte: %02X\n",Address1 >>8, Address2 >>8)
		}
	}

	return cross
}


// Decode Two's Complement
func DecodeTwoComplement(num byte) int8 {

	var sum int8 = 0

	for i := 0 ; i < 8 ; i++ {
		// Sum each bit and sum the value of the bit power of i (<<i)
		sum += (int8(num) >> i & 0x01) << i
	}

	return sum
}

// BCD - Binary Coded Decimal
func BCD(number byte) byte {

	var tmp_hundreds, tmp_tens, tmp_ones, bcd	byte

	// Split the Decimal Value
	tmp_hundreds = number / 100		// Hundreds
	tmp_tens = (number / 10)  % 10	// Tens
	tmp_ones = (number % 100) % 10	// Ones

	fmt.Printf("H: %d\tT: %d\tO: %d\n", tmp_hundreds, tmp_tens, tmp_ones)

	// Combine in one decimal number
	bcd = (tmp_hundreds * 100) + (tmp_tens * 10) + tmp_ones

	return bcd
}
