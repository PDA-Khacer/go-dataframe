package io

import (
	"github.com/PDA-Khacer/go-dataframe/dataframes"
	"github.com/PDA-Khacer/go-dataframe/series"
	"github.com/PDA-Khacer/go-dataframe/vendor_lib/json2jsonschema/parse"
	"github.com/samber/lo"
	"strconv"
)

//func DataFrameJsonNormalize(df dataframes.DataFrame[string]) dataframes.DataFrame[string] {
//	// loop item if item can convert to json
//	return
//}

func SeriesJsonNormalize(seriesData series.Series[string]) (*dataframes.DataFrame[string], error) {
	var indexDf []string // skip that no need, add option execute that
	var colDf []string
	mapColDf := map[string]bool{}
	mapFlatMap := map[int]map[string]string{}
	dataDf := map[string][]*string{}
	lo.ForEach(seriesData.Values, func(item *string, index int) {
		// convert to flat field
		if item != nil {
			flatMap, err := parse.JsonStringToFlatMap(*item)
			if err != nil {
				return
			}
			for k, _ := range flatMap {
				mapColDf[k] = true
			}
			mapFlatMap[index] = flatMap
		}
	})
	colDf = lo.Keys(mapColDf)

	for i, flatMap := range mapFlatMap {
		var row []*string
		indexDf = append(indexDf, strconv.Itoa(i))
		// loop each col
		for _, c := range colDf {
			if v, ok := flatMap[c]; ok {
				row = append(row, &v)
			} else {
				row = append(row, nil)
			}
		}
		dataDf[strconv.Itoa(i)] = row
	}

	return dataframes.NewDataframeWithRowMap[string](indexDf, colDf, dataDf)
}
