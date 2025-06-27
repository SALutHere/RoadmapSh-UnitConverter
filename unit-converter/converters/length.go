package converters

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var lengthUnits = map[string]float64{
	"mm": 0.001,
	"cm": 0.01,
	"m":  1,
	"km": 1000,
	"in": 0.0254,
	"ft": 0.3048,
	"yd": 0.9144,
	"mi": 1609.34,
}

func ConvertLength(valueStr, from, to string) (string, error) {
	valueStr = strings.TrimSpace(valueStr)
	from = strings.ToLower(strings.TrimSpace(from))
	to = strings.ToLower(strings.TrimSpace(to))

	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return "", errors.New("invalid number")
	}

	fromFactor, ok := lengthUnits[from]
	if !ok {
		return "", errors.New("unknown from-unit: " + from)
	}

	toFactor, ok := lengthUnits[to]
	if !ok {
		return "", errors.New("unknown to-unit: " + to)
	}

	result := value * fromFactor / toFactor

	formatted := fmt.Sprintf("%f", result)
	trimmed := strings.TrimRight(strings.TrimRight(formatted, "0"), ".")

	return trimmed, nil
}
