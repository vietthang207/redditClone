package utils

import (
	"fmt"
	"testing"
)

func AssertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	message = fmt.Sprintf("%s: %v != %v", message, a, b)
	t.Errorf(message)
}

func AssertTrue(t *testing.T, a bool, message string) {
	AssertEqual(t, a, true, message)
}
func AssertFalse(t *testing.T, a bool, message string) {
	AssertEqual(t, a, false, message)
}
