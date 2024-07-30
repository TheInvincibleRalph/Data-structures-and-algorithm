package main

import "fmt"

func main() {
	// maxInt := int(^uint(0) >> 1)
	// fmt.Println(maxInt)

	// // or

	// maxIn := uint(0)
	// fmt.Println(maxIn)

	// maxIng := ^maxIn //produces a number with all bits set to 1 in binary representation (2^64 - 1 in decimal notation) which is the maximum value that can be stored in an unsigned integer of 64 bits.
	// fmt.Println(maxIng)

	// maxIntg := maxIng >> 1 //shifts all bits one position to the right, which results in a number with all bits set to 1 except the leftmost one which is set to 0 (2^63 - 1 in decimal notation) which is the maximum value that can be stored in a signed integer of 64 bits.
	// fmt.Println(maxIntg)

	// maxTing := int(maxIntg) //the value remains the same, but the type is changed from uint to int.
	// fmt.Println(maxTing)

	// Further Explanation:

	//  Signed intergers use one bit (usually the leftmost one) to represent
	// the sign of the number (0 for positive and 1 for negative),
	// so the maximum value that can be stored in a signed integer of 64 bits
	// is half of the maximum value that can be stored in an unsigned integer of 64 bits.
	// So in maxIntg above, the 0 bit is used to represent the sign of the number,
	// and the remaining 63 bits are used to represent the value of the number.

	//  Therefore the maximum value that can be stored in an unsigned integer of 64 bits is 2^64 - 1,
	// and that in a signed integer of 64 bits is 2^63 - 1 because of the leftmost 0 in signed integers used for sign

	res := factorial(100)
	fmt.Println(res)
}

func factorial(num uint64) uint64 {
	if num == 0 || num == 1 {
		return 1
	} else {
		return num * factorial(num-1)
	}
}
