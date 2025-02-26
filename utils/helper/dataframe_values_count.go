package helper

import (
	"dataframe/common"
	"dataframe/dataframes"
	"dataframe/series"
	"github.com/samber/lo"
)

func DataframeValuesCount[S common.Frame, D int](s *dataframes.DataFrame[S]) *dataframes.DataFrame[int] {
	// case using Series
	var lstNewSer []*series.Series[int]
	/*
		mapColSeries: example:
		{
			"http_code": {"500": 10, "400": 4},
			"status": {"success" 13, "fail": 1},
			"data": {"500": 9, "success" 3}
		}
	*/
	mapColSeries := map[string]map[string]*int{}
	for _, ser := range s.Series {
		// get SeriesValuesCount each series
		newSer, err := series.Apply(ser, SeriesValuesCount[S, int])
		if err != nil {
			return nil
		}
		lstNewSer = append(lstNewSer, newSer)
		mapColSeries[ser.Name] = newSer.GetMapData()
	}
	/*
		Start combine mapColSeries because http_code having 500, 400 but status having success, fail
		But data having 500 & success
		expect:
			     http_code |  status | data
		500         10     |   nil   |  9
		400         4      |   nil   |  nil
		success     nil    |   13    |  3
		fail        nil    |   1     |  nil
	*/
	cols := lo.Keys(mapColSeries) //  http_code , status , data

	// fill all Indexes to Set
	setIndexes := map[string]bool{}
	for _, serCounter := range lstNewSer {
		if serCounter == nil {
			continue
		}
		for _, idx := range serCounter.Indexes {
			setIndexes[idx] = true
		}
	}

	indexes := lo.Keys(setIndexes) // 500, 400, success, fail

	// fill data
	var data [][]*int
	for _, idx := range indexes {
		var row []*int
		for _, col := range cols {
			if v, ok := mapColSeries[col][idx]; ok {
				row = append(row, v)
			} else {
				row = append(row, nil)
			}
		}
		data = append(data, row)
	}

	resultDf, err := dataframes.NewDataframe(indexes, cols, data)
	if err != nil {
		return nil
	}

	return resultDf

}
