package series

import (
	"fmt"
	"github.com/PDA-Khacer/go-dataframe/utils"
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
			utils.PrettyPrint2(series)
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
			utils.PrettyPrint2(series)
		})
	})
}

func TestSeriesApplyFunction(t *testing.T) {
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
			utils.PrettyPrint2(series)
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
			utils.PrettyPrint2(series)
		})
	})
}

func TestSeriesBasicFunction(t *testing.T) {
	Convey("Test Basic function Series", t, func() {
		var index []string
		Convey("Drop nil function", func() {
			index = []string{"a", "b", "c", "d"}
			number1 := 1
			number2 := 2
			number3 := 3
			data := map[string]*int{
				"a": &number1, "b": nil, "c": &number3, "d": &number2,
			}

			series, err := NewSeriesPointer[int]("", index, data)
			if err != nil {
				panic(err)
			}
			idx, _ := series.DropNil(true)
			So(idx, ShouldEqual, []string{"b"})
			utils.PrettyPrint2(series)
		})

		Convey("Drop nil if function", func() {
			Convey("Valid nil", func() {
				index = []string{"a", "b", "c", "d"}
				number1 := 1
				number2 := 2
				number3 := 3
				data := map[string]*int{
					"a": &number1, "b": nil, "c": &number3, "d": &number2,
				}

				series, err := NewSeriesPointer[int]("", index, data)
				if err != nil {
					panic(err)
				}
				idx, _ := series.DropIf(true, func(i *int) bool {
					return i == nil
				})
				So(idx, ShouldEqual, []string{"b"})
				utils.PrettyPrint2(series)
			})

			Convey("Valid string", func() {
				index = []string{"a", "b", "c", "d"}
				string1 := "1"
				string2 := ""
				string3 := "3"
				data := map[string]*string{
					"a": &string1, "b": nil, "c": &string2, "d": &string3,
				}

				series, err := NewSeriesPointer[string]("", index, data)
				if err != nil {
					panic(err)
				}
				idx, _ := series.DropIf(true, func(i *string) bool {
					if i != nil && *i == "" {
						return true
					}
					return false
				})
				So(idx, ShouldEqual, []string{"c"})
				utils.PrettyPrint2(series)
			})
		})
	})
}
