package series

import (
	"dataframe/common"
	"errors"
	"fmt"
	"github.com/samber/lo"
)

type Series[T common.Frame] struct {
	Name    string
	Indexes []string
	Values  []*T
	DType   string
}

type ISeries[T common.Frame] interface {
	Apply(func(*T) *T)
	UpdateDType()
	GetMapData() map[string]*T
	//Array() []T
	//T() Series[any]
}

func (s *Series[T]) Apply(f func(*T) *T) {
	s.Values = lo.Map(s.Values, func(item *T, index int) *T {
		return f(item)
	})
}

func (s *Series[T]) UpdateDType() {
	s.DType = fmt.Sprintf("%T", *new(T))
}

func (s *Series[T]) GetMapData() map[string]*T {
	result := map[string]*T{}

	for idx, value := range s.Values {
		result[s.Indexes[idx]] = value
	}
	return result
}

/*
NewSeries input:

	data := map[string]int{
	"a": 1,
	"b": 2,
	"c": 3,
	}
*/
func NewSeries[T common.Frame](
	name string,
	indexes []string,
	values map[string]T,
) (*Series[T], error) {
	s := &Series[T]{
		Name:    name,
		Indexes: indexes,
		Values:  make([]*T, 0),
		DType:   "",
	}
	s.DType = fmt.Sprintf("%T", *new(T))
	switch s.DType {
	case "int":
	case "int64":
	case "string":
	case "bool":
	case "float32":
	case "float64":
	default:
		return nil, errors.New(fmt.Sprintf("unknown type: %s", s.DType))
	}

	if indexes == nil || len(indexes) == 0 {
		for k, v := range values {
			indexes = append(indexes, k)
			s.Values = append(s.Values, &v)
		}
	} else {
		lo.ForEach(indexes, func(indexName string, _ int) {
			if val, oke := values[indexName]; oke {
				s.Values = append(s.Values, &val)
			} else {
				s.Values = append(s.Values, nil)
			}
		})
	}
	return s, nil
}

func NewSeriesWithList[T common.Frame](
	name string,
	indexes []string,
	valuesList []*T,
) (*Series[T], error) {
	s := &Series[T]{
		Name:    name,
		Indexes: indexes,
		Values:  make([]*T, 0),
		DType:   "",
	}

	s.DType = fmt.Sprintf("%T", *new(T))
	switch s.DType {
	case "int":
	case "int64":
	case "string":
	case "bool":
	case "float32":
	case "float64":
	default:
		return nil, errors.New(fmt.Sprintf("unknown type: %s", s.DType))
	}

	if valuesList != nil || len(valuesList) != 0 {
		s.Values = append(s.Values, valuesList...)
	}

	return s, nil
}

func _implement[T common.Frame]() ISeries[T] {
	return &Series[T]{
		Indexes: nil,
		Values:  nil,
		DType:   "",
	}
}
