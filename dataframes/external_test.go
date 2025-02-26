package dataframes

import (
	"dataframe/utils"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestFullAsType(t *testing.T) {
	Convey("Test AsType function", t, func() {
		Convey("Convert int to string", func() {
			dataIntMap, index2 := utils.SampleIntMapMatrix()
			column := []string{"col1", "col2", "col3"}
			frame, err := NewDataframeWithRowMap[int](index2, column, dataIntMap)
			if err != nil {
				panic(err)
			}
			utils.PrettyPrint2(frame)

			dfString, err := AsType[int, string](frame)
			if err != nil {
				panic(err)
			}
			utils.PrettyPrint2(dfString)
		})
	})
}
