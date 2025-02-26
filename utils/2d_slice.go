package utils

import "errors"

func GetColValuesOf2DSlice[T any](_2dSlice [][]*T, indexCol int) []*T {
	if len(_2dSlice) > 0 {
		totalCol := len(_2dSlice[0])
		if totalCol-1 < indexCol {
			return nil
		}
		var re []*T
		for i := 0; i < len(_2dSlice); i++ {
			re = append(re, _2dSlice[i][indexCol])
		}
		return re
	}
	return nil
}

/*
Convert2DMapRowTo2DSlice input:

	data := map[string][]int{
			"row1": {1,2,3},
			"row2": {4,5,6},
			"row3": {7,8,9}
		}

Output:

	data := [][]int{
			{1,2,3},
			{4,5,6},
			{7,8,9},
	}

index: row1, row2, row3
*/
func Convert2DMapRowTo2DSlice[T any](_2dMap map[string][]*T) (result [][]*T, indexName []string) {
	for k, v := range _2dMap {
		indexName = append(indexName, k)
		result = append(result, v)
	}
	return
}

func GetColValuesOf2DMapRow[T any](_2dMap map[string][]*T, indexCol int) (result []*T, indexName []string, err error) {
	for k, v := range _2dMap {
		indexName = append(indexName, k)
		if len(v) <= indexCol {
			return nil, nil, errors.New("indexCol greater than number col")
		}
		result = append(result, v[indexCol])
	}
	return
}

func GetColValuesOf2DMapRowAndIndex[T any](_2dMap map[string][]*T, indexCol int, indexName []string) (result []*T, err error) {
	for _, idx := range indexName {
		if v, ok := _2dMap[idx]; !ok {
			return nil, errors.New("index name not exist: " + idx)
		} else {
			if len(v) <= indexCol {
				return nil, errors.New("indexCol greater than number col")
			}
			result = append(result, v[indexCol])
		}
	}
	return
}

func GetNumberColOfMapRow[T any](_2dMap map[string][]*T) int {
	for _, v := range _2dMap {
		return len(v)
	}
	return 0
}
