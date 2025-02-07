package utils

import "math"

var ScreenWidth float64 = 1080
var ScreenHeight float64 = 2217
var Density float64 = 2.625

var CellWidth float64

var CellHeight float64

func DP(value float64) float64 {
	if value == 0 {
		return 0
	}
	return math.Ceil(Density * value)
}
