package utils

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test2DSlice(t *testing.T) {
	Convey("Test 2DSlice", t, func() {
		var data [][]*int
		Convey("Happy GetColValuesOf2DSlice case", func() {
			data = SampleIntMatrix()
			a := GetColValuesOf2DSlice[int](data, 2)
			tt, _ := PrettyPrint(a)
			fmt.Println(string(tt))
		})

		Convey("Happy GetColValuesOf2DMapRow case", func() {
			dataMap, _ := SampleIntMapMatrix()
			a, _, _ := GetColValuesOf2DMapRow[int](dataMap, 2)
			tt, _ := PrettyPrint(a)
			fmt.Println(string(tt))
		})
	})
}
