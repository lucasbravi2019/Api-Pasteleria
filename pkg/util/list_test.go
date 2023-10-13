package util

import "testing"

func TestToList(t *testing.T) {
	expected := []int{1, 2, 3}
	result := ToList(1, 2, 3)

	sizeExpected := len(expected)
	sizeResult := len(result)

	AssertEquals(t, sizeExpected, sizeResult)

	for i := 0; i < len(result); i++ {
		AssertEquals(t, expected[i], result[i])
	}
}

func TestAdd(t *testing.T) {
	list := []int{}
	expected := []int{1, 2, 4}
	Add(&list, 1, 2, 3)

	expectedSize := len(expected)
	resultSize := len(list)
	AssertEquals(t, expectedSize, resultSize)

	for i := 0; i < resultSize; i++ {
		expectedValue := expected[i]
		resultValue := list[i]

		AssertEquals(t, expectedValue, resultValue)
	}
}
