package convert

import (
	"fmt"
	"math"
)

const (
	milesConst  = 1.60934
	degreeConst = 32
	bytesConst  = 1024
)

func MilesToKilometers(miles float64) float64 {
	return miles * milesConst
}

func KilometersToMiles(kilometers float64) float64 {
	return kilometers / milesConst
}

func CelsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + degreeConst
}

func FahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - degreeConst) * 5 / 9
}

func BytesToKilobytes(bytes int64) string {
	KB := float64(bytes) / bytesConst
	result := fmt.Sprintf("%g KB", math.Round(KB*100)/100)
	return result
}

func KilobytesToBytes(kilobytes float64) int64 {
	return int64(kilobytes * bytesConst)
}
