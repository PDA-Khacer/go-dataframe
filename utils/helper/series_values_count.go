package helper

import (
	"github.com/PDA-Khacer/go-dataframe/common"
	"github.com/PDA-Khacer/go-dataframe/series"
	"github.com/PDA-Khacer/go-dataframe/utils/converter"
)

func SeriesValuesCount[S common.Frame, D int](source *series.Series[S]) *series.Series[int] {
	result := &series.Series[int]{
		Name:    source.Name,
		Indexes: []string{},
		Values:  []*int{},
		DType:   "",
	}

	mapCounter := map[S]int{}
	counterNil := 0
	for _, v := range source.Values {
		if v == nil {
			counterNil += 1
			continue
		}
		if count, ok := mapCounter[*v]; ok {
			mapCounter[*v] = count + 1
		} else {
			mapCounter[*v] = 1
		}
	}
	// update index and value
	for key, val := range mapCounter {
		tempInx := converter.ConvertGenericsToString(&key)
		if tempInx == nil {
			counterNil += 1
			continue
		}
		result.Indexes = append(result.Indexes, *tempInx)
		result.Values = append(result.Values, &val)
	}

	if counterNil > 0 {
		result.Indexes = append(result.Indexes, "nil")
		result.Values = append(result.Values, &counterNil)
	}
	return result
}
