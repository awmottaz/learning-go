// Package tempconv performs temperature conversions between Fahrenheit,
// Celcius, and Kelvin.
package tempconv

import (
	"fmt"
)

// Fahrenheit temperature units
type Fahrenheit float64

// Celsius temperature units
type Celsius float64

// Kelvin temperature units
type Kelvin float64

const (
	// AbsoluteZeroC equals the absolute zero temperature in Celsius
	AbsoluteZeroC Celsius = -273.15
	// FreezingC equals the freezing point of water in Celsius
	FreezingC Celsius = 0
	// BoilingC equals the boiling point of water in Celsius
	BoilingC Celsius = 0
)

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%gK", k)
}
