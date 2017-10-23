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

func AssertMap(t *testing.T, a map[int]int, b map[int]int, message string) {
	if len(a) != len(b) {
		message = fmt.Sprintf("%s. Different length: %v != %v", message, len(a), len(b))
		t.Errorf(message)
		return
	}
	for k, v := range a {
		if _, ok := b[k]; !ok || v != b[k] {
			message = fmt.Sprintf("%s: Different in key %v: %v != %v", message, k, a[k], b[k])
			t.Errorf(message)
			return
		}
	}
}

func AssertList(t *testing.T, a []int, b []int, message string) {
	if len(a) != len(b) {
		message = fmt.Sprintf("%s. Different length: %v != %v", message, len(a), len(b))
		t.Errorf(message)
		return
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			message = fmt.Sprintf("%s: Different in key %v: %v != %v", message, i, a[i], b[i])
			t.Errorf(message)
			return
		}
	}
}

func AssertTrue(t *testing.T, a bool, message string) {
	AssertEqual(t, a, true, message)
}
func AssertFalse(t *testing.T, a bool, message string) {
	AssertEqual(t, a, false, message)
}
