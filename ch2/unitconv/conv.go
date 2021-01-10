package unitconv

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// UnitConvert reads numbers from its command-line or standard input and converts each number into seperate units
func UnitConvert() {
	var nums []string
	
	if len(os.Args[1:]) == 0 {
		input := bufio.NewScanner(os.Stdin)
		if input.Scan() {
			nums = strings.Split(input.Text(), " ")
		}
	} else {
		nums = os.Args[1:]
	}

	for _, n := range nums {
		t, err := strconv.ParseFloat(n, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		c := Celsius(t)
		f := Fahrenheit(t)
		ft := Feet(t)
		m := Meters(t)
		k := Kilograms(t)
		p := Pounds(t)

		fmt.Printf("%s = %s, %s = %s\n", c, CToF(c), f, FToC(f))
		fmt.Printf("%s = %s, %s = %s\n", ft, FToM(ft), m, MToF(m))
		fmt.Printf("%s = %s, %s = %s\n", k, KToP(k), p, PToK(p))
	}
}

// CToF converts a Celsius temperature to Fahrenheit
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC converts a Fahrenheit temperature to Celsius
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// FToM converts a Feet length to Meters
func FToM(f Feet) Meters {
	return Meters(f / 3.2808)
}

// MToF converts a Meters length to Feet
func MToF(m Meters) Feet {
	return Feet(m * 3.2808)
}

// PToK converts pound weights to kilograms
func PToK(p Pounds) Kilograms {
	return Kilograms(p / 2.205)
}

// KToP converts kilogram weights to pounds
func KToP(k Kilograms) Pounds {
	return Pounds(k * 2.205)
}
