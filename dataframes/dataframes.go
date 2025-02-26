package dataframes

import (
	"dataframe/common"
	"dataframe/series"
	"dataframe/utils"
	"dataframe/utils/converter"
	"errors"
	"github.com/samber/lo"
	"strconv"
	"strings"
)

type DataFrame[T common.Frame] struct {
	Indexes []string
	Columns []string
	Values  [][]*T
	Series  []*series.Series[T]
	DType   string
}

func _implement[T common.Frame]() IDataFrame[T] {
	return &DataFrame[T]{
		Indexes: nil,
		Columns: nil,
		Values:  nil,
		Series:  nil,
		DType:   "",
	}
}

type IDataFrame[T common.Frame] interface {
	Apply(func(*T) *T) *DataFrame[T]
	GetSeries(string) *series.Series[T]
	Agg([]string) (*DataFrame[float64], error)
	Drop([]string) *DataFrame[T]
	//Array() []T
	//T() Series[any]
}

func (df *DataFrame[T]) Drop(colName []string) *DataFrame[T] {
	if df.Series != nil {
		df.Series = lo.Filter(df.Series, func(item *series.Series[T], _ int) bool {
			if item != nil && lo.Contains(colName, item.Name) {
				return false
			}
			return true
		})
	}

	if df.Values != nil {
		lo.ForEach(colName, func(item string, index int) {
			idx := lo.IndexOf(df.Columns, item)
			if idx != -1 {
				df.Values = lo.Map(df.Values, func(row []*T, _ int) []*T {
					return lo.Filter(row, func(_ *T, index int) bool {
						if index == idx {
							return false
						}
						return true
					})
				})
			}
		})
	}

	// remove col
	df.Columns = lo.Filter(df.Columns, func(item string, _ int) bool {
		return !lo.Contains(colName, item)
	})

	return df
}

func (df *DataFrame[T]) Apply(f func(*T) *T) *DataFrame[T] {
	if df.Series != nil {
		df.Series = lo.Map(df.Series, func(item *series.Series[T], _ int) *series.Series[T] {
			item.Apply(f)
			return item
		})
	}

	if df.Values != nil {
		df.Values = lo.Map(df.Values, func(row []*T, _ int) []*T {
			return lo.Map(row, func(item *T, _ int) *T {
				return f(item)
			})
		})
	}

	return df
}

/*
Agg Method support only: Max, Min, Sum, Avg
*/
func (df *DataFrame[T]) Agg(methods []string) (*DataFrame[float64], error) {
	index := methods
	// loop all item
	var seriesDf []*series.Series[float64]
	if df.Series != nil {
		seriesDf = lo.Map(df.Series, func(item *series.Series[T], _ int) *series.Series[float64] {
			if item != nil && len(item.Values) > 0 {
				maxVal := converter.ConvertGenericsToInt(item.Values[0])
				minVal := converter.ConvertGenericsToInt(item.Values[0])
				sumVal := converter.ConvertGenericsToInt(item.Values[0])

				for i := 1; i < len(item.Values); i++ {
					maxVal = utils.MaxPointer(maxVal, converter.ConvertGenericsToInt(item.Values[i]))
					minVal = utils.MinPointer(minVal, converter.ConvertGenericsToInt(item.Values[i]))
					sumVal = utils.SumPointer(sumVal, converter.ConvertGenericsToInt(item.Values[i]))
				}
				// save val to series
				seriesMapData := map[string]float64{}
				for _, s := range index {
					if strings.ToLower(s) == "sum" {
						seriesMapData["sum"] = float64(*sumVal)
					}

					if strings.ToLower(s) == "min" {
						seriesMapData["min"] = float64(*minVal)
					}

					if strings.ToLower(s) == "max" {
						seriesMapData["max"] = float64(*maxVal)
					}

					if strings.ToLower(s) == "mean" {
						seriesMapData["mean"] = float64(*sumVal / len(item.Values))
					}
				}

				newSeries, err := series.NewSeries(item.Name, index, seriesMapData)
				if err != nil {
					return nil
				}
				return newSeries
			}
			return nil
		})
		return &DataFrame[float64]{
			Indexes: index,
			//Columns: nil,
			//Values:  nil,
			Series: seriesDf,
			DType:  "",
		}, nil
	}

	if df.Values != nil {
		// TODO
	}

	return nil, nil
}

func (df *DataFrame[T]) GetSeries(Name string) *series.Series[T] {
	a, _ := lo.Find(df.Series, func(item *series.Series[T]) bool {
		if item != nil && item.Name == Name {
			return true
		}
		return false
	})
	return a
}

func NewDataframe[T common.Frame](
	indexes []string,
	columns []string,
	data [][]*T,
) (*DataFrame[T], error) {
	df := &DataFrame[T]{
		Indexes: indexes,
		Columns: columns,
		Values:  data,
		Series:  []*series.Series[T]{},
		DType:   "",
	}
	// check len col & data size
	if len(columns) > 0 && len(data) > 0 && len(columns) != len(data[0]) {
		return nil, errors.New("columns and data must have the same length")
	}

	// create data for Series
	for i := 0; i < len(columns); i++ {
		ser, err := series.NewSeriesWithList(columns[i], indexes, utils.GetColValuesOf2DSlice[T](data, i))
		if err != nil {
			return nil, err
		}
		df.Series = append(df.Series, ser)
	}

	return df, nil
}

/*
NewDataframeWithRowMap create data frame with map Row data

	data := map[string][]int{
		"row1": {1,2,3},
		"row2": {4,5,6},
		"row3": {7,8,9}
	}
*/
func NewDataframeWithRowMap[T common.Frame](
	indexes []string,
	columns []string,
	data map[string][]*T,
) (*DataFrame[T], error) {
	df := &DataFrame[T]{
		Indexes: indexes,
		Columns: columns,
		Values:  [][]*T{},
		Series:  []*series.Series[T]{},
		DType:   "",
	}

	if len(columns) == 0 {
		for i := 0; i < utils.GetNumberColOfMapRow(data); i++ {
			columns = append(columns, strconv.Itoa(i))
		}
	}
	var err error
	for i := 0; i < len(columns); i++ {
		var ser *series.Series[T]
		if len(indexes) > 0 {
			colData, err := utils.GetColValuesOf2DMapRowAndIndex[T](data, i, indexes)
			if err != nil {
				return nil, err
			}
			ser, err = series.NewSeriesWithList(columns[i], indexes, colData)
			if err != nil {
				return nil, err
			}
		} else {
			colData, index, err := utils.GetColValuesOf2DMapRow[T](data, i)
			if err != nil {
				return nil, err
			}
			ser, err = series.NewSeriesWithList(columns[i], index, colData)
			if err != nil {
				return nil, err
			}
			df.Indexes = index
		}
		df.Series = append(df.Series, ser)
	}
	return df, err
}

//func BuildDataframeFromSeries[T common.Frame]([]*series.Series[T]) *DataFrame[T] {
//
//}
