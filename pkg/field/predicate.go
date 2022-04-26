package field

import (
	"fmt"
	"strings"
)

type PredicateType int

const (
	PredicateTypeEq = iota
	PredicateTypeContains
	PredicateTypeGt
	PredicateTypeGtE
	PredicateTypeLt
	PredicateTypeLtE
	PredicateTypeIn
)

var (
	presMap = map[PredicateType]string{
		PredicateTypeEq:       "Eq",
		PredicateTypeContains: "Contains",
		PredicateTypeGt:       "Gt",
		PredicateTypeGtE:      "GTE",
		PredicateTypeLt:       "Lt",
		PredicateTypeLtE:      "LTE",
		PredicateTypeIn:       "In",
	}
	presEntMap = map[PredicateType]string{
		PredicateTypeEq:       "EQ",
		PredicateTypeContains: "Contains",
		PredicateTypeGt:       "GT",
		PredicateTypeGtE:      "GTE",
		PredicateTypeLt:       "LT",
		PredicateTypeLtE:      "LTE",
		PredicateTypeIn:       "In",
	}
	StrToPreMap = map[string]PredicateType{
		"eq":       PredicateTypeEq,
		"contains": PredicateTypeContains,
		"gt":       PredicateTypeGt,
		"gte":      PredicateTypeGtE,
		"lt":       PredicateTypeLt,
		"lte":      PredicateTypeLtE,
		"in":       PredicateTypeIn,
	}
)

func (pred PredicateType) String() string {
	return presMap[pred]
}

func (pred PredicateType) EntString() string {
	return presEntMap[pred]
}

func NewPredicateType(s string) PredicateType {
	predicateType, ok := StrToPreMap[strings.ToLower(s)]
	if !ok {
		panic(fmt.Sprintf("unknown PredicateType: %s", s))
	}
	return predicateType
}

type Predicate struct {
	Name      string
	Type      PredicateType
	FieldType string
	EntName   string
}
