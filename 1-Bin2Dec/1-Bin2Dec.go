package main

import (
	"fmt"
	"github.com/thoas/go-funk"
	"math"
	"os"
)

func main() {
	// check if argument given
	if len(os.Args) < 2 {
		fmt.Println("To few argument")
		return
	}

	// get first argument/binary digits
	bin := os.Args[1]

	// reverse binary digits
	binReverse := funk.ReverseString(bin)

	// declare decimal representation
	var dec float64

	// iterate over reversed binary digits
	for pos, char := range binReverse {
		// check if digit '0' or '1'
		if char != '0' && char != '1' {
			fmt.Printf("Invalid binary digit: %c\n", char)
			return
		}

		// ignore '0'
		if char == '0' {
			continue
		}

		dec += math.Exp2(float64(pos))
	}

	fmt.Printf("%s = %f\n", bin, dec)
}
