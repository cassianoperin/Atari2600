package VGS

import	"os"
import	"fmt"

//-------------------------------------------------- Processor Flags --------------------------------------------------//

// Zero Flag
func flags_Z(value byte) {
	if Debug {
		fmt.Printf("\tFlag Z: %d ->", P[1])
	}
	// Check if final value is 0
	if value == 0 {
		P[1] = 1
	} else {
		P[1] = 0
	}
	if Debug {
		fmt.Printf(" %d\n", P[1])
	}
}

// Negative Flag
func flags_N(value byte) {
	if Debug {
		fmt.Printf("\tFlag N: %d ->", P[7])
	}
	// Set Negtive flag to the the MSB of the value
	P[7] = value >> 7

	if Debug {
		fmt.Printf(" %d | Value = %08b\n", P[7], value)
	}
}

// Carry Flag
func flags_C(value1, value2 byte) {
	if Debug {
		fmt.Printf("\tFlag C: %d ->", P[0])
	}

	// Check if final value is 0
	if value1 >= value2 {
		P[0] = 1
	} else {
		P[0] = 0
	}

	if Debug {
		fmt.Printf(" %d\n", P[0])
	}
}

// Carry Flag for Subtractions (SBC and CMP)
func flags_C_Subtraction(value1, value2 byte) {
	if Debug {
		fmt.Printf("\tFlag C: %d ->", P[0])
	}

	// If the new value is bigger than the original clear the flag
	// if value1 <= value2 {
	if value1 < value2 {
		P[0] = 0
	} else {
		P[0] = 1
	}


	if Debug {
		fmt.Printf(" %d (SBC)\n" , P[0])
	}
}

// oVerflow Flag for ADC
// value1 = Accumulator, value2 = memory, value3 = carry flag
func Flags_V_ADC(value1, value2, value3 byte) {
	var (
		carry_OUT 		byte = 0	// Keep the value of the Carry OUT
		carry_6	 		byte = 0	// Keep the value of the Carry of bit 6 [address7 in the array]

		// First SUM
		sum_v1_v2			[8]byte	// Keep the result of the sum of value1 and value2
		carry_v1_v2		[8]byte	// Keep the carry array of sum of value1 and value2

		// Second SUM
		sum_1stsum_v3		[8]byte	// Keep the result of the first sum + value 3
		carry_1stsum_v3	[8]byte	// Keep the carry array of the first sum and value 3
	)

	if Debug {
		fmt.Printf("\tFlag V: %d ->", P[6])
	}

	// Make the magic
	for i:=0 ; i <= 7 ; i++ 	{
		// sum the bit from value one + bit from value 2 + carry value
		tmp := (value1 >> byte(i) & 0x01) + (value2 >> byte(i) & 0x01) + carry_v1_v2[i]
		if tmp >= 2 {
			// set the carry out
			if i == 7 {
				carry_OUT = 1
			} else {
				carry_v1_v2[i+1] = 1
			}
			// Sum
			if tmp == 2 {
				sum_v1_v2[i]=0
			} else if tmp == 3 {
				sum_v1_v2[i]=1
			} else {
				fmt.Printf("Flags_V_ADC unexpected result.\n")
				os.Exit(2)
			}
		} else {
			sum_v1_v2[i]=tmp
		}
	}
	
	// fmt.Printf("\n\n  ")
	//
	// for i := 7 ; i >= 0 ; i-- {
	// 	fmt.Printf("%d",carry_v1_v2[i])
	// }
	// fmt.Printf("\n")
	// fmt.Printf("\n  %08b\t%d",value1,value1)
	// fmt.Printf("\n  %08b\t%d",value2,value2)
	// fmt.Printf("\n  ---------\n  ")
	// for i := 7 ; i >= 0 ; i-- {
	// 	fmt.Printf("%d",sum_v1_v2[i])
	// }
	// fmt.Printf("\n")


	// Sum the first result with accumulator
	// Make the magic
	for i:=0 ; i <= 7 ; i++ 	{
		// sum the bit from value one + bit from value 2 + carry value
		tmp := (sum_v1_v2[i]) + (value3 >> byte(i) & 0x01) + carry_1stsum_v3[i]
		if tmp >= 2 {
			// set the carry out
			if i == 7 {
				carry_OUT = 1
			} else {
				carry_1stsum_v3[i+1] = 1
			}
			// Sum
			if tmp == 2 {
				sum_1stsum_v3[i]=0
			} else if tmp == 3 {
				sum_1stsum_v3[i]=1
			} else {
				fmt.Printf("Flags_V_ADC unexpected result.\n")
				os.Exit(2)
			}
		} else {
			sum_1stsum_v3[i]=tmp
		}
	}

	// fmt.Printf("\n\n  ")
	//
	// for i := 7 ; i >= 0 ; i-- {
	// 	fmt.Printf("%d",carry_1stsum_v3[i])
	// }
	// fmt.Printf("\n\n  ")
	// for i := 7 ; i >= 0 ; i-- {
	// 	fmt.Printf("%d",sum_v1_v2[i])
	// }
	// // fmt.Printf("\n")
	// fmt.Printf("\n  %08b\t%d",value3,value3)
	// fmt.Printf("\n  ---------\n  ")
	// for i := 7 ; i >= 0 ; i-- {
	// 	fmt.Printf("%d",sum_1stsum_v3[i])
	// }
	// fmt.Printf("\n\n\n")


	// Calculate the last bit of carry based on two sums
	if carry_v1_v2[7] == 1 || carry_1stsum_v3[7] == 1 {
		carry_6 = 1
	}

	// Formula to calculate: V = C6 xor C7
	P[6] = carry_6 ^ carry_OUT
	// fmt.Printf("\nV: %d", P[6])

	if Debug {
		fmt.Printf(" %d\n", P[6])
	}
}


