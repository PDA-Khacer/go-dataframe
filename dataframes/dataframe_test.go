package dataframes

import (
	"fmt"
	"github.com/PDA-Khacer/go-dataframe/utils"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDataFrameCreate(t *testing.T) {
	Convey("Test create DataFrame", t, func() {
		var index []string
		var column []string
		Convey("Create DataFrame with 2D slice", func() {
			data := utils.SampleIntMatrix()
			column = []string{"col1", "col2", "col3"}
			frame, err := NewDataframe[int](index, column, data)
			if err != nil {
				return
			}

			b, _ := utils.PrettyPrint(frame)
			fmt.Println(string(b))
		})

		Convey("Create DataFrame with 2D map slice", func() {
			dataMap, index2 := utils.SampleIntMapMatrix()
			Convey("Create DataFrame with 2D map slice having index", func() {
				column = []string{"col1", "col2", "col3"}
				frame, err := NewDataframeWithRowMap[int](index2, column, dataMap)
				if err != nil {
					return
				}
				b, _ := utils.PrettyPrint(frame)
				fmt.Println(string(b))
			})

			Convey("Create DataFrame with 2D map slice no index", func() {
				column = []string{"col1", "col2", "col3"}
				frame, err := NewDataframeWithRowMap[int](index, column, dataMap)
				if err != nil {
					return
				}
				b, _ := utils.PrettyPrint(frame)
				fmt.Println(string(b))
			})
		})
	})
}

func TestDataFrameAgg(t *testing.T) {
	Convey("Test Agg function", t, func() {
		var index []string
		var column []string
		Convey("Create DataFrame with 2D slice", func() {
			data := utils.SampleIntMatrix()
			column = []string{"col1", "col2", "col3"}
			frame, err := NewDataframe[int](index, column, data)
			if err != nil {
				return
			}

			utils.PrettyPrint2(frame)
			agg, err := frame.Agg([]string{"max", "min"})
			if err != nil {
				return
			}
			utils.PrettyPrint2(agg)
		})
	})
}

func TestDataFrameBasicFunction(t *testing.T) {
	Convey("Test Basic function", t, func() {
		var index []string
		var column []string
		Convey("Create DataFrame with 2D slice", func() {
			data := utils.SampleDataFrameHavingCol2thSameValue(2)
			column = []string{"col1", "col2", "col3"}
			frame, err := NewDataframe[int](index, column, data)
			if err != nil {
				return
			}

			utils.PrettyPrint2(frame)
			valuesNeedDrop := 2
			col, agg := frame.DropColIfAllValueIs(&valuesNeedDrop, true)
			So(col, ShouldEqual, []string{"col2"})
			utils.PrettyPrint2(agg)
		})
	})
}
