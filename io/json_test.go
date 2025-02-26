package io

import (
	"dataframe/series"
	"dataframe/utils"
	. "github.com/smartystreets/goconvey/convey"
	"strconv"
	"testing"
)

func TestSeriesJsonNormalize(t *testing.T) {
	Convey("Test SeriesJsonNormalize", t, func() {
		Convey("Happy case", func() {
			index := []string{"1", "2", "3"}
			data := map[string]string{
				"1": `{"customer": { "email": "1111@onepay.vn", "id": "1", "name": "Test User 1", "phone": "*********21" }}`,
				"2": `{"customer": { "email": "2222@onepay.vn", "id": "2", "name": "Test Usr 2", "phone": "*********21" }}`,
				"3": `{"customer": { "email222": "2222@onepay.vn", "id11": "2", "name": "Test Usr 2", "phone22": "*********21" }}`,
			}

			seriesMock, err := series.NewSeries[string]("user_info", index, data)
			if err != nil {
				return
			}
			utils.PrettyPrint2(seriesMock)

			// convert json string
			normalize, err := SeriesJsonNormalize(*seriesMock)
			if err != nil {
				return
			}

			normalize.Apply(func(s *string) *string {
				// get len
				if s != nil {
					l := strconv.Itoa(len(*s))
					return &l
				}
				return nil
			})

			utils.PrettyPrint2(normalize)
		})
	})
}
