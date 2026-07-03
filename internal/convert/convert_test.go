package convert

import (
	"math"
	"testing"
)

const epsilon = 1e-9

func TestConverter(t *testing.T) {
	tests := []struct {
		name   string
		input  float64
		want   float64
		method func(float64) float64
	}{
		{
			name:   "miles to kilometers",
			input:  1,
			want:   1.60934,
			method: MilesToKilometers,
		},
		{
			name:   "kilometers to miles",
			input:  1.60934,
			want:   1,
			method: KilometersToMiles,
		},
		{
			name:   "celsius to fahrenheit",
			input:  0,
			want:   32,
			method: CelsiusToFahrenheit,
		},
		{
			name:   "fahrenheit to celsius",
			input:  32,
			want:   0,
			method: FahrenheitToCelsius,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.method(tt.input); math.Abs(got-tt.want) > epsilon {
				t.Errorf("Converter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConverterBytes(t *testing.T) {
	tests := []struct {
		name   string
		input  int64
		want   string
		method func(int64) string
	}{
		{
			name:   "bytes to kilobytes",
			input:  1024,
			want:   "1 KB",
			method: BytesToKilobytes,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.method(tt.input); got != tt.want {
				t.Errorf("Converter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConverterKiloBytes(t *testing.T) {
	tests := []struct {
		name   string
		input  float64
		want   int64
		method func(float64) int64
	}{
		{
			name:   "kilobytes to bytes",
			input:  1,
			want:   1024,
			method: KilobytesToBytes,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.method(tt.input); got != tt.want {
				t.Errorf("Converter() = %v, want %v", got, tt.want)
			}
		})
	}
}
