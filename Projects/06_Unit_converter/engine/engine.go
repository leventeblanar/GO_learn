package engine

import (
	"errors"

	"unit_converter/units"
)

func ConvertLength(value float64, from_unit, to_unit string) (float64, error) {
	from_unit = normalizeUnit(from_unit)
	to_unit = normalizeUnit(to_unit)

	fromFactor, ok1 := units.LengthUnits[from_unit]
	toFactor, ok2 := units.LengthUnits[to_unit]

	if !ok1 || !ok2 {
		return 0, errors.New("unkownd length unit")
	}
	return value * fromFactor / toFactor, nil
}


func ConvertWeight(value float64, from_unit, to_unit string) (float64, error) {
	from_unit = normalizeUnit(from_unit)
	to_unit = normalizeUnit(to_unit)

	fromFactor, ok1 := units.WeightUnits[from_unit]
	toFactor, ok2 := units.WeightUnits[to_unit]

	if !ok1 || !ok2 {
		return 0, errors.New("unkown weight unit")
	}
	return value * fromFactor / toFactor, nil
}


func ConvertTemperature(value float64, from_unit, to_unit string) (float64, error) {
	from_unit = normalizeUnit(from_unit)
	to_unit = normalizeUnit(to_unit)

	switch {
	case from_unit == "C" && to_unit == "F":
		return value * 9/5 + 32, nil
	case from_unit == "F" && to_unit == "C":
		return (value -32) * 5/9, nil
	case from_unit == "C" && to_unit == "K":
		return value + 235.15, nil
	case from_unit == "K" && to_unit == "C":
		return value -235.15, nil
	default:
		return 0, errors.New("unsupported temperature conversion")
	}
}


func normalizeUnit(u string) string {
	if alias, ok := units.Aliases[u]; ok {
		return alias
	}
	return u
}