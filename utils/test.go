package utils

import (
	"github.com/samber/lo"
	"strconv"
)

func SampleIntMatrix() (data [][]*int) {
	for i := 0; i < 3; i++ {
		var temp []*int
		j1 := i*3 + 1
		temp = append(temp, &j1)
		j2 := i*3 + 2
		temp = append(temp, &j2)
		j3 := i*3 + 3
		temp = append(temp, &j3)
		data = append(data, temp)
	}
	return
}

func SampleIntMapMatrix() (data map[string][]*int, index []string) {
	data = map[string][]*int{}
	for i := 0; i < 3; i++ {
		var temp []*int
		j1 := i*3 + 1
		temp = append(temp, &j1)
		j2 := i*3 + 2
		temp = append(temp, &j2)
		j3 := i*3 + 3
		temp = append(temp, &j3)
		index = append(index, strconv.Itoa(i))
		data[strconv.Itoa(i)] = temp
	}
	return
}

func SampleStringMapMatrix() (data map[string][]*string, index []string) {
	data = map[string][]*string{}
	for i := 0; i < 3; i++ {
		var temp []*string
		j1 := strconv.Itoa(i*3 + 1)
		temp = append(temp, &j1)
		j2 := strconv.Itoa(i*3 + 2)
		temp = append(temp, &j2)
		j3 := strconv.Itoa(i*3 + 3)
		temp = append(temp, &j3)
		index = append(index, strconv.Itoa(i))
		data[strconv.Itoa(i)] = temp
	}
	return
}

func ConvertArrayToArrayPointer[T any](source []T) []*T {
	return lo.Map(source, func(item T, _ int) *T {
		return &item
	})
}
