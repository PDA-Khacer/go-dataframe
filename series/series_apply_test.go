package series

import (
	"fmt"
	"github.com/PDA-Khacer/go-dataframe/utils"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSeriesApply(t *testing.T) {
	Convey("Test Apply function Series", t, func() {
		var index []string
		Convey("Apply function square", func() {
			index = []string{"a", "b", "c"}
			data := map[string]int{
				"a": 1, "b": 2, "c": 3,
			}

			series, err := NewSeries[int]("", index, data)
			if err != nil {
				panic(err)
			}
			series.Apply(func(i *int) *int {
				if i != nil {
					temp := *i * *i
					return &temp
				}
				return i
			})
			s, _ := utils.PrettyPrint(series)
			fmt.Println(string(s))
		})

		Convey("Apply function count length", func() {
			index = []string{"a", "b", "c"}
			data := map[string]int{
				"a": 1, "b": 2, "c": 3,
			}

			series, err := NewSeries[int]("", index, data)
			if err != nil {
				panic(err)
			}
			series.Apply(func(i *int) *int {
				if i != nil {
					temp := len(fmt.Sprintf("%d", *i))
					return &temp
				}
				return i
			})
			s, _ := utils.PrettyPrint(series)
			fmt.Println(string(s))
		})
	})
}
