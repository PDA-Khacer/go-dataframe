package converter

import (
	"fmt"
	"github.com/PDA-Khacer/go-dataframe/common"
	"strconv"
)

func ConvertGenerics[S common.Frame, D common.Frame](s *S) *D {
	var re any // pointer
	var casted D
	typeDes := fmt.Sprintf("%T", *new(D))
	switch typeDes {
	case "int":
		temp := ConvertGenericsToInt(s)
		if temp != nil {
			re = *temp
		} else {
			return nil
		}
	case "string":
		temp := ConvertGenericsToString(s)
		if temp != nil {
			re = *temp
		} else {
			return nil
		}
	case "float32":
		temp := ConvertGenericsToFloat32(s)
		if temp != nil {
			re = *temp
		} else {
			return nil
		}
	case "float64":
		temp := ConvertGenericsToFloat64(s)
		if temp != nil {
			re = *temp
		} else {
			return nil
		}
	default:
		return nil
	}
	casted = re.(D)
	return &casted
}

func ConvertGenericsToInt[T any](t *T) *int {
	if t == nil {
		return nil
	}
	switch t := any(*t).(type) {
	// handle all non-approximate type cases first
	case int:
		return &t // t is int
	case string:
		i, err := strconv.Atoi(t)
		if err != nil {
			return nil
		}
		return &i
	case float32:
		i := int(t)
		return &i
	case float64:
		i := int(t)
		return &i
	default:
		return nil
	}
}

func ConvertGenericsToString[T any](t *T) *string {
	if t == nil {
		return nil
	}
	switch t := any(*t).(type) {
	// handle all non-approximate type cases first
	case int:
		s := strconv.Itoa(t)
		return &s // t is int
	case int64:
		s := strconv.FormatInt(t, 10)
		return &s
	case string:
		return &t
	case float32, float64:
		s := fmt.Sprintf("%f", t)
		return &s
	case bool:
		s := fmt.Sprintf("%t", t)
		return &s
	default:
		return nil
	}
}

func ConvertGenericsToFloat64[T any](t *T) *float64 {
	if t == nil {
		return nil
	}
	switch t := any(*t).(type) {
	// handle all non-approximate type cases first
	case int:
		f := float64(t)
		return &f // t is int
	case string:
		f, err := strconv.ParseFloat(t, 64)
		if err != nil {
			return nil
		}
		return &f
	case float64:
		return &t
	case float32:
		f := float64(t)
		return &f
	default:
		return nil
	}
}

func ConvertGenericsToFloat32[T any](t *T) *float32 {
	if t == nil {
		return nil
	}
	switch t := any(*t).(type) {
	// handle all non-approximate type cases first
	case int:
		f := float32(t)
		return &f // t is int
	case string:
		f, err := strconv.ParseFloat(t, 32)
		if err != nil {
			return nil
		}
		f2 := float32(f)
		return &f2
	case float64:
		f := float32(t)
		return &f
	case float32:
		return &t
	default:
		return nil
	}
}
