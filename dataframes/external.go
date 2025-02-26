package dataframes

import (
	"dataframe/common"
	"dataframe/series"
	"dataframe/utils/converter"
	"errors"
	"github.com/samber/lo"
)

func AsType[S common.Frame, D common.Frame](sourceDf *DataFrame[S]) (*DataFrame[D], error) {
	if sourceDf == nil {
		return nil, errors.New("sourceDf must not be nil")
	}

	var seriesDf []*series.Series[D]

	seriesDf = lo.Map(sourceDf.Series, func(s *series.Series[S], index int) *series.Series[D] {
		valSer := lo.Map(s.Values, func(val *S, index int) *D {
			return converter.ConvertGenerics[S, D](val)
		})
		result := &series.Series[D]{
			Name:    s.Name,
			Indexes: s.Indexes,
			Values:  valSer,
			DType:   s.DType,
		}
		result.UpdateDType()
		return result
	})

	var dfValues [][]*D
	if sourceDf.Values != nil && len(sourceDf.Values) > 0 {
		dfValues = lo.Map(sourceDf.Values, func(row []*S, index int) []*D {
			return lo.Map(row, func(val *S, index int) *D {
				return converter.ConvertGenerics[S, D](val)
			})
		})
	}

	return &DataFrame[D]{
		Indexes: sourceDf.Indexes,
		Columns: sourceDf.Columns,
		DType:   sourceDf.DType,
		Values:  dfValues,
		Series:  seriesDf,
	}, nil
}

func Apply[S common.Frame, D common.Frame](s *DataFrame[S], fn func(*DataFrame[S]) *DataFrame[D]) (*DataFrame[D], error) {
	if s == nil {
		return nil, errors.New("source dataframe is nil")
	}

	re := fn(s)

	return re, nil
}
