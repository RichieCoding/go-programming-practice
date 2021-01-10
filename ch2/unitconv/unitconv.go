// Package unitconv takes
package unitconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Feet float64
type Meters float64
type Pounds float64
type Kilograms float64

func (c Celsius) String() string {
	return fmt.Sprintf("%gC", c)
}
func (f Fahrenheit) String() string {
	return fmt.Sprintf("%gF", f)
}
func (f Feet) String() string {
	return fmt.Sprintf("%gFeet", f)
}
func (m Meters) String() string {
	return fmt.Sprintf("%gMeters", m)
}
func (p Pounds) String() string {
	return fmt.Sprintf("%gPounds", p)
}
func (k Kilograms) String() string {
	return fmt.Sprintf("%gKilograms", k)
}

// write a general-purpose unit conversion program analogous to cf that reads numbers from its command-line arguments or from the standard input if there are no aruments, and converts each number into units like temperature in Celsius and Fahrenheit, length in feet and meters, weight in pounds and kilograms, and the like
