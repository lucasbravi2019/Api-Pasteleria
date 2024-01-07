package util

import "strconv"

func Int64SliceToString(arr *[]int64) *[]string {
	newArr := NewList[string]()

	if arr == nil {
		return &newArr
	}

	for _, number := range *arr {
		Add(&newArr, strconv.FormatInt(number, 10))
	}

	return &newArr
}
