package util

import (
	"reflect"
	"testing"
)

func AssertEquals[T any](t *testing.T, expectedValue, resultValue T) {
	if !equals(expectedValue, resultValue) {
		FailValue(t, expectedValue, resultValue)
	}
}

func equals[T any](expectedValue, resultValue T) bool {
	return reflect.DeepEqual(expectedValue, resultValue)
}

func FailValue[T any](t *testing.T, valueExpected T, valueResult T) {
	t.Errorf("Value expected %v, Value result %v", valueExpected, valueResult)
}
