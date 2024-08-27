package conv

import (
	"fmt"
	"math"
)

func ConverterLength(convertFrom, convertTo string, value float64) string {
	val1, val2 := 1.0, 1.0

	switch convertFrom {
	case "km":
		val1 *= 10e4

	case "m":
		val1 *= 10e3

	case "cm":
		val1 *= 10
	}

	switch convertTo {
	case "km":
		val2 *= 10e4

	case "m":
		val2 *= 10e3

	case "cm":
		val2 *= 10
	}
	res := val1/val2 * value
	factor := math.Pow(10, 6) // 10^6 = 1000000
    truncated := math.Floor(res*factor) / factor
	return fmt.Sprintf("%.6f", truncated)
}