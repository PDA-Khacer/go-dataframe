package series

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSeriesCreate(t *testing.T) {
	Convey("Test create Series", t, func() {
		var index []string
		Convey("Create Series with index having", func() {
			index = []string{"a", "b", "c"}
			data := map[string]int{
				"a": 1, "b": 2, "c": 3,
			}

			series, err := NewSeries[int]("", index, data)
			if err != nil {
				return
			}
			fmt.Println(series)
		})

		Convey("Create Series with index not existed", func() {
			index = []string{"a", "b", "c"}
			data := map[string]int{
				"a": 1, "e": 2, "1": 3,
			}

			series, err := NewSeries[int]("", index, data)
			if err != nil {
				return
			}
			fmt.Println(series)
		})

	})
}
