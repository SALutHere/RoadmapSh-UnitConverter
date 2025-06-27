package converters

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var fromC = map[string]func(float64) float64{
	"r": func(c float64) float64 { return c * 0.8 },
	"f": func(c float64) float64 { return c*9/5 + 32 },
	"c": func(c float64) float64 { return c },
	"k": func(c float64) float64 { return c + 273.15 },
}

var toC = map[string]func(float64) float64{
	"r": func(v float64) float64 { return v * 5 / 4 },
	"f": func(v float64) float64 { return 5.0 / 9.0 * (v - 32) },
	"c": func(v float64) float64 { return v },
	"k": func(v float64) float64 { return v - 273.15 },
}

func ConvertTemperature(valueStr, from, to string) (string, error) {
	valueStr = strings.TrimSpace(valueStr)
	from = strings.ToLower(strings.TrimSpace(from))
	to = strings.ToLower(strings.TrimSpace(to))

	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return "", errors.New("invalid number")
	}

	if _, ok := toC[from]; !ok {
		return "", errors.New("unknown from-unit: " + from)
	}

	if _, ok := fromC[to]; !ok {
		return "", errors.New("unknown to-unit: " + from)
	}

	result := fromC[to](toC[from](value))

	formatted := fmt.Sprintf("%f", result)
	trimmed := strings.TrimRight(strings.TrimRight(formatted, "0"), ".")

	return trimmed, nil
}
