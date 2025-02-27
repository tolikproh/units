package units

import (
	"encoding/json"
	"errors"
)

// UnitType определяет тип единицы измерения
type UnitType int

// Константы для различных типов единиц измерения
const (
	NilType UnitType = iota
	LengthType
	WeigthType
	ThingsType
	SetType
	PackageType
	BayType
)

func unmarshalJSON(data []byte) (UnitType, error) {
	var qj quantityJSON
	if err := json.Unmarshal(data, &qj); err != nil {
		return NilType, err
	}

	return qj.Type, nil

}

func NewJSON(data []byte) (Quantiter, error) {
	types, err := unmarshalJSON(data)
	if err != nil {
		return nil, err
	}

	switch types {
	case LengthType:
		return NewLengthJSON(data)
	case WeigthType:
		return NewWeigthJSON(data)
	case ThingsType:
		return NewThingsJSON(data)
	case SetType:
		return NewSetJSON(data)
	case PackageType:
		return NewPackageJSON(data)
	case BayType:
		return NewBayJSON(data)
	}

	return nil, errors.New("json is not stuct correct")
}
