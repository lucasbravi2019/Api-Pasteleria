package util

import "testing"

func TestToLong(t *testing.T) {
	var expected1 int64 = 3
	result1, _ := ToLong("3")
	var expected2 int64 = -3
	result2, _ := ToLong("-3")

	AssertEquals[int64](t, expected1, result1)
	AssertEquals[int64](t, expected2, result2)
}
