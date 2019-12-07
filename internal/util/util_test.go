package util

import (
	"testing"
)

const expectedValueString = "foo"
const expectedValueInt = 10

func TestPickStringLast(t *testing.T) {
	if v := PickString("", "", expectedValueString); v != expectedValueString {
		t.Fatalf("expected %s, got %s", expectedValueString, v)
	}
}

func TestPickStringNoValue(t *testing.T) {
	if v := PickString(""); v != "" {
		t.Fatalf("expected '%s', got '%s'", "", v)
	}
}

func TestPickStringFirst(t *testing.T) {
	if v := PickString(expectedValueString, "bar"); v != expectedValueString {
		t.Fatalf("expected %s, got %s", expectedValueString, v)
	}
}

func TestPickIntLast(t *testing.T) {
	if v := PickInt(0, 0, expectedValueInt); v != expectedValueInt {
		t.Fatalf("expected %d, got %d", expectedValueInt, v)
	}
}

func TestPickIntNoValue(t *testing.T) {
	if v := PickInt(0); v != 0 {
		t.Fatalf("expected %d, got %d", 0, v)
	}
}
func TestPickIntFirst(t *testing.T) {
	if v := PickInt(expectedValueInt, 5); v != expectedValueInt {
		t.Fatalf("expected %d, got %d", expectedValueInt, v)
	}
}