// func Flags_V_ADC(value1, value2 byte) {
// 	var (
// 		carry_bit		[8]byte
// 		carry_OUT 	byte = 0
// 	)
//
// 	if Debug {
// 		fmt.Printf("\tFlag V: %d ->", P[6])
// 	}
//
// 	// Make the magic
// 	for i:=0 ; i <= 7 ; i++ 	{
// 		// sum the bit from value one + bit from value 2 + carry value
// 		tmp := (value1 >> byte(i) & 0x01) + (value2 >> byte(i) & 0x01) + carry_bit[i]
// 		if tmp >= 2 {
// 			// set the carry out
// 			if i == 7 {
// 				carry_OUT = 1
// 			} else {
// 				carry_bit[i+1] = 1
// 			}
// 		}
// 	}
// 	for i : = 7 ; i >= 0 ; i-- {
// 		fmt.Printf("%d",carry_bit[i])
// 	}
// 	fmt.Printf("\n\n")
// 	fmt.Printf("\n  %08b\t%d",value1,value1)
// 	fmt.Printf("\n  %08b\t%d",value2,value2)
//
// 	// Formula to calculate: V = C6 xor C7
// 	P[6] = carry_bit[7] ^ carry_OUT
// 	// fmt.Printf("\nV: %d", P[6])
//
// 	if Debug {
// 		fmt.Printf(" %d\n", P[6])
// 	}
// }

// oVerflow Flag for SBC
func Flags_V_SBC(value1, value2 byte) {
	var (
		carry_bit		= [8]byte{}
		carry_OUT 	byte = 0
	)

	// fmt.Printf("\n  \t %d\t(carry IN)",P[0])
	// fmt.Printf("\n  %08b\tDecimal: %d",value1,value1)
	// Since internall subtraction is just addition of the ones-complement
	// N can simply be replaced by 255-N in the formulas of Flags_V_ADC
	value2 = 255 - value2
	// fmt.Printf("\n  %08b\tDecimal: %d",value2,value2)

	// Set the carry flag on bit 0 of carry_bit Array to bring the carry if exists
	carry_bit[0] = P[0]

	if Debug {
		fmt.Printf("\tFlag V: %d ->", P[6])
	}

	// Make the magic
	for i:=0 ; i <= 7 ; i++{
		// sum the bit from value one + bit from value 2 + carry value
		tmp := (value1 >> byte(i) & 0x01) + (value2 >> byte(i) & 0x01) + carry_bit[i]
		if tmp >= 2 {
			// set the carry out
			if i == 7 {
				carry_OUT = 1
			} else {
				carry_bit[i+1] = 1
			}
		}
	}

	// fmt.Printf("\n  %08b\tDecimal: %d\t(SUM)",value1+value2+carry_bit[0],value1+value2+carry_bit[0])
	//
	// fmt.Printf("\n\n%d ",carry_OUT)
	// for i:=7 ; i >=0 ; i--{
	// 	fmt.Printf("%d",carry_bit[i])
	// }
	// fmt.Printf("\t(carry OUT | carry array)")

	// Formula to calculate: V = C6 xor C7 (if they are different is a overflow)
	P[6] = carry_bit[7] ^ carry_OUT
	// fmt.Printf("\nV: %d", P[6])

	if Debug {
		fmt.Printf(" %d\n", P[6])
	}

}
