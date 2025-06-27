package converters

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var weightUnits = map[string]float64{
	"mg":   0.000001,
	"ct":   0.0002,
	"g":    0.001,
	"oz":   0.0283495,
	"lb":   0.453592,
	"kg":   1,
	"pood": 16.3807,
	"q":    100,
	"t":    1000,
}

func ConvertWeight(valueStr, from, to string) (string, error) {
	valueStr = strings.TrimSpace(valueStr)
	from = strings.ToLower(strings.TrimSpace(from))
	to = strings.ToLower(strings.TrimSpace(to))

	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return "", errors.New("invalid number")
	}

	fromFactor, ok := weightUnits[from]
	if !ok {
		return "", errors.New("unknown from-unit: " + from)
	}

	toFactor, ok := weightUnits[to]
	if !ok {
		return "", errors.New("unknown to-unit: " + from)
	}

	result := value * fromFactor / toFactor

	formatted := fmt.Sprintf("%f", result)
	trimmed := strings.TrimRight(strings.TrimRight(formatted, "0"), ".")

	return trimmed, nil
}
