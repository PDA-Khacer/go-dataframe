package helper

import (
	"github.com/PDA-Khacer/go-dataframe/dataframes"
	"github.com/PDA-Khacer/go-dataframe/series"
	"github.com/PDA-Khacer/go-dataframe/utils"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSeriesValuesCount(t *testing.T) {
	Convey("Test SeriesValuesCount function", t, func() {
		data := []int{
			500, 400, 500, 401, 404, 404, 200, 500, 400, 401, 401, 401, 404,
		}

		s, err := series.NewSeriesWithList[int]("http_code", nil, utils.ConvertArrayToArrayPointer(data))
		if err != nil {
			panic(err)
		}
		Convey("Counter string", func() {
			out := SeriesValuesCount(s)
			utils.PrettyPrint2(out)
		})

		Convey("Test with apply", func() {
			out, err := series.Apply(s, SeriesValuesCount[int, int])
			if err != nil {
				return
			}
			utils.PrettyPrint2(out)
		})
	})
}

func TestDataFrameValuesCount(t *testing.T) {
	Convey("Test DataframeValuesCount function", t, func() {
		data1 := []string{
			"500", "400", "500", "401", "404", "404", "200", "500", "400", "401", "401", "401", "404",
		}

		s1, err := series.NewSeriesWithList[string]("http_code", nil, utils.ConvertArrayToArrayPointer(data1))
		if err != nil {
			panic(err)
		}

		data2 := []string{
			"20000", "40092", "20000", "40100", "77777", "20000", "77777", "20000", "40092", "40100", "20000", "40092", "77777",
		}

		s2, err := series.NewSeriesWithList[string]("status_code", nil, utils.ConvertArrayToArrayPointer(data2))
		if err != nil {
			panic(err)
		}

		data3 := []string{
			"500", "400", "500", "400", "77777", "200", "77777", "20000", "401", "40100", "401", "40092", "77777",
		}

		s3, err := series.NewSeriesWithList[string]("fake_col", nil, utils.ConvertArrayToArrayPointer(data3))
		if err != nil {
			panic(err)
		}

		df := &dataframes.DataFrame[string]{
			Indexes: nil,
			Columns: []string{"http_code", "status_code", "fake_col"},
			Values:  nil,
			Series:  []*series.Series[string]{s1, s2, s3},
			DType:   "",
		}

		Convey("Uint test", func() {
			out := DataframeValuesCount(df)
			utils.PrettyPrint2(out)
		})
	})
}
